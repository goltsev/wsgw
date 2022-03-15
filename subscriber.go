package main

import (
	"fmt"
	"sync"
)

type WriterJSON interface {
	WriteJSON(v interface{}) error
}

type Subscriber struct {
	id   int
	List []string
	Conn WriterJSON
	m    sync.RWMutex
}

func (s *Subscriber) Update(n *Notification) error {
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

func (s *Subscriber) SetID(id int) {
	s.m.Lock()
	s.id = id
	s.m.Unlock()
}

func (s *Subscriber) ID() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.id
}
