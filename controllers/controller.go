package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func GetQuote() {
	c := colly.NewCollector()
	quotes := []string{}

	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
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
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_AUTH"),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(os.Getenv("MY_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_NUMBER"))
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
