# MakePNG

[![Go Report Card](https://goreportcard.com/badge/github.com/KingGenius5/makesite)](https://goreportcard.com/report/github.com/KingGenius5/makesite)

MakePNG is a ChromeDP CLI that takes a screenshot of any webpage and saves it as a PNG. Currently takes pictures of the NYT front page; you'll have to manually change the URL until the args update comes along.

![NYT](nyt.png)

### Getting Started

```bash
    $ git clone git@github.com:KingGenius5/makePNG.git
    $ cd makePNG
    $ docker pull chromedp/headless-shell
    $ go build
    $ go run main.go
```

And you'll have a PNG image filed saved of the NYT front page at your fingertips.
