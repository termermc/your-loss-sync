package util

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
	"io/fs"
	"path/filepath"
	"slices"
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

	// Strip away any leading or trailing whitespace
	res = strings.TrimSpace(res)

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

var fatNames = []string{
	"fat",
	"fat32",
	"vfat",
	"msdos",
}

// IsPathFat32 returns whether a path appears to be on a FAT32 filesystem.
func IsPathFat32(path string) (bool, error) {
	// Get absolute path to handle relative paths correctly
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Get all partitions
	partitions, err := disk.Partitions(true)
	if err != nil {
		return false, fmt.Errorf("failed to get partitions: %w", err)
	}

	// Find the longest matching partition mount point
	var matchingPartition *disk.PartitionStat
	longestMatch := 0

	for i, partition := range partitions {
		mountPoint := partition.Mountpoint
		if strings.HasPrefix(absPath, mountPoint) && len(mountPoint) > longestMatch {
			longestMatch = len(mountPoint)
			matchingPartition = &partitions[i]
		}
	}

	if matchingPartition == nil {
		return false, fmt.Errorf("no matching partition found for path: %s", path)
	}

	// Check if filesystem is FAT32
	fstype := strings.ToLower(matchingPartition.Fstype)
	return slices.Contains(fatNames, fstype), nil
}
