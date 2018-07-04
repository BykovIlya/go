package main

import (
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"os"
	"fmt"
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
	itemid_int int
}

type Visitor struct {
	visitorid_string string
	visitorid_int int
	items []*Items
}

func findInEvents(events []*Events, finder string) int {
	for i := 0; i < len(events); i++ {
		if events[i].visitorid == finder {
			return i;
		}
	}
	return -1;
}

/**
find element from array
 */
func find(buf []*Events, events []*Events, visitor []*Visitor) {
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
			Remove(buf, buf[i])
		}
	}
	Remove(buf,buf[0])
}

/**
delete elem from array
 */
func Remove(list []*Events, item *Events) []*Events {
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
	fmt.Println(len(events))
	//var MatrixOfSales [][] int64;

//	for i := 0;i < len(events);i++  {
//		fmt.Println(/*events[i].timestamp, " ", */events[i].visitorid, " ", /*events[i].event_, " ",*/ events[i].itemid/*, " ", events[i].transactionid*/);
//	}
	//var buf []*Events
	buf:= events
	var visitors []*Visitor
	//var items []*Items
	find(buf,events,visitors)
	//eventsJson, _ := json.Marshal(events)
	//fmt.Println(string(eventsJson))
	//fmt.Println("hello\n")
}
