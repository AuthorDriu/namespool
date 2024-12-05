package path

import (
	"os"
	"path/filepath"
)

func FromRoot(path string) string {
	rootPath := os.Getenv("NAMESPOOL_ROOT")
	newPath := filepath.Join(rootPath, path)
	return newPath
}
