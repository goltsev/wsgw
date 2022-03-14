package main

import (
	"fmt"
	"time"
)

type Instrument struct {
	Action string `json:"action"`
	Data   []struct {
		LastPrice float64   `json:"lastPrice"`
		Symbol    string    `json:"symbol"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"data"`
}

type Notification struct {
	Timestamp time.Time `json:"timestamp"`
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
}

type Request struct {
	Action  string   `json:"action"`
	Symbols []string `json:"symbols"`
}

func BuildNotification(i *Instrument) (*Notification, error) {
	if i == nil || len(i.Data) == 0 {
		return nil, fmt.Errorf("empty data")
	}
	return &Notification{
		Timestamp: i.Data[0].Timestamp,
		Symbol:    i.Data[0].Symbol,
		Price:     i.Data[0].LastPrice,
	}, nil
}
