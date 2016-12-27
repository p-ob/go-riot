package lolgo

import (
	"context"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// lolgo metadata
const Name = "lolgo"
const Version = 0.1

type BaseParams struct {
	ApiKey string `url:"api_key,omitempty"`
}

type RiotApiError struct {
	Status RiotApiErrorStatus `json:"status,omitempty"`
}

type RiotApiErrorStatus struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"messasge,omitempty"`
}

type Region int

const (
	Na   Region = 0
	Euw  Region = 1
	Eune Region = 2
	Kr   Region = 3
	Lan  Region = 4
	Las  Region = 5
	Jp   Region = 6
	Ru   Region = 7
	Tr   Region = 8
	Oce  Region = 9
	Br   Region = 10
	Pbe  Region = 11
)

// The base URL serving the API. Override this for testing.
var BaseURL = "https://%s.api.pvp.net/"

type Client struct {
	*sling.Sling

	ApiKey string
	Region Region

	// The API Client uses these resources
	Summoner        *SummonerService
	Match           *MatchService
	MatchList       *MatchListService
	Champion        *ChampionService
	ChampionMastery *ChampionMasteryService
	CurrentGame *CurrentGameService
	FeaturedGames *FeaturedGamesService
	Game *GameService
	Stats *StatsService
}

const defaultTimeout = 30*time.Second + 500*time.Millisecond

func NewClient(apiKey string, region Region, httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	slingClient := sling.New().Client(httpClient).Base(addRegionToString(BaseURL, region))
	userAgentString := "%s/%v sling/1.1.0 (https://github.com/dghubble/sling) %s (%s/%s)"
	slingClient.Set(
		"User-Agent",
		fmt.Sprintf(
			userAgentString,
			Name,
			Version,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH))

	c := &Client{Sling: slingClient, ApiKey: apiKey, Region: region}
	c.Summoner = &SummonerService{client: c}
	c.Match = &MatchService{client: c}
	c.MatchList = &MatchListService{client: c}
	c.Champion = &ChampionService{client: c}
	c.ChampionMastery = &ChampionMasteryService{client: c}
	c.CurrentGame = &CurrentGameService{client: c}
	c.FeaturedGames = &FeaturedGamesService{client: c}
	c.Game = &GameService{client: c}
	c.Stats = &StatsService{client: c}

	return c
}

func (c *Client) GetResource(ctx context.Context, pathPart string, sid string, params interface{}, v interface{}) error {
	sidPart := pathPart
	if sid != "" {
		sidPart = strings.Join([]string{pathPart, sid}, "/")
	}
	baseParams := BaseParams{}
	baseParams.ApiKey = c.ApiKey
	req, err := c.New().Get(sidPart).QueryStruct(params).QueryStruct(baseParams).Request()
	if err == nil {
		// TODO error interface (not map[string]interface{})
		req.WithContext(ctx)
		c.Do(req, v, new(map[string]interface{}))
	}
	return err
}

// private methods

func addRegionToString(str string, region Region) string {
	stringRegion := mapRegionToString(region)
	return fmt.Sprintf(str, stringRegion)
}

func mapRegionToString(region Region) string {
	switch region {
	case Na:
		return "na"
	case Eune:
		return "eune"
	case Euw:
		return "euw"
	case Kr:
		return "kr"
	case Lan:
		return "lan"
	case Las:
		return "las"
	case Jp:
		return "jp"
	case Ru:
		return "ru"
	case Tr:
		return "tr"
	case Oce:
		return "oce"
	case Br:
		return "br"
	case Pbe:
		return "pbe"
	default:
		return "na"
	}
}

func mapRegionToLocationString(region Region) string {
	switch region {
	case Na:
		return "NA1"
	case Eune:
		return "EUN1"
	case Euw:
		return "EUW1"
	case Kr:
		return "KR"
	case Lan:
		return "LA1"
	case Las:
		return "LA2"
	case Jp:
		return "JP1"
	case Ru:
		return "RU"
	case Tr:
		return "TR1"
	case Oce:
		return "OC1"
	case Br:
		return "BR1"
	case Pbe:
		return "PBE1"
	default:
		return "NA1"
	}
}

func int64ArrayToCommaDelimitedList(intArray []int64) string {
	var strArray []string
	for _, intVal := range intArray {
		strArray = append(strArray, strconv.FormatInt(intVal, 10))
	}

	return strings.Join(strArray, ",")
}
