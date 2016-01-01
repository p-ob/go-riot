package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	BaseUrl                 = "https://%s.api.pvp.net/api/lol/%s/"
	SummonersByNameEndpoint = "v1.4/summoner/by-name/%s"
	SummonersByIdEndpoint   = "v1.4/summoner/%s"
)

type RiotError struct {
	StatusCode int
	Reason     string
}

func (err RiotError) Error() string {
	return fmt.Sprintf("Http Status: %d; %s", err.StatusCode, err.Reason)
}

type ApiInfo struct {
	Key       string `api key given by https://developer.riotgames.com`
	Region    string `region to be queried against`
	RateLimit int    `rate limit given by https://developer.riotgames.com`
}

// public methods
func (api *ApiInfo) GetSummoners(summonerNames ...string) []Summoner {
	url := BaseUrl + SummonersByNameEndpoint
	url = fmt.Sprintf(url, api.Region, api.Region, strings.Join(summonerNames, ","))
	url += "?api_key=" + api.Key

	summonersMap := make(map[string]Summoner)
	err := makeRequest(url, &summonersMap)

	if err != nil {
		log.Fatal(err)
	}

	var summonersArray []Summoner
	for _, summonerName := range summonerNames {
		normalizedSummonerName := strings.ToLower(summonerName)
		normalizedSummonerName = strings.Replace(normalizedSummonerName, " ", "", -1)
		if val, ok := summonersMap[normalizedSummonerName]; ok {
			summonersArray = append(summonersArray, val)
		}
	}

	return summonersArray
}

func (api *ApiInfo) GetSummonersById(summonerIds ...int) []Summoner {
	var summonerIdsStrings []string
	for _, summonerId := range summonerIds {
		summonerIdsStrings = append(summonerIdsStrings, strconv.Itoa(summonerId))
	}

	url := BaseUrl + SummonersByIdEndpoint
	url = fmt.Sprintf(url, api.Region, api.Region, strings.Join(summonerIdsStrings, ","))
	url += "?api_key=" + api.Key
	summonersMap := make(map[string]Summoner)
	err := makeRequest(url, &summonersMap)

	if err != nil {
		log.Fatal(err)
	}

	var summonersArray []Summoner
	for _, summonerId := range summonerIdsStrings {
		if val, ok := summonersMap[summonerId]; ok {
			summonersArray = append(summonersArray, val)
		}
	}

	return summonersArray
}

// private methods
func makeRequest(url string, v interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return RiotError{resp.StatusCode, resp.Status}
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	json.Unmarshal(body, v)

	return nil
}
