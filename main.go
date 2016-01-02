// lolgo project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	TEST_SUMMONERS             = false
	TEST_MATCHLIST             = false
	TEST_MATCH                 = false
	TEST_STATIC_DATA_CHAMPIONS = false
	TEST_ITEMS                 = false
	TEST_MASTERIES             = false
	TEST_RUNES                 = false
	TEST_SUMMONER_SPELLS       = false
	TEST_CHAMPIONS             = true
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

	if TEST_STATIC_DATA_CHAMPIONS {
		staticChampions := a.GetAllChampionsStaticData(false)
		fmt.Printf("All staticchampions: %+v\n", staticChampions)

		staticThresh := a.GetChampionStaticData(412, true)
		fmt.Printf("Static Thresh: %+v\n", staticThresh)
	}

	if TEST_ITEMS {
		items := a.GetAllItems(false)
		fmt.Printf("All items: %+v\n", items)

		zekes := a.GetItem(3050, true)
		fmt.Printf("Zeke's: %+v\n", zekes)
	}

	if TEST_MASTERIES {
		masteries := a.GetAllMasteries(false)
		fmt.Printf("All masteries: %+v\n", masteries)

		thunderlords := a.GetMastery(6362, true)
		fmt.Printf("Thunderlord's: %+v\n", thunderlords)
	}

	if TEST_RUNES {
		runes := a.GetAllRunes(false)
		fmt.Printf("All runes: %+v\n", runes)
	}

	if TEST_SUMMONER_SPELLS {
		summonerSpells := a.GetAllSummonerSpells(false)
		fmt.Printf("All summoner spells: %+v\n", summonerSpells)
	}

	if TEST_CHAMPIONS {
		champions := a.GetAllChampions(false)
		fmt.Printf("All champions: %+v\n", champions)

		thresh := a.GetChampion(412)
		fmt.Printf("Thresh: %+v\n", thresh)

	}
}
