package main

import "github.com/Hhanri/toll_calculator/types"

type Storer interface {
	Insert(types.Distance) error
}

type MemoryStore struct {
	data map[int64]float64
}

func NewMemoryStore() Storer {
	return &MemoryStore{
		data: make(map[int64]float64),
	}
}

func (m *MemoryStore) Insert(distance types.Distance) error {
	m.data[int64(distance.ObuId)] += (distance.Value)
	return nil
}
