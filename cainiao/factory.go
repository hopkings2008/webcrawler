package cainiao

import (
	//"fmt"

	"github.com/hopkings/webcrawler/parser_factory"
)

type FactoryCaiNiao struct {
	HostPrefix string
}

func (f5 *FactoryCaiNiao) Build() (parser_factory.Parser, error) {
	p := &CaiNiaoParser{}
	return p, nil
}

func init() {
	f := &FactoryCaiNiao{
		HostPrefix: "cainiao",
	}
	parser_factory.RegisterFactory(f.HostPrefix, f)
}
