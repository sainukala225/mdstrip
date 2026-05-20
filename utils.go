package main

import (
	"fmt"
	"os"
	"strings"
)

func printhelpmessage() {
	fmt.Println("Usage: mdstrip [--output <file.txt>] <file.md>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --filename <file.md>    input markdown file")
	fmt.Println("  --output <file.txt>     output text file (default: stdout)")
	fmt.Println("  --help                  print this help message")
}

func isvalidfile(filename, filetype string) bool {
	fileinfo, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if fileinfo.IsDir() {
		fmt.Println("path should be a file")
		return false
	}

	if !strings.HasSuffix(filename, filetype) {
		fmt.Printf("filename must end with %s", filetype)
		return false
	}
	return true
}
