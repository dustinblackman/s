package main

import (
	"os"
	"strings"
	"sync"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/huandu/facebook"
	"github.com/urfave/cli"
)

var (
	version = "HEAD"
)

// SCtx holds the context for the current exeuction
type SCtx struct {
	ctx     *cli.Context
	message string
	wg      *sync.WaitGroup
}

// postTwitter submits a tweet through Twitter's API.
func (S *SCtx) postTwitter() {
	defer S.wg.Done()

	log.Info("Tweeting: " + S.message)
	anaconda.SetConsumerKey(S.ctx.String("twitter-consumer-key"))
	anaconda.SetConsumerSecret(S.ctx.String("twitter-consumer-secret"))
	api := anaconda.NewTwitterApi(S.ctx.String("twitter-acesss-token"), S.ctx.String("twitter-access-secret"))
	_, err := api.PostTweet(S.message, nil)
	if err != nil {
		log.Error(err)
	}
}

func (S *SCtx) postFacebook() {
	defer S.wg.Done()

	log.Info("Submitting to Facebook: " + S.message)
	app := facebook.New(S.ctx.String("facebook-app-key"), S.ctx.String("facebook-app-secret"))
	session := app.Session(S.ctx.String("facebook-user-token"))
	_, err := session.Post("/me/feed", facebook.Params{"message": S.message})
	if err != nil {
		log.Error(err)
	}
}

func (S *SCtx) checkMissingKeys(keys []string) bool {
	missingKey := false
	for _, key := range keys {
		if S.ctx.String(key) == "" {
			log.Error(key + " is not defined.")
			missingKey = true
		}
	}

	return !missingKey
}

// verifyTwitterConfig checks if the configs for Twitter exists before continuing, otherwise exits early.
func (S *SCtx) twitterConfigExists() bool {
	return S.checkMissingKeys([]string{"twitter-consumer-key", "twitter-consumer-secret", "twitter-acesss-token", "twitter-access-secret"})
}

func (S *SCtx) facebookConfigExists() bool {
	return S.checkMissingKeys([]string{"facebook-app-key", "facebook-app-secret", "facebook-user-token"})
}

// processContext parses the context passed from CLI.
func processContext(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		log.Error("You need to write a message to update your status... Who would want to read nothing?")
		cli.ShowAppHelp(ctx)
		return cli.NewExitError("", 1)
	}

	var wg sync.WaitGroup
	message := strings.Join(ctx.Args(), " ")
	S := SCtx{ctx, message, &wg}

	if !ctx.Bool("twitter") && !ctx.Bool("facebook") {
		if S.twitterConfigExists() {
			wg.Add(1)
			go S.postTwitter()
		}
		if S.facebookConfigExists() {
			wg.Add(1)
			go S.postFacebook()
		}
	}

	if ctx.Bool("twitter") {
		if !S.twitterConfigExists() {
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}
		wg.Add(1)
		go S.postTwitter()
	}

	if ctx.Bool("facebook") {
		if !S.facebookConfigExists() {
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}
		wg.Add(1)
		go S.postFacebook()
	}

	wg.Wait()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "S"
	app.Usage = "A simple command line utility for posting status messages to social networks"
	app.Version = version
	app.Author = "Dustin Blackman"
	app.Copyright = "(c) 2016 " + app.Author
	app.EnableBashCompletion = true
	app.Action = processContext

	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}
	https://github.com/dustinblackman/s/

EXAMPLE USAGE:
	s -t Going out for poutine.
{{if len .Authors}}
AUTHOR(S):
	{{range .Authors}}{{ . }}{{end}}{{end}}
{{if .VisibleFlags}}
GLOBAL OPTIONS:
{{range .VisibleFlags}}	{{.}}
{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
	{{.Copyright}}
{{end}}{{if .Version}}
VERSION:
	{{.Version}}
{{end}}
`

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "twitter, t",
			Usage: "Post tweet to Twitter",
		},
		cli.BoolFlag{
			Name:  "facebook, f",
			Usage: "Post status to Facebook",
		},
		cli.StringFlag{
			Name:   "twitter-consumer-key, tck",
			Usage:  "Twitter consumer key",
			EnvVar: "TWITTER_CONSUMER_KEY",
		},
		cli.StringFlag{
			Name:   "twitter-consumer-secret, tcs",
			Usage:  "Twitter consumer secret",
			EnvVar: "TWITTER_CONSUMER_SECRET",
		},
		cli.StringFlag{
			Name:   "twitter-acesss-token, tat",
			Usage:  "Twitter access token",
			EnvVar: "TWITTER_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "twitter-access-secret, tas",
			Usage:  "Twitter access token secret",
			EnvVar: "TWITTER_ACCESS_SECRET",
		},
		cli.StringFlag{
			Name:   "facebook-app-key, fak",
			Usage:  "Facebook app key",
			EnvVar: "FACEBOOK_APP_KEY",
		},
		cli.StringFlag{
			Name:   "facebook-app-secret, fas",
			Usage:  "Facebook app secret",
			EnvVar: "FACEBOOK_APP_SECRET",
		},
		cli.StringFlag{
			Name:   "facebook-user-token, fut",
			Usage:  "Facebook user token",
			EnvVar: "FACEBOOK_USER_TOKEN",
		},
	}

	app.Run(os.Args)
}
