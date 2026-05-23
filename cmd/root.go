package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/sainukala225/mdstrip/converter"
	"github.com/sainukala225/mdstrip/utils"
	"github.com/urfave/cli/v3"
)

var RootCmd = &cli.Command{
	Name:    "mdstrip",
	Usage:   "Cli tool to convert markdown files to plain text",
	Version: "v1.0.0",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "filename",
			Aliases: []string{"f"},
			Usage:   "`Filename` to read from",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "`Filename` to write to",
		},
	},
	Arguments: []cli.Argument{
		&cli.StringArg{
			Name: "filename",
		},
		&cli.StringArgs{
			Max: -1,
		},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		var file string
		if cmd.String("filename") != "" {
			file = cmd.String("filename")
		} else if cmd.StringArg("filename") != "" {
			file = cmd.StringArg("filename")
		} else {
			file = ""
		}
		err := run(file, cmd.String("output"))
		return err
	},
}

func run(filename string, output string) (err error) {

	var filePtr *os.File

	if filename != "" {
		if err = utils.IsValidFile(filename, ".md"); err != nil {
			return err
		}
		filePtr, err = os.Open(filename)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := filePtr.Close()
			if err == nil && closeErr != nil {
				err = closeErr
			}
		}()
	} else {
		filePtr = os.Stdin
	}

	plaintext := converter.ConvertMdToTxt(filePtr)
	if output == "" {
		fmt.Println(plaintext)
	} else {
		var outputFilPtr *os.File
		if err = utils.IsValidFile(output, ".txt"); err != nil {
			return err
		}

		outputFilPtr, err = os.Create(output)
		if err != nil {
			return err
		}

		defer func() {
			closeErr := outputFilPtr.Close()
			if err == nil && closeErr != nil {
				err = closeErr
			}
		}()
		_, err = outputFilPtr.Write([]byte(plaintext))
		if err != nil {
			return err
		}

	}
	return err
}
