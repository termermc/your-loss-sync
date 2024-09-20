package util

import (
	"io/fs"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// EscapeFilename escapes a filename so that it can be safely used on most filesystems.
// Do not use on paths, only use on filenames.
// The filenames will be truncated to 255 characters if they are longer.
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
		default:
			builder.WriteRune(r)
		}
	}

	res := builder.String()

	// Windows doesn't allow folders to end with a dot
	if strings.HasSuffix(res, ".") {
		res += "_"
	}

	// Truncate to 255 characters if the filename is longer.
	// Try to preserve the extension if possible.
	if runeCount := utf8.RuneCountInString(res); runeCount > 255 {
		ext := filepath.Ext(res)
		if ext == "" {
			res = res[:255]
		} else {
			withoutExt := res[:len(res)-len(ext)]
			res = withoutExt[:255-len(ext)] + ext
		}
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
