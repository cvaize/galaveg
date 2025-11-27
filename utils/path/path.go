package path

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//https://github.com/golang/go/blob/master/src/cmd/go/internal/base/path.go
//https://github.com/golang/go/blob/9e3b1d53a012e98cfd02de2de8b1bd53522464d4/src/cmd/go/internal/modload/init.go#L1504C1-L1522C2

// Cwd returns the current working directory.
func Cwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func FindModuleRoot(dir string) (roots string) {
	if dir == "" {
		panic("dir not set")
	}
	dir = filepath.Clean(dir)

	// Look for enclosing go.mod.
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
	}
	return ""
}

func CollectFilepathBySuffix(dir, suffix string) ([]string, error) {
	suffix = strings.ToLower(suffix)
	var paths []string
	err := filepath.WalkDir(dir, func(fullPath string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), suffix) {
			paths = append(paths, fullPath)
		}
		return nil
	})
	return paths, err
}
