# Goupfile [![Build Status](https://travis-ci.org/goupfile/server.svg?branch=master)](https://travis-ci.org/goupfile/server) [![Go Report Card](https://goreportcard.com/badge/github.com/goupfile/server)](https://goreportcard.com/report/github.com/goupfile/server)

Goupfile is a service that allows you to easily upload and share files without
leaving your terminal.

## Status

Goupfile is currently in alpha, which means that anything can change, and files
could disappear without warning.

Nevertheless, you can try it out: https://goupfile.com

## Features

What makes this file upload site different?

- Share multiple files under one URL
- URLs are short, memorable, and don't have ambiguous characters
- Web UI is free of bloat
- There's a [CLI tool](https://github.com/goupfile/goup)
  (under development) that makes it easy to upload files from your terminal

## Developing

`go get` will fetch, build, and install the package. You can then run the
server locally.

```
go get github.com/goupfile/server
$GOPATH/bin/server
```

To upload a file, send a multipart/form-data `POST` request to `127.0.0.1:8080`.
The form should contain a key named "file".

To download a file, send a `GET` request to `127.0.0.1:8080/[filename]`

### With Docker

You will need [Docker Engine](https://docs.docker.com/install/) and
[Docker Compose](https://docs.docker.com/compose/).

```
git@github.com:goupfile/server.git
cd server
docker-compose up
```

## License

MIT
