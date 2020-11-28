package page

import (
	"endito/files"
	"fmt"
	"os"
	"strings"
)

// GetAll returns the html files descendant from a base directory
func GetAll(base string) ([]string, error) {
	// get a list of files
	fs, err := files.FromDir(base, nil)
	if err != nil {
		return nil, err
	}

	// get base directory from environment variables
	bd := os.Getenv("BASE_DIR")
	if bd == "" {
		return nil, fmt.Errorf("$BASE_DIR not set")
	}

	// build path to editor file
	edtr := fmt.Sprintf("%s/editor/index.html", strings.TrimRight(bd, "/"))

	// ignore editor file
	return files.Filter(fs, edtr), nil
}
