<p align="center">
<img height="200" src="https://i.imgur.com/r4khD2u.png">
<br />
<a href="https://travis-ci.org/dustinblackman/s"><img src="https://img.shields.io/travis/dustinblackman/s.svg" alt="Build Status"></a> <a href="https://goreportcard.com/report/github.com/dustinblackman/s"><img src="https://goreportcard.com/badge/github.com/dustinblackman/s"></a> <img src="https://img.shields.io/github/release/dustinblackman/s.svg?maxAge=2592000">
</p>

<p align="center">A simple command line utility for posting to social networks.</p>

__Currently Supported:__
- Twitter
- Facebook

## Usage
Without any parameters, S will post to all social networks with available configs.

```bash
s Going out for poutine.
```

You can also specifiy just a single social network. For example, just Twitter:

```bash
s -t Posting this wonderful tweet from command line!
```

## Install

Grab the latest release from the [releases](https://github.com/dustinblackman/s/releases) page, or build from source and install directly from master. S is currently built and tested against Go 1.7.

```bash
git pull https://github.com/dustinblackman/s.git
cd ./s
make install
```

## Configuration

Configuration for social networks can be done by setting the required keys in your environment variables, but it's also possible to pass them in as parameters. See `s --help` for more details.

__Twitter:__

Creating a twitter application can be done [here](https://apps.twitter.com/app/new). You can then generate keys and save them in your environment variables.

```bash
export TWITTER_CONSUMER_KEY=""
export TWITTER_CONSUMER_SECRET=""
export TWITTER_ACCESS_TOKEN=""
export TWITTER_ACCESS_SECRET=""
```

__Facebook:__

Create an application on Facebook [here](https://developers.facebook.com/docs/apps/register). Afterwards use the [Graph explorer](https://developers.facebook.com/docs/apps/register) to create a user access token that has the `publish_actions` scope. You can extend the life of the key by clicking the `I` next to the token to open the access token tools.

```bash
export FACEBOOK_APP_KEY=""
export FACEBOOK_APP_SECRET=""
export FACEBOOK_USER_TOKEN=""
```

## Alfred Workflows

Alfred workflows are available [here](./alfred-workflows/) which allows you to use S directly from Alfred. Due to S needing access to environment variables for configuration, there are two editions to the work flows for `bash` and `zsh`.
