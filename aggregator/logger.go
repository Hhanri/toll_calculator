package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	AggregateDistance(err error, startTime time.Time)
}

type LogrusLogger struct{}

func (l *LogrusLogger) AggregateDistance(err error, startTime time.Time) {
	logrus.WithFields(
		logrus.Fields{
			"took":  time.Since(startTime),
			"error": err,
		},
	).Info("aggregate distance")
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}
