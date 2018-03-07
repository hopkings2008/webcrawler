package parser_factory

import (
	"fmt"
	"net/url"
	"strings"
)

type ParserFactory interface {
	Build() (Parser, error)
}

var factories map[string]ParserFactory

func init() {
	factories = make(map[string]ParserFactory)
}

func RegisterFactory(host string, pf ParserFactory) {
	key := getKey(host)
	factories[key] = pf
}

func BuildFactory(urlstr string) (ParserFactory, error) {
	ul, err := url.Parse(urlstr)
	if err != nil {
		fmt.Printf("failed to parse %s, err: %v\n", urlstr, err)
		return nil, err
	}
	key := getKey(ul.Host)
	factory, ok := factories[key]
	if !ok {
		err = fmt.Errorf("cannot find parser factory for %s", urlstr)
		return nil, err
	}
	return factory, nil
}

func getKey(host string) string {
	h := strings.ToLower(host)
	list := strings.Split(h, ".")
	key := ""
	num := len(list)
	if num >= 2 {
		key = list[num-2]
	} else {
		key = h
	}
	return key
}
