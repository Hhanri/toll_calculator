package main

import (
	"fmt"

	"github.com/Hhanri/toll_calculator/types"
)

const basePrice = 3.14

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("Processing and inserting distance in the storage:", distance)
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(id int) (*types.Invoice, error) {
	dist, err := i.store.Get(id)
	if err != nil {
		return nil, err
	}
	inv := &types.Invoice{
		ObuId:         id,
		TotalDistance: dist,
		TotalAmount:   dist * basePrice,
	}
	return inv, nil
}
