# Goupfile core [![Go Report Card](https://goreportcard.com/badge/github.com/goupfile/core)](https://goreportcard.com/report/github.com/goupfile/core)

The main server application for Goupfile. It handles file upload, download, and
storage.

## Developing

`go get` will fetch, build, and install the package. You can then run the
server locally.

```
go get github.com/goupfile/core
$GOPATH/bin/core
```

To upload a file, send a multipart/form-data `POST` request to `127.0.0.1:8080`.
The form should contain a key named "file".

To download a file, send a `GET` request to `127.0.0.1:8080/[filename]`

### With Docker

You will need [Docker Engine](https://docs.docker.com/install/) and
[Docker Compose](https://docs.docker.com/compose/).

```
git@github.com:goupfile/core.git
cd core
docker-compose up
```

## License

Unlicense
