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
	// 1: location, 2: class, 3: serviceclass, 4: serviceregion.
	infoType := 0
	doc.Find(".goods-detail-left div").Each(func(i int, sel *goquery.Selection) {
		data := yp.trimStr(sel.Text())
		if strings.Contains(data, "仓库地址") {
			whi.IsValid = 1
			infoType = 1
			return
		}
		if strings.Contains(data, "服务行业") {
			infoType = 2
			return
		}
		if strings.Contains(data, "服务类型") {
			infoType = 3
			return
		}
		if strings.Contains(data, "服务范围") {
			infoType = 4
			return
		}

		switch infoType {
		case 1:
			{
				whi.Location = yp.trimStr(sel.Text())
			}
		case 2:
			{
				whi.Class = yp.trimStr(sel.Text())
			}
		case 3:
			{
				whi.ServiceClass = yp.trimStr(sel.Text())
			}
		case 4:
			{
				whi.ServiceRegion = yp.trimStr(sel.Text())
			}
		}
	})

	doc.Find(".basis-info-content div").Each(func(i int, sel *goquery.Selection) {
		/*val, exists := sel.Attr("data-reactid")
		if !exists {
			return
		}*/
		name := yp.trimStr(sel.Text())
		if !strings.Contains(name, "建筑面积") {
			return
		}
		sel.Find("span").Each(func(i int, spn *goquery.Selection) {
			whi.Square = yp.trimStr(spn.Text())
		})
	})

	doc.Find(".basis-function-content div").Each(func(i int, sel *goquery.Selection) {
		name := yp.trimStr(sel.Text())
		if !strings.Contains(name, "地面材料") {
			return
		}
		sel.Find("span").Each(func(i int, spn *goquery.Selection) {
			whi.FloorInfo = yp.trimStr(spn.Text())
		})
	})

	doc.Find(".basis-fire-content div").Each(func(i int, sel *goquery.Selection) {
		sel.Find("span").Each(func(i int, spn *goquery.Selection) {
			whi.FireSystem = yp.trimStr(spn.Text())
		})
	})

	return whi, nil
}

func (yp *CaiNiaoParser) trimStr(text string) string {
	data := strings.Replace(text, " ", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	return data
}
