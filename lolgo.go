package lolgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
	"github.com/dghubble/sling"
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
	Message    string `json:"message,omitempty"`
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
	sling  *sling.Sling
	ApiKey string
	region Region

	// The API Client uses these resources
	Summoner        *SummonerService
	Match           *MatchService
	MatchList       *MatchListService
	Champion        *ChampionService
	ChampionMastery *ChampionMasteryService
	CurrentGame     *CurrentGameService
	FeaturedGames   *FeaturedGamesService
	Game            *GameService
	Stats           *StatsService
	League          *LeagueService
	Status          *StatusService
	StaticData      *StaticDataService
}

const defaultTimeout = 30*time.Second + 500*time.Millisecond

// Constructs a client to handle API calls to the Riot League of Legends public API
// apiKey: unique key given by registering with https://developer.riotgames.com/
// region: the region to make queries against (Na, Euw, etc.)
// httpClient: if desired, provide your own instance of an httpClient; pass nil otherwise
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

	c := &Client{sling: slingClient, ApiKey: apiKey, region: region}
	c.Summoner = &SummonerService{client: c}
	c.Match = &MatchService{client: c}
	c.MatchList = &MatchListService{client: c}
	c.Champion = &ChampionService{client: c}
	c.ChampionMastery = &ChampionMasteryService{client: c}
	c.CurrentGame = &CurrentGameService{client: c}
	c.FeaturedGames = &FeaturedGamesService{client: c}
	c.Game = &GameService{client: c}
	c.Stats = &StatsService{client: c}
	c.League = &LeagueService{client: c}
	c.Status = &StatusService{client: c}
	c.StaticData = &StaticDataService{client: c}

	return c
}

// returns the region set for this Client
func (c Client) Region() Region {
	return c.region
}

func (c *Client) getResource(ctx context.Context, pathPart string, sid string, params interface{}, v interface{}) error {
	sidPart := pathPart
	if sid != "" {
		sidPart = strings.Join([]string{pathPart, sid}, "/")
	}
	riotError := new(RiotApiError)
	baseParams := new(BaseParams)
	baseParams.ApiKey = c.ApiKey
	req, err := c.sling.New().Get(sidPart).QueryStruct(params).QueryStruct(baseParams).Request()
	if err == nil {
		req.WithContext(ctx)
		c.sling.Do(req, v, riotError)
	}
	if riotError.Status.StatusCode >= 400 {
		errorMsg := fmt.Sprintf(
			"Status: %v; Reason: %s",
			riotError.Status.StatusCode,
			riotError.Status.Message)
		err = errors.New(errorMsg)
	}
	return err
}

// private helper methods

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
