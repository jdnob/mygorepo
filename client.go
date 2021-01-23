package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Data struct {
	Channel string `json:"channel"`
}

type MessageToListen struct {
	Event string `json:"event"`
	Data Data    `json:"data"`
}
var addr = flag.String("addr", "ws.bitstamp.net", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	done := make(chan struct {})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	m := MessageToListen{Event: "bts:subscribe", Data: Data{Channel: "live_order_ltcbtc"}}

	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Println("error converting")
	}
	strjson := string(jsonData)



	for {
		select {
			case <- done:
				return
			case t:= <- ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte(strjson))
				if err != nil {
					log.Println("write: ", err)
					return
				}
				case <- interrupt:
					log.Println("interrupt")

					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						return
					}

					select {
					case <- done:
					case <- time.After(time.Second):
					}
					return
		}
	}
}
