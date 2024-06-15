package main

import (
	"time"

	"github.com/Hhanri/toll_calculator/types"
)

type LogMiddleware struct {
	logger Logger
	next   CalculatorServicer
}

func NewLogMiddleware(logger Logger, next CalculatorServicer) *LogMiddleware {
	return &LogMiddleware{
		logger: logger,
		next:   next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.ObuData) (dist float64, err error) {
	defer func() {
		m.logger.CalculateDistance(dist, time.Now())
	}()
	dist, err = m.next.CalculateDistance(data)
	return
}
