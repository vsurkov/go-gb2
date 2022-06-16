package main

import "sync"

type RWSet struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewRWSet() *RWSet {
	return &RWSet{
		mm: map[int]struct{}{},
	}
}

func (s *RWSet) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *RWSet) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}
