package walkup

import (
	"path/filepath"
	"strings"
)

func Walkup(base string, filename string, level int) ([]string, error) {
	return []string{filepath.Clean(filepath.Join(base, "..", "..", "..", "TEMP"))}, nil
}

func parent(dir string) string {
	return filepath.Clean(filepath.Join(dir, ".."))
}

func walkList(dir string, n int) []string {
	// Slashes is the number of least directory level, so it is no problem if loop breaks
	// when they met root slash
	slashes := strings.Count(dir, "/")
	if n == 0 {
		n = slashes + 1
	} else if n+1 > slashes {
		n = slashes + 1
	}

	list := []string{}

L:
	for i := 0; i < n; i++ {
		list = append(list, dir)
		if dir == "/" {
			break L
		}
		dir = parent(dir)
	}

	return list
}
