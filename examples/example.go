package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/p-ob/lolgo"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const region = lolgo.Na

func main() {
	apiKey, err := getAPIKey()
	if err != nil {
		printError(err)
		return
	}
	client := lolgo.NewClient(apiKey, region, nil)
	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter summoner to query: ")
	summonerName, _ := reader.ReadString('\n')
	// purge spaces since the json won't return these
	summonerName = strings.TrimSpace(strings.ToLower(summonerName))
	s, err := client.Summoner.GetByName(ctx, summonerName)
	if err != nil {
		printError(err)
		return
	}
	thisSummoner := (*s)[strings.ToLower(summonerName)]
	fmt.Printf("Summoner:\n%+v\n", thisSummoner)

	m, err := client.MatchList.GetBySummoner(ctx, thisSummoner.ID)
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("MatchList: \n%+v\n", *m)

	g, err := client.Game.GetRecent(ctx, thisSummoner.ID)
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("Games: \n%+v\n", *g)

	championMastery, err := client.ChampionMastery.GetAll(ctx, thisSummoner.ID)
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("Champion mastery: \n%+v\n", *championMastery)
}

func getAPIKey() (string, error) {
	file, err := filepath.Abs("key.txt")
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func printError(err error) {
	fmt.Printf("Error: %v", err)
}
