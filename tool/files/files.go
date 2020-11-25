package files

import (
	"io/ioutil"
	"strings"
)

const pkg = "endito"

func ReadDir(dir string, files []string) ([]string, error) {
	rdir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(rdir) == 1 && strings.Contains(rdir[0].Name(), ".html") {
		return append(files, dir+rdir[0].Name()), nil
	}

	for _, f := range rdir {
		if f.IsDir() && !strings.Contains(f.Name(), ".git") {
			files, err = ReadDir(dir+f.Name()+"/", files)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(f.Name(), ".html") {
			files = append(files, dir+f.Name())
		}
	}

	return files, nil
}

func GetRelativePaths(rltvDir string, absPaths []string) ([]string, error) {
	rltvPaths := make([]string, 0)

	for _, fp := range absPaths {
		segs := strings.Split(fp, pkg+"/")

		if len(segs) < 2 {
			continue
		}

		rltvPaths = append(rltvPaths, segs[1])
	}

	return rltvPaths, nil
}
