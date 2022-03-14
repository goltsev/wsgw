package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	conf := ReadConfigViper()

	service := NewService()

	h, err := Header(conf.URL, conf.Secret, conf.Key, conf.Expires)
	if err != nil {
		return err
	}
	client, err := NewWSClient(service, conf.URL, h)
	if err != nil {
		return err
	}
	defer client.Close()

	server := NewWSServer(service)
	go func() {
		log.Println(server.Serve())
	}()
	return client.Listen()
}
