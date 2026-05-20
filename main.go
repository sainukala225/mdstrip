package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sainukala225/mdstrip/converter"
)

func main() {

	filename := flag.String("filename", "", "filename to read")
	output := flag.String("output", "", "filename to write")
	help := flag.Bool("help", false, "print help message")

	flag.Parse()

	if *help {
		printhelpmessage()
		return
	}

	args := flag.Args()
	var file string
	var fileptr *os.File
	var err error

	if *filename != "" || len(args) > 0 {
		if *filename != "" {
			file = *filename
		} else {
			file = args[0]
		}
		if !isvalidfile(file, ".md") {
			return
		}
		fileptr, err = os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			err = fileptr.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()
	} else {
		fileptr = os.Stdin
	}

	plaintext := converter.ConvertMdToTxt(fileptr)
	if *output == "" {
		fmt.Println(plaintext)
	} else {

		if !strings.HasSuffix(*output, ".txt") {
			fmt.Println("output filename must end with .txt")
			return
		}

		outputfileptr, err := os.Create(*output)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer func() {
			err = outputfileptr.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()
		_, err = outputfileptr.Write([]byte(plaintext))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
