# autoimport

[![Reference][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]

A faster goimports implementation. It assumes that the files are part of a Go module and uses the information inside `go.mod` to find the import paths for external dependencies.

## Installation

```bash
go get github.com/akyoto/autoimport/...
```

## CLI

```bash
autoimport hello.go world.go
```

Work in progress (currently API only, CLI is not available yet).

## API

### New

```go
importer := autoimport.New("/home/user/projects/helloworld")
```

Creates a new importer inside the given project directory. The root directory must have a `go.mod` file. Otherwise it will traverse the directories going upwards until it finds a `go.mod` file.

### Files

```go
importer.Files([]string{
	"hello.go",
	"world.go",
})
```

Work in progress. Automatically adds import statements to the given files.

### Source

```go
importer.Source([]byte{"GO CODE HERE"})
```

Work in progress.

[godoc-image]: https://godoc.org/github.com/akyoto/autoimport?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/autoimport
[report-image]: https://goreportcard.com/badge/github.com/akyoto/autoimport
[report-url]: https://goreportcard.com/report/github.com/akyoto/autoimport
[tests-image]: https://cloud.drone.io/api/badges/akyoto/autoimport/status.svg
[tests-url]: https://cloud.drone.io/akyoto/autoimport
[codecov-image]: https://codecov.io/gh/akyoto/autoimport/graph/badge.svg
[codecov-url]: https://codecov.io/gh/akyoto/autoimport
[license-image]: https://img.shields.io/badge/license-MIT-blue.svg
[license-url]: https://github.com/akyoto/autoimport/blob/master/LICENSE