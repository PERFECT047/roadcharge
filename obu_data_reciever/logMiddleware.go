package main

import (
	"time"

	"github.com/perfect047/roadcharge/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuId":     data.OBUID,
			"longitude": data.Longitude,
			"latitude":  data.Latitude,
			"timeTaken": time.Since(start),
		}).Info("Producing to Kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
