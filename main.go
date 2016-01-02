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
	a := Api{apiKey, "na", 10}
	summoners := a.GetSummoners("drunk7irishman", "rastarockit", "ohsnap62")
	fmt.Printf("Summoners by name: %+v\n", summoners)

	sId := summoners[0].Id
	matchList := a.GetRankedMatchList(sId)
	fmt.Printf("Match list for %d: %+v\n", sId, matchList)

	mId := matchList.Matches[0].MatchId
	match := a.GetMatch(mId)
	fmt.Printf("Match details for match %d: %+v\n", mId, match)

	champions := a.GetAllChampions(false)
	fmt.Printf("All champions: %+v\n", champions)

	thresh := a.GetChampion(412, true)
	fmt.Printf("Thresh: %+v\n", thresh)
}
