package cainiao

import (
	//"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"github.com/hopkings/webcrawler/parser_factory"
)

type CaiNiaoParser struct {
}

func (yp *CaiNiaoParser) Parse(doc *goquery.Document) (*parser_factory.WarehouseInfo, error) {
	//get the location
	whi := parser_factory.CreateWarehouseInfo()
	doc.Find(".goods-detail-left div").Each(func(i int, sel *goquery.Selection) {
		//udata := mahonia.NewDecoder("gbk").ConvertString(sel.Text())
		val, exists := sel.Attr("data-reactid")
		if !exists {
			return
		}
		switch val {
		case "81":
			{
				udata := yp.trimStr(sel.Text())
				if udata == "仓库地址" {
					whi.IsValid = 1
				} else {
					return
				}
			}
		case "82":
			{
				whi.Location = yp.trimStr(sel.Text())
			}
		}
		//elems := strings.Split(udata, "\n")
		//fmt.Printf("%s ", udata)
	})

	/*doc.Find(".cb210_c2").Each(func(i int, sel *goquery.Selection) {
		sel.Find("p").Each(func(ii int, elems *goquery.Selection) {
			udata := mahonia.NewDecoder("gbk").ConvertString(elems.Text())
			udata = strings.Replace(udata, "\n", "", -1)
			fmt.Printf("%s ", udata)
		})
	})

	fmt.Printf("\n")*/

	return whi, nil
}

func (yp *CaiNiaoParser) trimStr(text string) string {
	data := strings.Replace(text, " ", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	return data
}
