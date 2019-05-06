# {name}

{go:header}

A faster goimports implementation. It assumes that the files are part of a Go module and uses the information inside `go.mod` to find the import paths for external dependencies.

{go:install}

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

{go:footer}
