package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/list", func(m *tb.Message) {
		res, err := http.Get("https://pure-shore-25232.herokuapp.com/repuestos")
		if err != nil {
			fmt.Println(err)
		}
		data, _ := ioutil.ReadAll(res.Body)
		b.Send(m.Sender, string(data))
	})

	b.Handle("nombre", func(m *tb.Message) {
		b.Send(m.Sender, "You entered "+m.Text)
	})

	b.Start()

}
