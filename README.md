# autoimport

A faster goimports implementation. It assumes that the files are part of a Go module and uses the information inside `go.mod` to find the import paths for external dependencies.

## Installation

```bash
go install github.com/akyoto/autoimport/...
```

## CLI

```bash
autoimport hello.go world.go
```

## API

### New

```go
importer := autoimport.New("/home/user/projects/helloworld")
```

Creates a new importer inside the given project directory. The root directory must have a `go.mod` file.

### Files

```go
importer.Files([]string{
	"hello.go",
	"world.go",
})
```

Automatically adds import statements to the given files.

### Source

```go
importer.Source([]byte{"GO CODE HERE"})
```
