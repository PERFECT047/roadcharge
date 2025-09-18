package main

import (
	"time"

	"github.com/perfect047/roadcharge/types"
	"github.com/sirupsen/logrus"
)

type LogMiddlware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddlware{
		next: next,
	}
}

func (m *LogMiddlware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"timeTaken": time.Since(start),
			"err":       err,
			"dist":      dist,
		}).Info("Calculate distance")
	}(time.Now())

	dist, err = m.next.CalculateDistance(data)
	return
}
