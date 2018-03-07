package parser_factory

import (
	"github.com/PuerkitoBio/goquery"
)

type Parser interface {
	Parse(doc *goquery.Document) (*WarehouseInfo, error)
}

func BuildParser(url string) (Parser, error) {
	factory, err := BuildFactory(url)
	if err != nil {
		return nil, err
	}
	parser, err := factory.Build()
	return parser, err
}
