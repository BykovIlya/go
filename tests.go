package main

import "fmt"

func test1 () {
	prefs := MakeRatingMatrix([]float64{
		2, 3, 4, 1, 5,
		3, 0, 3, 3, 0,
		4, 4, 1, 2, 3,
		2, 4, 0, 3, 4,
		3, 1, 3, 0, 4}, 5, 5)
	products := []string{"Spiderman", "Big Momma's House", "Vanilla Sky", "Pacific Rim", "The Mask"}
	prods, scores, err := GetRecommendations(prefs, 1, products)
	if err != nil {
		fmt.Println("WHAT!?")
	}
	fmt.Printf("\nRecommended Products are: %v, with scores: %v", prods, scores)
}

func test2 () {
	prefs := MakeRatingMatrix([] float64 {
		0,1,3,0,2,1,0,3,2,1,

		1,0,0,0,0,0,0,0,2,3,

		2,3,5,6,3,2,5,0,0,0,

		4,2,5,4,0,0,0,0,0,1,

		5,1,2,1,2,1,0,3,2,1,

		4,3,0,0,0,5,0,0,0,3,

		3,3,4,5,3,3,4,0,0,4,

		1,1,1,1,1,1,1,1,1,1,

		1,0,0,0,0,0,0,0,0,0,
	},9,10)
	products := []string{"1","2","3","4","5","6","7","8","9","10"}
	prods, scores, err := GetRecommendations(prefs, 0, products)
	if err != nil {
		fmt.Println("WHAT!?")
	}
	fmt.Printf("\nRecommended Products are: %v, with scores: %v", prods, scores)
}
