# ForgeStatus

[![CircleCI](https://circleci.com/gh/chrisstowe/forgestatus.svg?style=svg)](https://circleci.com/gh/chrisstowe/forgestatus) [![Go Report Card](https://goreportcard.com/badge/github.com/chrisstowe/forgestatus)](https://goreportcard.com/report/github.com/chrisstowe/forgestatus) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/chrisstowe/forgestatus)

## Description

A distributed system status and metric checker. 🔎

A visual dashboard for this service can be found at [forgestatus-dashboard](https://github.com/chrisstowe/forgestatus-dashboard)

### Live Builds

[forgestatus.com](http://forgestatus.com)

[dev.forgestatus.com](http://dev.forgestatus.com)

## How to run locally

### Docker

```
$ docker-compose up
```

### Run the worker/server

Requires `GOPATH` to be set and `GOPATH/bin` in your path.

More info on why this is required: [How to Write Go Code](https://golang.org/doc/code.html)

```
$ go install ./..
$ forgestatus-server
$ forgestatus-worker
```
