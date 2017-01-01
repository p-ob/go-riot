package main

import (
	"github.com/p-ob/lolgo"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"bufio"
	"os"
	"strings"
)

const Region = lolgo.Na

func main() {
	apiKey, err := getApiKey()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	client := lolgo.NewClient(apiKey, Region, nil)
	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter summoner to query: ")
	summonerName, _ := reader.ReadString('\n')
	// purge spaces since the json won't return these
	summonerName = strings.TrimSpace(strings.ToLower(summonerName))
	s, err := client.Summoner.GetByName(ctx, summonerName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	thisSummoner := (*s)[strings.ToLower(summonerName)]
	fmt.Printf("Summoner:\n%+v\n", thisSummoner)

	m, _ := client.MatchList.GetBySummoner(ctx, thisSummoner.Id)
	fmt.Printf("MatchList: \n%+v\n", *m)

	g, _ := client.Game.GetRecent(ctx, thisSummoner.Id)
	fmt.Printf("Games: \n%+v\n", *g)

	championMastery, _ := client.ChampionMastery.GetAll(ctx, thisSummoner.Id)
	fmt.Printf("Champion mastery: \n%+v\n", *championMastery)
}

func getApiKey() (string, error) {
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
