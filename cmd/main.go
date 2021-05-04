package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shmsr/tokr"
)

func main() {
	doc := flag.String("doc", "", "path of the target file for analysis")
	keywords := flag.String("keywords", "", "path of the keyword file")
	flag.Parse()

	err := tokr.Run(*doc, *keywords)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
