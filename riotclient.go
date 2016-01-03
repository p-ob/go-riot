package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	BaseUrl                   = "https://%s.api.pvp.net/api/lol/%s/%s?"
	SummonersByNameEndpoint   = "v1.4/summoner/by-name/%s"
	SummonersByIdEndpoint     = "v1.4/summoner/%s"
	MatchListEndpoint         = "v2.2/matchlist/by-summoner/%s"
	MatchEndpoint             = "v2.2/match/%s"
	ChampionsEndpoint         = "v1.2/champion"
	ChampionByIdEndpoint      = "v1.2/champion/%s"
	ItemsEndpoint             = "v1.2/item"
	ItemByIdEndpoint          = "v1.2/item/%s"
	MasteriesEndpoint         = "v1.2/mastery"
	MasteryByIdEndpoint       = "v1.2/mastery/%s"
	RunesEndpoint             = "v1.2/rune/"
	RuneByIdEndpoint          = "v1.2/rune/%s"
	SummonerSpellsEndpoint    = "v1.2/summoner-spell"
	SummonerSpellByIdEndpoint = "v1.2/summoner-spell/%s"
	RecentGamesEndpoint       = "v1.3/game/by-summoner/%s/recent"
	RankedStatsEndpoint       = "v1.3/stats/by-summoner/%s/ranked"
	SummaryStatsEndpoint      = "v1.3/stats/by-summoner/%s/summary"
	//
	StaticData = "static-data"
)

type RiotError struct {
	StatusCode int
	Reason     string
	Header     http.Header
}

func (err RiotError) Error() string {
	return fmt.Sprintf("Http Status: %d; %s", err.StatusCode, err.Reason)
}

type RiotClient struct {
	Key       string `api key given by https://developer.riotgames.com`
	Region    string `region to be queried against`
	RateLimit int    `rate limit given by https://developer.riotgames.com`
}

