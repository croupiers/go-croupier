package croupier

import "math/rand"

// Given a distribution of float64 numbers, returns a sample of a given size
func SampleFloat64(size int, dist chan float64) []float64 {
	arr := make([]float64, size)

	for i := 0; i < size; i++ {
		arr = append(arr, <-dist)
	}

	return arr
}

// Maps a distribution with a given fn
func MapDistributionFloat64(dist chan float64, fn func(float64) float64) chan float64 {
	out := make(chan float64)

	go func() {
		for {
			num, ok := <-dist
			if !ok {
				break
			}
			out <- fn(num)
		}
		close(out)
	}()

	return out
}

// Returns a distribution of floats in [0.0, 1.0)
func NewUniform() chan float64 {
	dist := make(chan float64)

	go func() {
		for {
			dist <- rand.Float64()
		}
	}()

	return dist
}

// Returns a distribution of floats in [min, max)
func NewUniformInRange(min, max float64) chan float64 {
	distRange := max - min

	return MapDistributionFloat64(NewUniform(), func(num float64) float64 {
		return num*distRange + min
	})
}

// Returns a new normal distribution, mean 0 and stddev 1
func NewNormal() chan float64 {
	dist := make(chan float64)

	go func() {
		for {
			dist <- rand.NormFloat64()
		}
	}()

	return dist
}

// Returns a new normal distribution with given means and stddev
func NewCustomNormal(mean, stddev float64) chan float64 {
	return MapDistributionFloat64(NewNormal(), func(num float64) float64 {
		return num*stddev + mean
	})
}
