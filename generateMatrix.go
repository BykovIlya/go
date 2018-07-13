package main

import "math/rand"

func generate(n int) [] float64 {
	a := make([]float64, n)
	//var a [n]float64
	for i := 0; i < n; i+=rand.Intn(800) {
		//a = append(a, float64(rand.Intn(10)))
		a[i] = float64(rand.Intn(2))
	}
	for i := 0; i < n; i+=rand.Intn(1500) {
		//a = append(a, float64(rand.Intn(10)))
		a[i] = float64(rand.Intn(5))
	}
	/*
	for i := 0; i < n; i += 2 {
		a[i] = 0;//float64(rand.Intn(1))
	}
	for i := 0; i < n; i += rand.Intn(4) {
		a[i] = 0
	}*/
	return a
}
