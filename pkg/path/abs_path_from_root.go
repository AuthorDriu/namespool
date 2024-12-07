package path

import (
	"os"
	"path/filepath"
)

func FromRoot(path string) string {
	rootPath := os.Getenv("NAMESPOOLROOT")
	newPath := filepath.Join(rootPath, path)
	return newPath
}
