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
	BaseUrl                 = "https://%s.api.pvp.net/api/lol/%s/%s?"
	SummonersByNameEndpoint = "v1.4/summoner/by-name/%s"
	SummonersByIdEndpoint   = "v1.4/summoner/%s"
	MatchListEndpoint       = "v2.2/matchlist/by-summoner/%s"
	MatchEndpoint           = "v2.2/match/%s"
	ChampionsEndpoint       = "v1.2/champion"
	ChampionByIdEndpoint    = "v1.2/champion/%s"
	//
	StaticData = "static-data"
)

type RiotError struct {
	StatusCode int
	Reason     string
}

func (err RiotError) Error() string {
	return fmt.Sprintf("Http Status: %d; %s", err.StatusCode, err.Reason)
}

type Api struct {
	Key       string `api key given by https://developer.riotgames.com`
	Region    string `region to be queried against`
	RateLimit int    `rate limit given by https://developer.riotgames.com`
}

// public methods
func (api *Api) GetSummoners(summonerNames ...string) []Summoner {
	url := api.constructUrl(SummonersByNameEndpoint, summonerNames...)

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

func (api *Api) GetSummonersById(summonerIds ...int64) []Summoner {
	var summonerIdsStrings []string
	for _, summonerId := range summonerIds {
		summonerIdsStrings = append(summonerIdsStrings, strconv.FormatInt(summonerId, 10))
	}

	url := api.constructUrl(SummonersByIdEndpoint, summonerIdsStrings...)
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

func (api *Api) GetRankedMatchList(summonerId int64) MatchList {
	url := api.constructUrl(MatchListEndpoint, strconv.FormatInt(summonerId, 10))

	matchList := MatchList{}
	err := makeRequest(url, &matchList)

	if err != nil {
		log.Fatal(err)
	}

	return matchList
}

func (api *Api) GetMatch(matchId int64) MatchDetail {
	url := api.constructUrl(MatchEndpoint, strconv.FormatInt(matchId, 10))

	match := MatchDetail{}
	err := makeRequest(url, &match)

	if err != nil {
		log.Fatal(err)
	}

	return match
}

func (api *Api) GetAllChampions(allData bool) ChampionListDto {
	url := api.constructStaticDataUrl(ChampionsEndpoint)
	if allData {
		url += "&champData=all"
	}

	champions := ChampionListDto{}
	err := makeRequest(url, &champions)

	if err != nil {
		log.Fatal(err)
	}

	return champions
}

func (api *Api) GetChampion(id int, allData bool) ChampionDto {
	url := api.constructStaticDataUrl(ChampionByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&champData=all"
	}

	champion := ChampionDto{}
	err := makeRequest(url, &champion)

	if err != nil {
		log.Fatal(err)
	}

	return champion
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

	//fmt.Println(string(body))
	resp.Body.Close()

	if err != nil {
		return err
	}

	json.Unmarshal(body, v)

	return nil
}

func (api *Api) constructUrl(endpoint string, args ...string) string {
	url := fmt.Sprintf(BaseUrl, api.Region, api.Region, endpoint)
	if len(args) > 0 {
		url = fmt.Sprintf(url, strings.Join(args, ","))
	}
	url += "api_key=" + api.Key

	return url
}

func (api *Api) constructStaticDataUrl(endpoint string, args ...string) string {
	staticDataUrlPart := fmt.Sprintf("%s/%s", StaticData, api.Region)
	url := fmt.Sprintf(BaseUrl, api.Region, staticDataUrlPart, endpoint)
	if len(args) > 0 {
		url = fmt.Sprintf(url, strings.Join(args, ","))
	}
	url += "api_key=" + api.Key

	return url
}
