package main

import (
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"fmt"
	"math"
)


/**
	the struct of events
 */
type Events struct {
//	timestamp string /*int64*/
	visitorid string /*int64*/
//	event_ string /*object*/
	itemid string /*int64*/
//	transactionid string /*float64*/

}

/**
	the struct of items
 */
type Items struct{
	itemid_string string
	itemid_count float64
}

/**
	the struct of visitors
 */
type Visitor struct {
	visitorid_string string
	items [] Items
}

/**
	get index of visitor
 */
func getIndVisitor (visitor [] Visitor, finder string) int {
	for i := 0; i < len(visitor); i++ {
		if visitor[i].visitorid_string == finder {
			return i
		}
	}
	return -1
}

/**
	get index of item
 */
func getIndItem (items [] string, finder string) int {
	for i := 0; i < len(items); i++ {
		if items[i] == finder {
			return i
		}
	}
	return -1
}

/**
	set the field visitorid_strnig of the structure Visitor to the value of unique visitors from the array buffer
 */
func initVisitors (visitor []Visitor, buffer [] string) {
	for i := 0; i < len(buffer); i++ {
		visitor[i].visitorid_string = buffer[i]
	}
}

/**
	set each visitor an array of items
 */
func addItemsToVisitor (visitor []Visitor, events []*Events){
	for i := 0; i < len(visitor); i++ {
		for j := 0; j < len(events); j++ {
			if visitor[i].visitorid_string == events[j].visitorid {
				visitor[i].items = append(visitor[i].items, Items{
					itemid_string: events[j].itemid,
					itemid_count: 1,
				})
			}
		}
	}
}

/*
func findVisitorInEvents(events []*Events, finder string) int {
	for i := 0; i < len(events); i++ {
		if events[i].visitorid == finder {
			return i
		}
	}
	return -1
}

func findItemsInEvents (events []*Events, finder string) int {
	for i := 0; i < len(events); i++ {
		if events[i].itemid == finder {
			return i
		}
	}
	return -1
}
*/

/**
	remove dublicates from visitors and itmes for make uniq arrays
 */
func removeDuplicates(array [] string) [] string{
	if len(array) == 1 || len(array) == 0 {
		return array
	}
	unique := 1
	for i := 1; i < len(array); i++{
		if array[i] != array[i - 1] {
			unique++;
		}
	}
	result := make([]string, unique)
	k := 0;
	if len(result) > 0 {
		result[k] = array[0]
		k++
	}
	for i := 1; i < len(array); i++ {
		if array[i] != array[i - 1] {
			result[k] = array[i];
			k++
		}
	}
	return result;
}

/**
	convert matrix to array
 */

 func toArray (matrix [][] float64, n int, m int, array [] float64) []float64 {
	 for i := 0; i < n ; i++  {
		 for j := 0; j < m; j++ {
			array = append(array, matrix[i][j])
		 }
	 }
	 return array
 }

/*
func initCountToResult (item []Items) {
	for i := 0; i < len(item); i++ {
		item[i].itemid_count = 1
	}
}
*/

/**
	find count of each items in array of items for each visitor
 */
func findCount (item []Items) [] Items{
	buffer := make( [] Items, 0);
	var prev string
	for i := 0; i < len(item); i++ {
		if (item[i].itemid_string != prev) {
			buffer = append(buffer, Items {
				item[i].itemid_string,
				1,
			})
		} else {
			buffer[len(buffer) - 1].itemid_count++
		}
		prev = item[i].itemid_string
	}
	return buffer
}

/*
func removeDuplicatesInItems(item []Items) []Items {
	unique := 1
	for i := 1; i < len(item); i++{
		if item[i] != item[i - 1] {
			unique++;
		}
	}
	result := make([]Items, unique)
	initCountToResult(result)
	k := 0;
	if len(result) > 0 {
		result[k].itemid_string = item[0].itemid_string
		k++
	}
	for i := 1; i < len(item); i++ {
		if item[i].itemid_string != item[i - 1].itemid_string {
			result[k].itemid_string = item[i].itemid_string;
			result[k].itemid_count++
			k++
		}
	}
	return result;
}
*/
/**
find element from array
 */
/*func find(buf []*Events, events []*Events, visitor []*Visitor) {
	for i := 1; i < len(buf); i++ {
		if buf[i].visitorid == buf[0].visitorid {
			resultInd := findInEvents(events, buf[i].visitorid)
			var itemsBuf []*Items
			itemsBuf = append(itemsBuf,&Items{
				itemid_string: buf[i].itemid,
				itemid_int: resultInd,
			})
			visitor = append(visitor, &Visitor{
				buf[i].visitorid,
				resultInd,
				itemsBuf,
			})
			remove(buf, buf[i])
		}
	}
	remove(buf,buf[0])
}
*/


/**
delete elem from array
 */
 /*
func remove(list []*Events, item *Events) []*Events {
	for i, v := range list {
		if v == item {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = nil
			list = list[:len(list)-1]
		}
	}
	return list
}
*/


/*
	remove unnecessary elements from score array
 */
func optimizeScores(scores [] float64, good [] float64) []float64{
	for i := 0; i < len(scores); i++ {
		if !math.IsNaN(scores[i]) {
			good = append(good, scores[i])
		}
	}
	return good
}

//func optimizeProds (prods [] )
/*
 main function
 */
