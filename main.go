package main

import (
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"os"
	"fmt"
	"sort"
)

type Events struct {
//	timestamp string /*int64*/
	visitorid string /*int64*/
//	event_ string /*object*/
	itemid string /*int64*/
//	transactionid string /*float64*/

}

type Items struct{
	itemid_string string
	itemid_count int
}

type Visitor struct {
	visitorid_string string
	items [] Items
}

func getIndVisitor (visitor [] Visitor, finder string) int {
	for i := 0; i < len(visitor); i++ {
		if visitor[i].visitorid_string == finder {
			return i
		}
	}
	return -1
}

func getIndItem (items [] string, finder string) int {
	for i := 0; i < len(items); i++ {
		if items[i] == finder {
			return i
		}
	}
	return -1
}

func initVisitors (visitor []Visitor, buffer [] string) {
	for i := 0; i < len(buffer); i++ {
		visitor[i].visitorid_string = buffer[i]
	}
}

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

func initCountToResult (item []Items) {
	for i := 0; i < len(item); i++ {
		item[i].itemid_count = 1
	}
}

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

func main() {
	csvFile, _ := os.Open("events_example.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []*Events
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//if line[2] == "transaction" {
				events = append(events, &Events{
			//	timestamp:     line[0],
				visitorid:     line[0],
			//	event_ :       line[2],
				itemid:        line[1],
			//	transactionid: line[4],
			})
		//}
	}
	// for visitors
	bufOfVisitors := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfVisitors[i] = events[i].visitorid
	}
	removeDublicatesOfVisitors := make ([] string, len(events))
	sort.Strings(bufOfVisitors)
	removeDublicatesOfVisitors = removeDuplicates(bufOfVisitors)
	bufOfItems := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfItems[i] = events[i].itemid
	}
	sort.Strings(bufOfItems)
	removeDublicatesOfItems := make ([] string, len(events))
	removeDublicatesOfItems = removeDuplicates(bufOfItems)
	visitors := make([] Visitor, len(removeDublicatesOfVisitors))
	initVisitors(visitors, removeDublicatesOfVisitors)
	addItemsToVisitor(visitors,events)
	for i := 0; i < len(visitors); i++  {
		sort.Slice(visitors[i].items, func(j, k int) bool { return visitors[i].items[j].itemid_string < visitors[i].items[k].itemid_string })
	}
	for i := 0; i < len(visitors); i++ {
		visitors[i].items = findCount(visitors[i].items)
	}
	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		fmt.Println(visitors[i])
	}
	matrixOfSales := make([][]int, len(removeDublicatesOfVisitors))
	for i := 0; i < len(removeDublicatesOfVisitors); i++  {
		matrixOfSales[i] = make([]int, len(removeDublicatesOfItems))
	}
	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		for j := 0; j < len(visitors[i].items); j++ {
			//if visitors[i].items[j].itemid_count > 0 {
				matrixOfSales[i][getIndItem(removeDublicatesOfItems,visitors[i].items[j].itemid_string)] = visitors[i].items[j].itemid_count;
			//}
		}
	}
	for i:=0; i < len(removeDublicatesOfVisitors); i++ {
		for j:=0; j < len(removeDublicatesOfItems); j++ {
			//matrixOfSales[i][j] = visitors[i].items[j].itemid_count;
			fmt.Print(matrixOfSales[i][j], " ")
		}
		fmt.Println()
	}
}

