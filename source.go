// used https://medium.com/@pradityadhitama/simple-logger-in-golang-f72dadf2c8c6
//this is the file writing log entries in logfile.log

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)
//define the log entry to have two fields (which are both string formats): Time and Text
type log_entry struct {
	Time string `json:"Time"`
	Text string `json:"Text"`
}

func main() {
	//open file
	file, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(file, "", 0)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	for {
		//create new entry as a combination of the current time and the text: "hello world!"
		curr := log_entry{
			Time: time.Now().Format(time.RFC3339),
			Text: "Hello World!",
		}
		//format the data into JSON 
		entry, err := json.Marshal(curr)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}
		//write the current entry into the logfile
		logger.Println(string(entry))
		time.Sleep(time.Second)
	}
}
