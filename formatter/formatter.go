package formatter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

func formatFootnote(doc *goquery.Document) {
	selection := doc.Find("a")
	fmt.Println(doc.Html())
	fmt.Println(selection.Length())
}

func LoadFile(file string) {
	var (
		f   []byte
		doc *goquery.Document
		err error
	)

	if f, err = ioutil.ReadFile(file); err != nil {
		fmt.Println(err)
	}

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(f)); err != nil {
		fmt.Println(err)
	}

	formatFootnote(doc)
}

func LoadFolder(folder string) {
}
