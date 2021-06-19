# Goupfile [![Build status](https://github.com/goupfile/server/actions/workflows/goupfile.yml/badge.svg)](https://github.com/goupfile/server/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/goupfile/server)](https://goreportcard.com/report/github.com/goupfile/server)

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

## HTTP API

```
GET    /                   Show the home page and upload/download from there
POST   /upload             Upload a file (use multipart/form-data)
GET    /download?id={id}   Download a file
```

## Developing

`go get` will fetch, build, and install the package. You can then run the
server locally.

```
go get github.com/goupfile/server
$GOPATH/bin/server
```

### Docker

Using Docker, you can build and run Goupfile without having Go installed and
without gcc (since `mattn/go-sqlite3` is a cgo package and relies on gcc).

You will need [Docker Engine](https://docs.docker.com/install/) and
[Docker Compose](https://docs.docker.com/compose/).

```
git clone git@github.com:goupfile/server.git
cd server
docker build -t goupfile
docker container run -p 8090:8090 goupfile
```

### CSS

This project uses [Tailwind CSS](https://tailwindcss.com/). The following will
create a CSS file with all Tailwind classes, which is helpful in development
because you can use any Tailwind CSS utility. The file produced by `css-dev` is
*over 3 MB*, so don't use it in production!

```sh
npm install
npm run css-dev
```

For a production build,

```sh
npm run css-prod
```

This will produce a CSS file with only the classes you used in the HTML.

## License

MIT
