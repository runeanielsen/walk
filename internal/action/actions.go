package action

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FilterOut(path, ext string, minSize int64, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}
	return false
}

func ListFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}