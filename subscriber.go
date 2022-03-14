package main

import (
	"fmt"
)

type WriterJSON interface {
	WriteJSON(v interface{}) error
}

type Subscriber struct {
	ID   int
	List []string
	Conn WriterJSON
}

func (s Subscriber) Update(n *Notification) error {
	if len(s.List) == 0 {
		if err := s.Conn.WriteJSON(n); err != nil {
			return fmt.Errorf("write json: %w", err)
		}
		return nil
	}
	for _, v := range s.List {
		if n.Symbol == v {
			if err := s.Conn.WriteJSON(n); err != nil {
				return fmt.Errorf("write json: %w", err)
			}
			return nil
		}
	}
	return nil
}
