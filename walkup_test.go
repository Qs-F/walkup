package walkup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkup(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	err = os.Chdir(filepath.Join(dir, "A", "B"))
	if err != nil {
		t.Error(err.Error())
	}

	current, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	// Walkup( /* basedir */, /* filename */, /* directory level (0 means to root dir) */ )
	filelist, err := Walkup(current, "TEMP", 0)
	if err != nil {
		t.Error(err.Error())
	}

	if !assert.Equal(t, filelist, []string{filepath.Join(dir, "TEMP")}) {
		t.Error("Files does not match")
	}
}
