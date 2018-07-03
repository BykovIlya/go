package HelloWorld

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"encoding/json"
	"fmt"
)

type Events struct {
	timestamp string `json:"timestamp"`
	visitorid string `json:"visitorid"`
	event_ string `json:"event_type"`
	transactionid string `json:"transactionid"`

}
func main() {
	csvFile, _ := os.Open("events.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []Events
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		events = append(events, Events{
			timestamp: line[0],
			visitorid:  line[1],
			event_:  line[2],
			transactionid: line[3],
		})
	}
	eventsJson, _ := json.Marshal(events)
	fmt.Println(string(eventsJson))

}
