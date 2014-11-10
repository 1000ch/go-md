package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	args := os.Args[1:]
	length := len(args)

	for i := 0; i < length; i++ {
		data, error := ioutil.ReadFile(args[i])

		if error != nil {
			panic(error)
		}

		unsafe := blackfriday.MarkdownCommon(data)
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		fmt.Print(string(html))
	}
}
