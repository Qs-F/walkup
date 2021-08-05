# pkg `walkup`

Files and directories walk function going through parent directories recursively, written in Go.

This package is mainly used for finding the root `package.json` or `go.mod` file.

## Installation

```
go get -u github.com/Qs-F/walkup
```

## Usge

```go
walkup.Walkup("./", ".hello", 0)
// returns []string of including file or directory names witch you seeked.
```

## License

MIT
