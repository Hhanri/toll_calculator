package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	CalculateDistance(dist float64, startTime time.Time) error
}

type LogrusLogger struct{}

func (l *LogrusLogger) CalculateDistance(dist float64, startTime time.Time) error {
	logrus.WithFields(
		logrus.Fields{
			"took":     time.Since(startTime),
			"distance": dist,
		},
	).Info("calculate distance")
	return nil
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}
