package utils

import (
	"fmt"
	"os"
	"strings"
)

type InvalidFileError struct {
	Path   string
	Reason string
}

func (e InvalidFileError) Error() string {
	return fmt.Sprintf("invalid file %s: %s", e.Path, e.Reason)
}

func IsValidFile(filename, filetype string) error {
	fileinfo, err := os.Stat(filename)

	if err != nil {
		return err
	}

	if fileinfo.IsDir() {
		return InvalidFileError{Path: filename, Reason: "path is a directory"}
	}

	if !strings.HasSuffix(filename, filetype) {
		return InvalidFileError{Path: filename, Reason: "must end with " + filetype}
	}
	return nil
}
