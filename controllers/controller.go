package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	quotes := []string{}

	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
		c.OnHTML("p", func(e *colly.HTMLElement) {
			_, err := strconv.Atoi(string(e.Text[0]))
			if err != nil {
				return
			}
			if strings.Contains(e.Text, `“`) && strings.Contains(e.Text, `David Goggins`) && strings.Contains(e.Text, `.`) {
				quotes = append(quotes, e.Text)
				fmt.Println(e.Text)
			}
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit("https://thestrive.co/best-david-goggins-quotes-for-motivation/")

	rand.Seed(time.Now().Unix())
	sendMessage(quotes[rand.Intn(len(quotes))])
}

func sendMessage(msg string) {
	// TODO: twilio or facebook api send quote to person
}
