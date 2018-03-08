package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/hopkings/webcrawler/cainiao"
	"github.com/hopkings/webcrawler/parser_factory"
)

func main() {
	testHtml := "./test.html"
	//warehouseHandle, err := os.Create("./warehouseinfo.txt")
	warehouseHandle, err := os.OpenFile("./warehouseinfo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to create warehouseinfo.txt, err: %v\n", err)
		return
	}
	defer warehouseHandle.Close()

	// get the parser from parser_factory.
	pf, err := parser_factory.BuildFactory("https://market.c.cainiao.com/detail/cpfw?id=1209&mid=4398128526840")
	if err != nil {
		fmt.Printf("failed to build factory, err: %v\n", err)
		return
	}
	// create the parser.
	parser, err := pf.Build()
	if err != nil {
		fmt.Printf("failed to get the parser, err: %v\n", err)
		return
	}
	// got the document by using goquery.
	content, err := ioutil.ReadFile(testHtml)
	if err != nil {
		fmt.Printf("failed to read test.html, err: %v\n", err)
		return
	}
	bodyReader := bytes.NewBuffer(content)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		fmt.Printf("failed to create the document from %s, err: %v\n", testHtml, err)
	}
	whi, err := parser.Parse(doc)
	if err != nil {
		fmt.Printf("failed to parse the document from %s, err: %v\n", testHtml, err)
	}
	if whi.IsValid != 1 {
		return
	}
	warehouseHandle.WriteString(whi.String() + "\n")
	warehouseHandle.Sync()
}
