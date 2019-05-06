# autoimport

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

A faster goimports implementation. It assumes that the files are part of a Go module and uses the information inside `go.mod` to find the import paths for external dependencies.

## Installation

```shell
go get -u github.com/blitzprog/autoimport/...
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
importer.Files([]string)
```

Work in progress. Automatically adds import statements to the given files.

### Source

```go
importer.Source([]byte)
```

Work in progress.

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://www.patreon.com/eduardurbach)

[godoc-image]: https://godoc.org/github.com/blitzprog/autoimport?status.svg
[godoc-url]: https://godoc.org/github.com/blitzprog/autoimport
[report-image]: https://goreportcard.com/badge/github.com/blitzprog/autoimport
[report-url]: https://goreportcard.com/report/github.com/blitzprog/autoimport
[tests-image]: https://cloud.drone.io/api/badges/blitzprog/autoimport/status.svg
[tests-url]: https://cloud.drone.io/blitzprog/autoimport
[coverage-image]: https://codecov.io/gh/blitzprog/autoimport/graph/badge.svg
[coverage-url]: https://codecov.io/gh/blitzprog/autoimport
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
