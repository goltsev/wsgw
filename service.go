package main

import (
	"sync"
)

type Service struct {
	subs  map[int]*Subscriber
	next  int
	mutex sync.Mutex
}

func NewService() *Service {
	return &Service{
		subs: make(map[int]*Subscriber),
	}
}

func (s *Service) Subscribe(sub *Subscriber) {
	s.mutex.Lock()
	s.next++
	sub.SetID(s.next)
	s.subs[sub.ID()] = sub
	s.mutex.Unlock()
}

func (s *Service) Unsubscribe(sub *Subscriber) {
	s.mutex.Lock()
	delete(s.subs, sub.ID())
	s.mutex.Unlock()
}

func (s *Service) Notify(n *Notification) {
	m := make(map[int]*Subscriber)
	s.mutex.Lock()
	for k, v := range s.subs {
		m[k] = v
	}
	s.mutex.Unlock()
	for _, sub := range m {
		sub.Update(n)
	}
}
