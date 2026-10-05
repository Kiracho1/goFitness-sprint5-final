package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be greater than 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}
	// Рассчитываем количество калорий
	result := (weight * MeanSpeed(steps, height, duration) * duration.Hours()) * walkingCaloriesCoefficient

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be greater than 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}
	// Рассчитываем количество калорий
	result := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return result, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	lengthStep := height * stepLengthCoefficient
	numSteps := lengthStep * float64(steps)
	return numSteps / mInKm
}