// public methods
func (api *RiotClient) GetSummoners(summonerNames ...string) []Summoner {
	url := api.constructUrl(SummonersByNameEndpoint, summonerNames...)

	summonersMap := make(map[string]Summoner)
	err := makeRequest(url, &summonersMap)

	if err != nil {
		handleError(err)
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

func (api *RiotClient) GetSummonersById(summonerIds ...int64) []Summoner {
	var summonerIdsStrings []string
	for _, summonerId := range summonerIds {
		summonerIdsStrings = append(summonerIdsStrings, strconv.FormatInt(summonerId, 10))
	}

	url := api.constructUrl(SummonersByIdEndpoint, summonerIdsStrings...)
	summonersMap := make(map[string]Summoner)
	err := makeRequest(url, &summonersMap)

	if err != nil {
		handleError(err)
	}

	var summonersArray []Summoner
	for _, summonerId := range summonerIdsStrings {
		if val, ok := summonersMap[summonerId]; ok {
			summonersArray = append(summonersArray, val)
		}
	}

	return summonersArray
}

func (api *RiotClient) GetRankedMatchList(summonerId int64) MatchList {
	url := api.constructUrl(MatchListEndpoint, strconv.FormatInt(summonerId, 10))

	matchList := MatchList{}
	err := makeRequest(url, &matchList)

	if err != nil {
		handleError(err)
	}

	return matchList
}

func (api *RiotClient) GetMatch(matchId int64) MatchDetail {
	url := api.constructUrl(MatchEndpoint, strconv.FormatInt(matchId, 10))

	match := MatchDetail{}
	err := makeRequest(url, &match)

	if err != nil {
		handleError(err)
	}

	return match
}

func (api *RiotClient) GetRecentGames(summonerId int64) RecentGamesDto {
	url := api.constructUrl(RecentGamesEndpoint, strconv.FormatInt(summonerId, 10))

	recentGames := RecentGamesDto{}
	err := makeRequest(url, &recentGames)

	if err != nil {
		handleError(err)
	}

	return recentGames
}

func (api *RiotClient) GetRankedStats(summonerId int64) RankedStatsDto {
	url := api.constructUrl(RankedStatsEndpoint, strconv.FormatInt(summonerId, 10))

	rankedStats := RankedStatsDto{}
	err := makeRequest(url, &rankedStats)

	if err != nil {
		handleError(err)
	}

	return rankedStats
}

func (api *RiotClient) GetSummaryStats(summonerId int64) PlayerStatsSummaryListDto {
	url := api.constructUrl(SummaryStatsEndpoint, strconv.FormatInt(summonerId, 10))

	summaryStats := PlayerStatsSummaryListDto{}
	err := makeRequest(url, &summaryStats)

	if err != nil {
		handleError(err)
	}

	return summaryStats
}

func (api *RiotClient) GetAllChampions(freeToPlay bool) ChampionListDto {
	url := api.constructUrl(ChampionsEndpoint)
	url += "&freeToPlay=" + strconv.FormatBool(freeToPlay)

	champions := ChampionListDto{}
	err := makeRequest(url, &champions)

	if err != nil {
		handleError(err)
	}

	return champions
}

func (api *RiotClient) GetChampion(id int) ChampionDto {
	url := api.constructUrl(ChampionByIdEndpoint, strconv.Itoa(id))

	champion := ChampionDto{}
	err := makeRequest(url, &champion)

	if err != nil {
		handleError(err)
	}

	return champion
}

func (api *RiotClient) GetAllChampionsStaticData(allData bool) StaticChampionListDto {
	url := api.constructStaticDataUrl(ChampionsEndpoint)
	if allData {
		url += "&champData=all"
	}

	champions := StaticChampionListDto{}
	err := makeRequest(url, &champions)

	if err != nil {
		handleError(err)
	}

	return champions
}

func (api *RiotClient) GetChampionStaticData(id int, allData bool) StaticChampionDto {
	url := api.constructStaticDataUrl(ChampionByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&champData=all"
	}

	champion := StaticChampionDto{}
	err := makeRequest(url, &champion)

	if err != nil {
		handleError(err)
	}

	return champion
}

func (api *RiotClient) GetAllItems(allData bool) ItemListDto {
	url := api.constructStaticDataUrl(ItemsEndpoint)
	if allData {
		url += "&itemListData=all"
	}

	items := ItemListDto{}
	err := makeRequest(url, &items)

	if err != nil {
		handleError(err)
	}

	return items
}

func (api *RiotClient) GetItem(id int, allData bool) ItemDto {
	url := api.constructStaticDataUrl(ItemByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&itemData=all"
	}

	item := ItemDto{}
	err := makeRequest(url, &item)

	if err != nil {
		handleError(err)
	}

	return item
}

func (api *RiotClient) GetAllMasteries(allData bool) MasteryListDto {
	url := api.constructStaticDataUrl(MasteriesEndpoint)
	if allData {
		url += "&masteryListData=all"
	}

	masteries := MasteryListDto{}
	err := makeRequest(url, &masteries)

	if err != nil {
		handleError(err)
	}

	return masteries
}

func (api *RiotClient) GetMastery(id int, allData bool) MasteryDto {
	url := api.constructStaticDataUrl(MasteryByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&masteryData=all"
	}

	mastery := MasteryDto{}
	err := makeRequest(url, &mastery)

	if err != nil {
		handleError(err)
	}

	return mastery
}

func (api *RiotClient) GetAllRunes(allData bool) RuneListDto {
	url := api.constructStaticDataUrl(RunesEndpoint)
	if allData {
		url += "&runeListData=all"
	}

	runes := RuneListDto{}
	err := makeRequest(url, &runes)

	if err != nil {
		handleError(err)
	}

	return runes
}

func (api *RiotClient) GetRune(id int, allData bool) RuneDto {
	url := api.constructStaticDataUrl(RuneByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&runeData=all"
	}

	runeObj := RuneDto{}
	err := makeRequest(url, &runeObj)

	if err != nil {
		handleError(err)
	}

	return runeObj
}

func (api *RiotClient) GetAllSummonerSpells(allData bool) SummonerSpellListDto {
	url := api.constructStaticDataUrl(SummonerSpellsEndpoint)
	if allData {
		url += "&spellData=all"
	}

	summonerSpells := SummonerSpellListDto{}
	err := makeRequest(url, &summonerSpells)

	if err != nil {
		handleError(err)
	}

	return summonerSpells
}

func (api *RiotClient) GetSummonerSpell(id int, allData bool) SummonerSpellDto {
	url := api.constructStaticDataUrl(SummonerSpellByIdEndpoint, strconv.Itoa(id))
	if allData {
		url += "&spellData=all"
	}

	summonerSpell := SummonerSpellDto{}
	err := makeRequest(url, &summonerSpell)

	if err != nil {
		handleError(err)
	}

	return summonerSpell
}

// private methods
func makeRequest(url string, v interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return RiotError{resp.StatusCode, resp.Status, resp.Header}
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

func handleError(err error) error {
	if riotErr, ok := err.(RiotError); ok {
		// handle TooManyRequests error
		if riotErr.StatusCode == 429 {
			header := riotErr.Header["Retry-After"]
			if len(header) > 0 {
				timeToSleep, err := strconv.Atoi(header[0])
				if err != nil {
					time.Sleep(time.Duration(timeToSleep) * time.Second)
				}
			}
		}
	}
	log.Fatal(err)
	return err
}

func (api *RiotClient) constructUrl(endpoint string, args ...string) string {
	url := fmt.Sprintf(BaseUrl, api.Region, api.Region, endpoint)
	if len(args) > 0 {
		url = fmt.Sprintf(url, strings.Join(args, ","))
	}
	url += "api_key=" + api.Key

	return url
}

func (api *RiotClient) constructStaticDataUrl(endpoint string, args ...string) string {
	staticDataUrlPart := fmt.Sprintf("%s/%s", StaticData, api.Region)
	url := fmt.Sprintf(BaseUrl, api.Region, staticDataUrlPart, endpoint)
	if len(args) > 0 {
		url = fmt.Sprintf(url, strings.Join(args, ","))
	}
	url += "api_key=" + api.Key

	return url
}