func main() {
//	csvFile, _ := os.Open("events_example.csv")
	csvFile, _ := os.Open("events.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []*Events
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[2] == "transaction" {
				events = append(events, &Events{
			//	timestamp:     line[0],
				visitorid:     line[1],
			//	event_ :       line[2],
				itemid:        line[3],
			//	transactionid: line[4],
			})
		}
	}
	fmt.Println("Number of transactions: ", len(events))
	// for visitors
	//t0_0 := time.Now()
	bufOfVisitors := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfVisitors[i] = events[i].visitorid
	}
	removeDublicatesOfVisitors := make ([] string, len(events))
	sort.Strings(bufOfVisitors)
	removeDublicatesOfVisitors = removeDuplicates(bufOfVisitors)
	fmt.Println("Number of uniq visitors: ", len(removeDublicatesOfVisitors))
	bufOfItems := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfItems[i] = events[i].itemid
	}
	sort.Strings(bufOfItems)
	removeDublicatesOfItems := make ([] string, len(events))
	removeDublicatesOfItems = removeDuplicates(bufOfItems)
	fmt.Println("Number of uniq items: ",len(removeDublicatesOfItems))
	visitors := make([] Visitor, len(removeDublicatesOfVisitors))
	initVisitors(visitors, removeDublicatesOfVisitors)
	addItemsToVisitor(visitors,events)


	for i := 0; i < len(visitors); i++  {
		sort.Slice(visitors[i].items, func(j, k int) bool { return visitors[i].items[j].itemid_string < visitors[i].items[k].itemid_string })
	}


	for i := 0; i < len(visitors); i++ {
		visitors[i].items = findCount(visitors[i].items)
	}

	/*
		output visitors
	 */
/*	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		fmt.Println(visitors[i])
	}
*/
	/*
		output visitors
	 */
	//fmt.Println("Visitors: ", removeDublicatesOfVisitors)

	/*
		output items
	 */
	//fmt.Println("Items: ", removeDublicatesOfItems)
	/*
		init matrix
	 */
	matrixOfSales := make([][] float64, len(removeDublicatesOfVisitors))
	for i := 0; i < len(removeDublicatesOfVisitors); i++  {
		matrixOfSales[i] = make([] float64, len(removeDublicatesOfItems))
	}

	/*
	make matrix
	 */
	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		for j := 0; j < len(visitors[i].items); j++ {
			//if visitors[i].items[j].itemid_count > 0 {
				matrixOfSales[i][getIndItem(removeDublicatesOfItems,visitors[i].items[j].itemid_string)] = visitors[i].items[j].itemid_count;
			//}
		}
	}
	/*
	init array of sales to get it into CA
	 */
	 arrayOfSales := make ([]float64, 0/*len (removeDublicatesOfVisitors) * len(removeDublicatesOfItems)*/)
	 arrayOfSales = toArray(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems), arrayOfSales)
	 /*
	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	 */
	/**
	output matrix
	 */
	/*for i:=0; i < len(removeDublicatesOfVisitors); i++ {
		for j:=0; j < len(removeDublicatesOfItems); j++ {
			//matrixOfSales[i][j] = visitors[i].items[j].itemid_count;
			fmt.Print(matrixOfSales[i][j], " ")
		}
		fmt.Println()
	}*/
	//fmt.Println()
	/*
		output array
	 */
	 /*
	for i := 0; i < len(removeDublicatesOfItems) * len(removeDublicatesOfVisitors); i++ {
		fmt.Print(arrayOfSales[i], " ")
	}
	 */
	//fmt.Println()
	//t0_1 := time.Now()
	//fmt.Println("Time spent preparing data: ", t0_1.Sub(t0_0))
	prefs := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
	products := removeDublicatesOfItems
	fmt.Print("Choose visitor: ")
	/*for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		fmt.Print(removeDublicatesOfVisitors[i] " ")
	}*/
	//fmt.Println()
	var myVisitor string //= "599528"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	myVisitor = scanner.Text()
	//myVisitor = "599528"
	//fmt.Println(getIndVisitor(visitors, myVisitor))
	//prods, scores, err := GetRecommendations(prefs, getIndVisitor(visitors, myVisitor), products)
	indexOfVisitor := getIndVisitor(visitors, myVisitor)
	if (indexOfVisitor == -1) {
		fmt.Println("Ërror: visitor doesn't found!")
		os.Exit(-1)
	}
	//t1_0 := time.Now()
	prods, scores, err := GetRecommendations(prefs, getIndVisitor(visitors, myVisitor), products)
	//t1_1 := time.Now()
	//real_scores := make ([]float64, 0)
	//real_scores = optimizeScores(scores, real_scores)
	if err != nil {
		fmt.Println("WHAT!?")
	}
	fmt.Print("Recommended Producs are: ")
	for i := 0; i < len(prods); i++  {
		if prods[i] != "" {
			fmt.Print(prods[i], " ")
		}
	}
//	fmt.Println(" with scores: ", real_scores)
	fmt.Print(" with scores: "/*, scores[i]*/)
	for i := 0; i < len(scores); i++ {
		if !math.IsNaN(scores[i]) {
			fmt.Print(scores[i], " ")
		}
	}
	var max_count float64
	for i := 0; i < len(arrayOfSales); i++  {
		if arrayOfSales[i] > max_count {
			max_count = arrayOfSales[i]
		}
	}
	fmt.Println(max_count)
	fmt.Println()
	//fmt.Println("Algorithm running time : ", t1_1.Sub(t1_0))
	//fmt.Println(len(prods))
	//fmt.Println(prods[5])
	//fmt.Printf("\nRecommended Products are: %v, with scores: %v", prods, scores)
	//fmt.Println()
}

