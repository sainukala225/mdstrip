package main

import (
	"context"
	"log"
	"os"

	"github.com/sainukala225/mdstrip/cmd"
)

func main() {

	if err := cmd.RootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
