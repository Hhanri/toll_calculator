package main

import (
	"math"

	"github.com/Hhanri/toll_calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.ObuData) (float64, error)
}

type CalculatorService struct {
	points map[int]types.Geo
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		points: make(map[int]types.Geo),
	}
}

func (s *CalculatorService) CalculateDistance(data types.ObuData) (float64, error) {
	defer s.saveLastPoint(data)
	prev, ok := s.points[data.ObuId]
	if !ok {
		return 0.0, nil
	}
	distance := calculateDistance(data.Geo.Lng, data.Geo.Lat, prev.Lng, prev.Lat)
	return distance, nil
}

func (s *CalculatorService) saveLastPoint(data types.ObuData) {
	s.points[data.ObuId] = data.Geo
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	a := math.Pow(x2-x1, 2)
	b := math.Pow(y2-y1, 2)
	return math.Sqrt(a + b)
}
