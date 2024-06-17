package main

import (
	"time"

	"github.com/Hhanri/toll_calculator/types"
)

type LogMiddleware struct {
	logger Logger
	next   Aggregator
}

func NewLogMiddleware(logger Logger, next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		logger: logger,
		next:   next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func() {
		m.logger.AggregateDistance(err, time.Now())
	}()
	err = m.next.AggregateDistance(distance)
	return
}

func (m *LogMiddleware) CalculateInvoice(id int) (invoice *types.Invoice, err error) {
	defer func() {
		m.logger.CalculateInvoice(err, time.Now(), id, invoice)
	}()
	invoice, err = m.next.CalculateInvoice(id)
	return
}
