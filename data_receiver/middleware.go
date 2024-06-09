package main

import (
	"time"

	"github.com/Hhanri/toll_calculator/types"
)

type LogMiddleware struct {
	logger Logger
	next   DataProducer
}

func NewLogMiddleware(logger Logger, next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *LogMiddleware) ProduceData(data types.ObuData) error {
	defer func() {
		l.logger.ProduceData(data, time.Now())
	}()
	return l.next.ProduceData(data)
}
