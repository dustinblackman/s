package main

import (
	"os"
	"strings"
	"sync"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "HEAD"
)

// twitter submits a tweet through Twitter's API.
func twitter(message string, ctx *cli.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Info("Tweeting: " + message)
	anaconda.SetConsumerKey(ctx.String("twitter-consumer-key"))
	anaconda.SetConsumerSecret(ctx.String("twitter-consumer-secret"))
	api := anaconda.NewTwitterApi(ctx.String("twitter-acesss-token"), ctx.String("twitter-access-secret"))
	api.PostTweet(message, nil)
}

// verifyTwitterConfig checks if the configs for Twitter exists before continuing, otherwise exits early.
func verifyTwitterConfig(ctx *cli.Context) {
	missingKey := false
	for _, key := range []string{"twitter-consumer-key", "twitter-consumer-secret", "twitter-acesss-token", "twitter-access-secret"} {
		if ctx.String(key) == "" {
			log.Error(key + " is not defined.")
			missingKey = true
		}
	}
	if missingKey {
		cli.ShowAppHelp(ctx)
		os.Exit(1)
	}
}

// processContext the context passed from CLI.
func processContext(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		log.Error("You need to write a message to update your status... Who would want to read nothing?")
		cli.ShowAppHelp(ctx)
		return cli.NewExitError("", 1)
	}

	message := strings.Join(ctx.Args(), " ")
	var wg sync.WaitGroup
	if !ctx.Bool("twitter") && !ctx.Bool("facebook") {
		wg.Add(1)
		verifyTwitterConfig(ctx)
		go twitter(message, ctx, &wg)
	}
	if ctx.Bool("twitter") {
		wg.Add(1)
		verifyTwitterConfig(ctx)
		go twitter(message, ctx, &wg)
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
	}

	app.Run(os.Args)
}
