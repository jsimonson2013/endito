package page

import (
	"encoding/json"
	"endito/files"
	"endito/git"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Load reads the uri of a file string and returns the contents of the file
func Load(bs []byte) ([]byte, error) {
	var body map[string]interface{}
	if err := json.Unmarshal(bs, &body); err != nil {
		return nil, err
	}

	// open file for reading
	f, err := os.OpenFile(body["uri"].(string), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(f.Name())
}

// Update uses various fields of the byte slice body to write to a file, and commit the changes
func Update(bs []byte) error {
	var body map[string]interface{}
	if err := json.Unmarshal(bs, &body); err != nil {
		return err
	}

	// check auth matches
	if body["uname"] != os.Getenv("USERNAME") || body["pword"] != os.Getenv("PASSWORD") {
		return fmt.Errorf("bad username or password")
	}

	// write contents to file
	if err := ioutil.WriteFile(body["uri"].(string), []byte(body["content"].(string)), os.ModePerm); err != nil {
		return err
	}

	// getrelative paths for git updates
	rltv, err := files.RelativePaths(os.Getenv("RLTV_DIR"), []string{body["uri"].(string)})
	if err != nil {
		return err
	}

	// commit and push changes
	if err := git.CommitAndPush(os.Getenv("RLTV_DIR"), fmt.Sprintf("%s updated %s", os.Getenv("GIT_UNAME"), strings.Join(rltv, ",")), rltv); err != nil {
		return err
	}

	return nil
}
