package main

import (
	"context"
	"fmt"
	"sync"
)

type WriterJSON interface {
	WriteJSON(v interface{}) error
}

type Subscriber struct {
	id   int
	conn WriterJSON
	m    *sync.RWMutex
	ch   chan *Notification

	List  []string
	Error error
}

func NewSubscriber(ctx context.Context, conn WriterJSON) *Subscriber {
	sub := &Subscriber{
		conn: conn,
		m:    new(sync.RWMutex),
		ch:   make(chan *Notification),
	}
	go sub.Run(ctx)
	return sub
}

func (s *Subscriber) Run(ctx context.Context) {
	defer close(s.ch)
	for {
		select {
		case <-ctx.Done():
			return
		case n := <-s.ch:
			if err := s.write(n); err != nil {
				s.Error = err
				return
			}
		}
	}
}

func (s *Subscriber) write(n *Notification) error {
	if len(s.List) == 0 {
		if err := s.conn.WriteJSON(n); err != nil {
			return fmt.Errorf("write json: %w", err)
		}
		return nil
	}
	for _, v := range s.List {
		if n.Symbol == v {
			if err := s.conn.WriteJSON(n); err != nil {
				return fmt.Errorf("write json: %w", err)
			}
			return nil
		}
	}
	return nil
}

func (s *Subscriber) Update(n *Notification) {
	s.ch <- n
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
