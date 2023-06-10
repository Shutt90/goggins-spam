package main

import (
	"net/http"

	"github.com/gocolly/colly"
	"google.golang.org/appengine"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(startScraper(w, r))
}

func startScraper(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	c := colly.NewCollector()
	c.Appengine(ctx)
	c.Visit("https://www.invajy.com/david-goggins-quotes/")
}
