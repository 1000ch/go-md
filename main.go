package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

func main() {

	// define command line flag
	o := flag.String("o", "", "Specify output file path")

	// parse arguments
	flag.Parse()

	// get arguments
	args := flag.Args()
	len := len(args)

	if len == 0 {
		// there is no markdown input
		panic("Specify input markdown")
	} else if len != 1 {
		// there are multiple markdown files
		panic("Cannot specify multiple input")
	}

	data, error := ioutil.ReadFile(args[0])

	if error != nil {
		panic(error)
	}

	unsafe := blackfriday.MarkdownCommon(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	if *o != "" {
		ioutil.WriteFile(*o, unsafe, 0644)
	} else {
		fmt.Print(string(html))
	}
}
