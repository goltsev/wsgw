package main

type Service struct {
	subs map[int]*Subscriber
	next int
}

func NewService() *Service {
	return &Service{
		subs: make(map[int]*Subscriber),
	}
}

func (s *Service) Subscribe(sub *Subscriber) {
	s.next++
	sub.ID = s.next
	s.subs[sub.ID] = sub
}

func (s *Service) Unsubscribe(sub *Subscriber) {
	delete(s.subs, sub.ID)
}

func (s *Service) Notify(n *Notification) {
	for _, sub := range s.subs {
		if err := sub.Update(n); err != nil {
			s.Unsubscribe(sub)
		}
	}
}
