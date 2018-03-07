package main

import (
	"fmt"
	"runtime"

	"github.com/gocolly/colly"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/webloop"
)

func main() {
	gtk.Init(nil)
	go func() {
		runtime.LockOSThread()
		gtk.Main()
	}()
	ctx := webloop.New()
	view := ctx.NewView()
	defer view.Close()
	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Println(link)
		// Visit link found on page
		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", "cna=2BRiEiRjRBECAQF3NiqbBj//; hng=CN%7Czh-CN%7CCNY%7C156; t=d4bbb940b99d788029cfb1a3ce6291f6; tracknick=hopkings2005; _tb_token_=757e137e7e378; cookie2=1ff9218305c575db592b8ce12cb42857; ctoken=xBjO8jnuddE0nL2un8Z3S0yG; csrf=Q0ARuaDm-1lW-U3oorErbv03VbUn6wvuX54w; csrf.sig=_noZZ2FqOUQKtueNTWwYmS5JzmFz7VtpN59nwguLFMI; isg=BODgX1-Rrr1Y6BJrY-t9d1o4sehyQQx9pt4WIFrxrPuOVYB_AvmUQ7ZH6f1VfnyL")
	})

	c.OnResponse(func(r *colly.Response) {
		view.Load(string(r.Body), "https://market.c.cainiao.com")
		err := view.Wait()
		if err != nil {
			fmt.Printf("Failed to load URL: %s", err)
		}
		res, err := view.EvaluateJavaScript("document.documentElement.outerHTML")
		if err != nil {
			fmt.Printf("Failed to run JavaScript: %s", err)
		}
		gotbody, _ := res.(string)
		r.Body = []byte(gotbody)
		fmt.Printf("%s\n", string(r.Body))
	})

	// Start scraping on https://en.wikipedia.org
	//c.Visit("https://en.wikipedia.org/")
	c.Visit("https://market.c.cainiao.com/search/?q=&pm=2")
}
