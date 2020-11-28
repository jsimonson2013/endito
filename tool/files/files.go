package files

import (
	"io/ioutil"
	"os"
	"strings"
)

// List is specific to a list of filenames
type List []string

const pkg = "endito"

// FromDir builds a list from a dir string and current list of filenames
func FromDir(dir string, fs List) (List, error) {
	// get file info in dir
	rdir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// if only one fileinfo and it's valid, return it
	if len(rdir) == 1 && isValid(rdir[0]) {
		return append(fs, dir+rdir[0].Name()), nil
	}

	// range over slice of fileinfo from dir
	for _, f := range rdir {
		// recursively dive into non-git directories
		if f.IsDir() && !isGitDir(f) {
			fs, err = FromDir(dir+f.Name()+"/", fs)
			if err != nil {
				return nil, err
			}
		} else if isValid(f) {
			// add valid fileinfo
			fs = append(fs, dir+f.Name())
		}
	}

	// finished ranging over root
	return fs, nil
}

// Filter removes a file from the list if it exists
func Filter(l List, f string) List {
	for i, v := range l {
		if v == f {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// RelativePaths builds relative paths from absolute paths given
// the relative directory
func RelativePaths(rltvDir string, absPaths []string) (List, error) {
	rltvPaths := make(List, 0)

	for _, fp := range absPaths {
		// fp = "path/to/package/folder/file.ext", segs = { "path/to/package/", "folder/file.ext" }
		segs := strings.Split(fp, pkg+"/")

		if len(segs) < 2 {
			continue
		}

		rltvPaths = append(rltvPaths, segs[1])
	}

	return rltvPaths, nil
}

// valid files are html files
func isValid(f os.FileInfo) bool {
	return strings.Contains(f.Name(), ".html")
}

// git dir can cause issues if treated like normal dir
func isGitDir(f os.FileInfo) bool {
	return strings.Contains(f.Name(), ".git")
}
