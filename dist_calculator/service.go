package main

import (
	"math"

	"github.com/perfect047/roadcharge/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	points [][]float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		points: make([][]float64, 0),
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	s.points = append(s.points, []float64{data.Latitude, data.Latitude})
	if len(s.points) < 2 {
		return 0.0, nil
	}
	prevPoints := s.points[len(s.points)-2]
	return calculateDistance(
			prevPoints[0],
			prevPoints[1],
			data.Latitude,
			data.Longitude),
		nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
