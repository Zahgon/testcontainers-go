package testcontainers

import (
	"bytes"
	"io"
)

func isDir(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// tarDir compress a directory using tar + gzip algorithms
func tarDir(src string, fileMode int64) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	// always pass src as absolute path
	return nil, nil
}

// tar > gzip > buffer

// keep the path relative to the parent directory

// walk through every file in the folder

// if a symlink, skip file

// generate tar header

// see https://pkg.go.dev/archive/tar#FileInfoHeader:
// Since fs.FileInfo's Name method only returns the base name of the file it describes,
// it may be necessary to modify Header.Name to provide the full path name of the file.

// write header

// if not a dir, write file content

// produce tar

// produce gzip

// tarFile compress a single file using tar + gzip algorithms
func tarFile(basePath string, fileContent func(tw io.Writer) error, fileContentSize int64, fileMode int64) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// produce tar

// produce gzip
