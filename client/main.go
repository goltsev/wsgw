package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	url := "ws://127.0.0.1:8080/"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	go func() {
		conn.WriteJSON(map[string]interface{}{
			"action":  "subscribe",
			"symbols": []string{"MANAUSDT", "XBTZ22", "asd", "qwe", "XBTUSDT"},
		})
		time.Sleep(time.Second * 10)
		conn.WriteJSON(map[string]interface{}{
			"action": "unsubscribe",
		})
	}()
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Println(string(p))
	}
}
