package main

import (
	"flag"
	"fmt"

	"github.com/korprulu/epub3formatter/formatter"
)

var file, folder string

func main() {
	flag.StringVar(&file, "f", "", "XHTML file path")
	flag.StringVar(&folder, "d", "", "XHTML folder path")
	flag.Parse()

	if folder == "" && file == "" {
		fmt.Print("Require file path or folder path.")
	} else if folder != "" {
		formatter.LoadFolder(folder)
	} else {
		formatter.LoadFile(file)
	}
}
