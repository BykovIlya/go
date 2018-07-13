package main

import "math/rand"

func generate(n int) [] float64 {
	a := make([]float64, 0)
	for i := 0; i < n; i++ {
		a = append(a, float64(rand.Intn(5)))
	}
	for i := 0; i < n; i += 2 {
		a[i] = float64(rand.Intn(1))
	}
	return a
}
