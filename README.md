# Goupfile core

The main server application. It handles file upload, download, and storage.

## Developing

`go get` will fetch, build, and install the package. You can then run the
server locally.

```
go get github.com/goupfile/core
$GOPATH/bin/core
```

To upload a file, send a multipart/form-data POST request to 127.0.0.1:8080.
The form should contain a key named "file".

To download a file, send a GET request to 127.0.0.1:8080/[filename]

## License

Unlicense
