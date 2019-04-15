# autoimport

A faster goimports implementation. It assumes that the files are part of a Go module and uses the information inside `go.mod` to find the import paths for external dependencies.

## Installation

```bash
go get github.com/blitzprog/autoimport/...
```

## CLI

```bash
autoimport hello.go world.go
```

## API

### Files

```go
autoimport.Files([]string{
	"hello.go",
	"world.go",
})
```

Automatically adds import statements to the given files.

### GetPackagesInDirectory

```go
autoimport.GetPackagesInDirectory(path string) map[string][]*Package
```

Returns a map of package names mapped to the possible import paths.