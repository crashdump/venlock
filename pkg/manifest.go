package pkg

import (
	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
)

func FindManifests(filename string, dir string) (out []string, err error) {
	ignores := []string{"node_modules"}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if slices.Contains(ignores, filepath.Base(path)) {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			if filepath.Base(path) == filename {
				out = append(out, path)
			}
		}

		return nil
	})
	if err != nil {
		return out, err
	}

	return out, nil
}
