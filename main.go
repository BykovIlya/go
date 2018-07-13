package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main1() {
	/* reading from file */
	events := readingTransactionsFromFile()

	/* get num of transactions */
	fmt.Println("Number of transactions: ", len(events))

	/* make uniq visitors */
	removeDublicatesOfVisitors := makeUniqArrayOfVisitors(events)
	fmt.Println("Number of uniq visitors: ", len(removeDublicatesOfVisitors))
	removeDublicatesOfItems := makeUniqArrayOfItems(events)

	/* make uniq items */
	fmt.Println("Number of uniq items: ",len(removeDublicatesOfItems))
	visitors := make([] Visitor, len(removeDublicatesOfVisitors))

	/* make struct of visitors */
	initVisitors(visitors, removeDublicatesOfVisitors)

	/* add items to each visitor */
	addItemsToVisitor(visitors,events)
	addCountToEachProductOfEachVisitor(visitors)

	/*
		make matrix of sales
		rows - visitors
		columns - items
		the cells show how many items the visitor bought
	*/
	matrixOfSales := makeMatrixOfSales(visitors, removeDublicatesOfVisitors, removeDublicatesOfItems)

	/* init array of sales to get it into CA */
	arrayOfSales := makeArrayOfSales(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems) )

	/* CA algorithm*/
	prefs := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
	//products := removeDublicatesOfItems
	products := make([]string, 0)
	for i := 0; i < len(removeDublicatesOfItems); i++ {
		products = append(products,strconv.Itoa(i))
	}
	//for i := 0; i < /*len(removeDublicatesOfVisitors)*/ 2; i++ {
		fmt.Print("Choose visitor: ")
		var myVisitor string //= "599528"
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		myVisitor = scanner.Text()
		indexOfVisitor := getIndVisitor(visitors, myVisitor)
		if (indexOfVisitor == -1) {
			fmt.Println("Error: visitor doesn't found!")
			os.Exit(-1)
		}
		recommendations, err := GetRecommendations(prefs, /*getIndVisitor(visitors, myVisitor)*/ indexOfVisitor, products)
		if err != nil {
			fmt.Println("WHAT!?")
		}
		//fmt.Print(" Recommended Producs are: ")
		//realProds := make([]string,0)
		//cntOfProds := 0
	if len(recommendations) > 0 {
		fmt.Println("For user ", myVisitor, " recommended products are x with scores y (x --> y)")
		for i := 0; i < len(recommendations); i++ {
			fmt.Println(recommendations[i].product, "-->", recommendations[i].mpRating)
		}
	} else {
		fmt.Println("There are no recommendations for user ", myVisitor)
	}
		fmt.Println()
		/*for i := 0; i < len(prods); i++ {
			if prods[i] != "" {
				//fmt.Print(prods[i], " ")
				realProds = append(realProds, prods[i])
				cntOfProds++
			}
		}
		realScores := make ([]float64, 0)
		//fmt.Print(" with scores: ")
		for i := 0; i < len(scores); i++ {
			if !math.IsNaN(scores[i]) {
				realScores = append(realScores, scores[i])
				//fmt.Print(scores[i], " ")
			}
		}*/
		/*
		if cntOfProds > 0 {
			fmt.Print("For user ", myVisitor)
			fmt.Print(" Recommended Producs are: ", realProds, " with scores: ", realScores)
			fmt.Println()
		}
		*/
	//}
	/*
	   for bayesian filter
	*/
	/* for i := 0; i < len(arrayOfSales); i++ {
	 	if arrayOfSales[i] == 0 {
	 		arrayOfSales[i] = math.NaN()
		}
	 }
	mat := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
	for i := 0; !os.IsExist(err); i++ {
		fmt.Print("Choose item: ")
		var myItem string //= "54058"
		scannerItem := bufio.NewScanner(os.Stdin)
		scannerItem.Scan()
		myItem = scannerItem.Text()
		indexOfItem := getIndItem(removeDublicatesOfItems, myItem)
		if (indexOfItem == -1) {
			fmt.Println("Error: item doesn't found!")
			os.Exit(-1)
		}
		predArray, predictedValue, err:= BayesianFilter(mat, indexOfVisitor, indexOfItem)
		if err != nil {
			fmt.Errorf("Error in BayesianFilter: %v", err)
			os.Exit(-2)
		}
		fmt.Println(predictedValue)
		fmt.Println(predArray)
	}
*/
	/*
		get binary recommendation
	 */
	/*for i := 0; i < len(arrayOfSales); i++ {
		if arrayOfSales[i] > 0 {
			arrayOfSales[i] = 1
		}
	}*/
/*
	binaryPrefs := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
	binProds, binScores, err := GetBinaryRecommendations(binaryPrefs, getIndVisitor(visitors, myVisitor), products)

	if err != nil {
		fmt.Println("WHAT!?")
	}
	fmt.Printf("\nRecommended Products are: %v, with scores: %v", binProds, binScores)
/*
	fmt.Print("Recommended Producs are: ")
	for i := 0; i < len(binProds); i++  {
		if binProds[i] != "" {
			fmt.Print(binProds[i], " ")
		}
	}
	//	fmt.Println(" with scores: ", real_scores)
	//fmt.Print(" with scores: ", scores[i])
	for i := 0; i < len(binScores); i++ {
		if !math.IsNaN(binScores[i]) {
			fmt.Print(binScores[i], " ")
		}
	}
	fmt.Println()
*/
//	foo()

	//fmt.Println(cntBuf)
	//n_factors := 1
	//n_iterations := 1
	//lambda := 0.1
	//Qhat, _ := ALS.Train(prefs, n_factors, n_iterations, lambda)
	//fmt.Println(Qhat)
	/*for i := 0; !os.IsExist(err); i++ {
		fmt.Print("Choose item: ")
		var myItem string //= "54058"
		scannerItem := bufio.NewScanner(os.Stdin)
		scannerItem.Scan()
		myItem = scannerItem.Text()
		indexOfItem := getIndItem(removeDublicatesOfItems, myItem)
		if (indexOfItem == -1) {
			fmt.Println("Error: item doesn't found!")
			os.Exit(-1)
		}
	*/
	//	fmt.Println(ALS.Predict(Qhat, getIndVisitor(visitors, myVisitor), indexOfItem))
	//	fmt.Println(ALS.GetTopNRecommendations(prefs, Qhat, getIndVisitor(visitors, myVisitor), 5, products))
	/*	for i := 0; i < len(arrayOfSales); i++ {
			if (arrayOfSales[i]) > 5 {
				arrayOfSales[i] = 5
			}
		}
		R := ALS.TrainImplicit(prefs, 1, 1, 0.1)
		for i := 0; !os.IsExist(err); i++ {
			fmt.Print("Choose item: ")
			var myItem string //= "54058"
			scannerItem := bufio.NewScanner(os.Stdin)
			scannerItem.Scan()
			myItem = scannerItem.Text()
			indexOfItem := getIndItem(removeDublicatesOfItems, myItem)
			if (indexOfItem == -1) {
				fmt.Println("Error: item doesn't found!")
				os.Exit(-1)
			}
			fmt.Println(ALS.Predict(R, getIndVisitor(visitors, myVisitor), indexOfItem))
		}*/
}

func main()  {
	//test1CF()
	//test2CF()	products := make( []string, 0)
	//	for i := 0; i < 100; i++ {
	//		products = append(products, strconv.Itoa(i))
	//	}
	//test2ALS()
	test4CF()
	test4ALS()
}