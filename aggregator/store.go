package main

import (
	"fmt"

	"github.com/Hhanri/toll_calculator/types"
)

type Storer interface {
	Insert(types.Distance) error
	Get(int64) (float64, error)
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

func (m *MemoryStore) Get(id int64) (float64, error) {
	dist, ok := m.data[id]
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obu [%d]", id)
	}
	return dist, nil
}
