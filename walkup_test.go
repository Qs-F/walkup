package walkup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Qs-F/mintest"
	"github.com/stretchr/testify/assert"
)

func TestWalkup(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	err = os.Chdir(filepath.Join(dir, "_testdata", "A", "B"))
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Chdir(dir)

	current, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	// Walkup( /* basedir */, /* filename */, /* directory level (0 means to root dir) */ )
	filelist := Walkup(filepath.Join(current, "_testdata"), "TEMP", 0)

	if !assert.Equal(t, filelist, []string{filepath.Join(dir, "_testdata", "TEMP")}) {
		t.Error("Files does not match")
	}
}

func TestParent(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}
	base := filepath.Join(dir, "_testdata", "A", "B")

	if parent(base) != filepath.Join(dir, "_testdata", "A") {
		t.Error("Not parent directory")
	} else {
		t.Log(parent(base))
	}
}

func TestWalkList(t *testing.T) {
	type test struct {
		Dir string
		N   int
	}
	tests := mintest.TestSlices{
		{
			Input: test{
				Dir: "/_testdata/A/B",
				N:   0,
			},
			Must: []string{
				"/_testdata/A/B",
				"/_testdata/A",
				"/_testdata",
				"/",
			},
			Error: true,
		},
		{
			Input: test{
				Dir: "/_testdata/A/B/C/D",
				N:   2,
			},
			Must: []string{
				"/_testdata/A/B/C/D",
				"/_testdata/A/B/C",
			},
			Error: true,
		},
		{
			Input: test{
				Dir: "/_testdata/A/B/C/D",
				N:   5,
			},
			Must: []string{
				"/_testdata/A/B/C/D",
				"/_testdata/A/B/C",
				"/_testdata/A/B",
				"/_testdata/A",
				"/_testdata",
				"/",
			},
			Error: true,
		},
	}

	for n := range tests {
		assert.Equal(t, tests[n].Must, walkList(tests[n].Input.(test).Dir, tests[n].Input.(test).N))
	}
}
