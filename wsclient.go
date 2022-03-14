package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var ()

type WSClient struct {
	service *Service
	conn    *websocket.Conn
}

func NewWSClient(service *Service, wsurl string, header http.Header) (*WSClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(wsurl, header)
	if err != nil {
		return nil, fmt.Errorf("websocket dial: %w", err)
	}

	return &WSClient{
		service: service,
		conn:    conn,
	}, nil
}

func (c *WSClient) Close() error {
	return c.conn.Close()
}

func (c *WSClient) Listen() error {
	i := new(Instrument)
	for {
		if err := c.Read(i); err != nil {
			return fmt.Errorf("error reading data: %w", err)
		}
		n, err := BuildNotification(i)
		if err != nil {
			log.Printf("error building notification: %v", err.Error())
		}
		c.service.Notify(n)
	}
}

func (c *WSClient) Read(v interface{}) error {
	if err := c.conn.ReadJSON(&v); err != nil {
		return err
	}
	return nil
}

func (c *WSClient) Write(i interface{}) error {
	if err := c.conn.WriteJSON(i); err != nil {
		return fmt.Errorf("write json: %w", err)
	}
	return nil
}

func Header(wsurl, secret, key string, expires time.Time) (http.Header, error) {
	urlo, err := url.Parse(wsurl)
	if err != nil {
		return nil, fmt.Errorf("parsing err: %w", err)
	}
	return header(expires, key, signatureString(secret, message("GET", urlo.Path, expires, nil))), nil
}

func toUnix(date time.Time) string {
	return fmt.Sprintf("%d", date.Unix())
}

func message(verb string, path string, expires time.Time, data []byte) string {
	buf := bytes.NewBuffer(make([]byte, 0, 20))
	buf.WriteString(verb)
	buf.WriteString(path)
	buf.WriteString(toUnix(expires))
	buf.Write(data)
	return buf.String()
}

func signature(secret []byte, message []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

func signatureString(secret string, message string) string {
	return signature([]byte(secret), []byte(message))
}

func header(expires time.Time, pubkey string, signature string) http.Header {
	return http.Header{
		"api-expires":   []string{toUnix(expires)},
		"api-key":       []string{pubkey},
		"api-signature": []string{signature},
	}
}
