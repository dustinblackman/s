<p align="center">
<img height="200" src="https://i.imgur.com/r4khD2u.png">
<br />
<a href="https://travis-ci.org/dustinblackman/s"><img src="https://img.shields.io/travis/dustinblackman/s.svg" alt="Build Status"></a> <a href="https://goreportcard.com/report/github.com/dustinblackman/s"><img src="https://goreportcard.com/badge/github.com/dustinblackman/s"></a>
</p>

<p align="center">A simple command line utility for posting to social networks.</p>

__Currently Supported:__
- Twitter

## Usage
Without any parameters, S will post to all social networks with available configs.

```bash
s Going out for poutine.
```

You can also specifiy just specific social networks. For example, just Twitter:

```bash
s -t Posting this wonderful tweet from command line!
```

## Install

Grab the latest release from the [releases](https://github.com/dustinblackman/s/releases) page, or install directly from master. S is currently built and tested against Go 1.7.

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

## Alfred Workflows

Alfred workflows are available [here](./alfred-workflows/) which allows you to use S directly from Alfred. Due to S needing access to environment variables for configuration, there are two editions to the work flows for `bash` and `zsh`.
