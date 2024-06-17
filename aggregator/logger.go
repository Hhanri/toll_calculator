package main

import (
	"time"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	AggregateDistance(err error, startTime time.Time)
	CalculateInvoice(err error, startTime time.Time, id int, invoice *types.Invoice)
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

func (l *LogrusLogger) CalculateInvoice(err error, startTime time.Time, id int, invoice *types.Invoice) {
	fields := logrus.Fields{
		"took":  time.Since(startTime),
		"error": err,
		"obuId": id,
	}
	if invoice != nil {
		fields["totalDistance"] = invoice.TotalDistance
		fields["totalAmount"] = invoice.TotalAmount
	}
	logrus.WithFields(fields).Info("calculate invoice")
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}
