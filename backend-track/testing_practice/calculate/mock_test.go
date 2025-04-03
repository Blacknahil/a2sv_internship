package main

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// shape services mocks the ShapeService interface
type ShapeServiceMock struct {
	mock.Mock
}

func (m *ShapeServiceMock) CalculateArea(radius float64) float64 {
	fmt.Println("Mocked area calculation function")
	fmt.Printf("Radius passed in: %f\n", radius)
	args := m.Called(radius)
	return args.Get(0).(float64)
}

func (m *ShapeServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// CircleServices represents a service for a circle-related claculations
type CircleService struct {
	shapeService ShapeService
}

//  using the provided radius

func (cs CircleServices) CalculateCircleAreas(radius float64) float64 {
	return cs.ShapeService
}
