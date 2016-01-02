// lolgo project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	TEST_SUMMONERS = true
	TEST_MATCHLIST = true
	TEST_MATCH     = true
	TEST_CHAMPIONS = true
	TEST_ITEMS     = true
)

func main() {
	fileData, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := string(fileData)
	a := Api{apiKey, "na", 10}

	if TEST_SUMMONERS || TEST_MATCHLIST || TEST_MATCH {
		summoners := a.GetSummoners("drunk7irishman", "rastarockit", "ohsnap62")
		fmt.Printf("Summoners by name: %+v\n", summoners)

		if TEST_MATCHLIST || TEST_MATCH {
			sId := summoners[0].Id
			matchList := a.GetRankedMatchList(sId)
			fmt.Printf("Match list for %d: %+v\n", sId, matchList)

			if TEST_MATCH {
				mId := matchList.Matches[0].MatchId
				match := a.GetMatch(mId)
				fmt.Printf("Match details for match %d: %+v\n", mId, match)
			}
		}
	}

	if TEST_CHAMPIONS {
		champions := a.GetAllChampions(false)
		fmt.Printf("All champions: %+v\n", champions)

		thresh := a.GetChampion(412, true)
		fmt.Printf("Thresh: %+v\n", thresh)
	}

	if TEST_ITEMS {
		items := a.GetAllItems(false)
		fmt.Printf("All items: %+v\n", items)

		zekes := a.GetItem(3050, true)
		fmt.Printf("Zeke's: %+v\n", zekes)
	}
}
