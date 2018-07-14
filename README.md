# pkg `walkup` - files and directories walk function going through parent directories recursively, written in Go

## Installation

```
go get -u github.com/Qs-F/walkup
```

## Usge

```go
walkup.Walkup("./", ".hello", 0)
```

## Returns

`[]string` of including file or directory names witch you seeked.
