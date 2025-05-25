package phrasegen_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	phrasegen "t1pw40p/tools/phrasegen/internal"
	"testing"
)

func TestClean(t *testing.T) {
	expected := "This 1s a TEST string"
	uncleaned := "'~This 1s a T-E-S-T! string.<>~'"
	cleaned := phrasegen.Clean([]byte(uncleaned))
	if cleaned != expected {
		t.Errorf("Cleaned (%s) did not match expected: %s", cleaned, expected)
	}
}

func TestSplitOnSpace(t *testing.T) {
	input := "This is a    sentence.\n\n\nIt    has    many   \n    spaces."
	expected := []string{"This", "is", "a", "sentence.", "It", "has", "many", "spaces."}
	out := phrasegen.SplitOnSpace(input)
	if !slices.Equal(out, expected) {
		t.Errorf("Split on space expected '%s' got '%s'", expected, out)
	}
}

func TestLoadFileExistsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	f, err := os.CreateTemp(t.TempDir(), "testfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	data, err := phrasegen.LoadFile(f.Name())
	if err != nil {
		t.Errorf("Encountered err reading valid file? %s", err)
	}
	if data != "" {
		t.Errorf("Expected empty data from empty temp file?: %s", data)
	}

	sample := "Some sample data!"
	f.WriteString(sample)
	data, err = phrasegen.LoadFile(f.Name())
	if err != nil {
		t.Errorf("Encountered err reading valid file? %s", err)
	}
	if data != sample {
		t.Errorf("Didn't receive expected data from temp file?: %s", data)
	}
}

func TestLoadFileDneIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	dir := t.TempDir()
	fname := filepath.Join(dir, "does_not_exist")
	_, err := phrasegen.LoadFile(fname)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Expected no such file, got %s", err)
	}
}
