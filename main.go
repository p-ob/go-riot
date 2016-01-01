// lolgo project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileData, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := string(fileData)
	a := ApiInfo{apiKey, "na", 10}
	data := a.GetSummoners("drunk7irishman", "rastarockit", "ohsnap62")
	fmt.Printf("Summoners by name: %+v\n", data)
}
