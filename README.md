# Goupfile [![Build Status](https://travis-ci.org/goupfile/server.svg?branch=master)](https://travis-ci.org/goupfile/server) [![Go Report Card](https://goreportcard.com/badge/github.com/goupfile/server)](https://goreportcard.com/report/github.com/goupfile/server)

Goupfile is a file sharing service.

## Features

What makes this one different?

- Share multiple files under one URL
- URLs are short, memorable, and don't have ambiguous characters
- QR codes so that you can upload files on one device and easily access them on another
- Upload from any browser at [goupfile.com](https://goupfile.com)
- There's a [CLI tool](https://github.com/goupfile/up) for uploading files from the terminal
- No dependencies: it uses a SQLite database and saves files to the local filesystem
- Easy to deploy: just download a single binary and run
- Lightweight: runs on any machine in the cloud

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
