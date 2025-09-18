package main

import (
	"math"

	"github.com/perfect047/roadcharge/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type OBUCoordinate struct {
	Longitude float64
	Latitude  float64
}

type CalculatorService struct {
	prevPoint *OBUCoordinate
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		prevPoint: nil,
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	prevPoint := s.prevPoint
	point := &OBUCoordinate{
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
	}
	s.prevPoint = point

	if prevPoint == nil {
		return 0.0, nil
	}

	return calculateDistance(
			prevPoint.Longitude,
			prevPoint.Latitude,
			data.Latitude,
			data.Longitude),
		nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
