package walkup

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Walkup(base string, filename string, level int) []string {
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	dirs := walkList(base, level)
	list := []string{}
	wg.Add(len(dirs))
	for _, dir := range dirs {
		go func(dir string) {
			fi, err := os.Lstat(filepath.Join(dir, filename))
			if err != nil {
				wg.Done()
				return
			}
			if fi.Mode().IsRegular() || fi.Mode().IsDir() {
				mutex.Lock()
				list = append(list, filepath.Join(dir, filename))
				mutex.Unlock()
			}
			wg.Done()
		}(dir)
		continue
	}
	wg.Wait()
	return list
}

func parent(dir string) string {
	return filepath.Clean(filepath.Join(dir, ".."))
}

func walkList(dir string, n int) []string {
	dir = filepath.Clean(dir)
	// Slashes is the number of least directory level, so it is no problem if loop breaks
	// when they met root slash
	slashes := strings.Count(dir, string(filepath.Separator))
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
