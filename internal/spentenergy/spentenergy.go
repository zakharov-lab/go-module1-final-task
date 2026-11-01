package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("некорректные данные")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("некорректные данные")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	return (weight * meanSpeed * durationMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	dist := stepLength * float64(steps)
	return dist / mInKm
}
