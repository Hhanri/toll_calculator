package main

import (
	"fmt"

	"github.com/Hhanri/toll_calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.ObuData) (float64, error)
}

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.ObuData) (float64, error) {
	fmt.Println("calculating the distance")

	return 0.0, nil
}
