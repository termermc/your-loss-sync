package util

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// EscapeFilename escapes a filename so that it can be safely used on most filesystems.
func EscapeFilename(filename string) string {
	builder := strings.Builder{}

	// Use unicode replacements for disallowed characters.
	// The replacements are prefixed with an underscore to avoid potential conflicts with existing files.
	for _, r := range filename {
		switch r {
		case '/':
			builder.WriteString("_⧸")
		case '\\':
			builder.WriteString("_⧹")
		case '?':
			builder.WriteString("_？")
		case '%':
			builder.WriteString("_％")
		case '*':
			builder.WriteString("_＊")
		case ':':
			builder.WriteString("_：")
		case '|':
			builder.WriteString("_｜")
		case '"':
			builder.WriteString("_＂")
		case '\'':
			builder.WriteString("_’")
		case '<':
			builder.WriteString("_＜")
		case '>':
			builder.WriteString("_＞")
		}
	}

	res := builder.String()

	// Windows doesn't allow folders to end with a dot
	if strings.HasSuffix(res, ".") {
		res += "_"
	}

	return res
}

// ScanDirFilesRecursive scans a directory recursively and returns a list of all files in the directory.
func ScanDirFilesRecursive(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})

	return files, err
}
