package main

import (
	"time"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	ProduceData(data types.ObuData, startDate time.Time) error
}

type LogrusLogger struct{}

func (l *LogrusLogger) ProduceData(data types.ObuData, startDate time.Time) error {
	logrus.WithFields(
		logrus.Fields{
			"obuId": data.ObuId,
			"lat":   data.Geo.Lat,
			"lng":   data.Geo.Lng,
			"took":  time.Since(startDate),
		},
	).Info("producing")
	return nil
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}
