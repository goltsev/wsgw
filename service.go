package main

import "sync"

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
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, sub := range s.subs {
		if err := sub.Update(n); err != nil {
			s.Unsubscribe(sub)
		}
	}
}
