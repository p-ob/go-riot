package lolgo

import (
	"context"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// lolgo metadata

// Name is the name of this package
const Name = "lolgo"

// Version is the version of this package
const Version = 0.1

// baseParams are required by (almost) all queries
type baseParams struct {
	APIKey string `url:"api_key,omitempty"`
}

// RiotAPIError is the error structure returned by Riot (along with the headers returned for a 429 response)
type RiotAPIError struct {
	Status          riotAPIErrorStatus `json:"status"`
	XRateLimitType  string             `json:"-"`
	RetryAfter      int                `json:"-"`
	XRateLimitCount int                `json:"-"`
}

// riotAPIErrorStatus is the container with status information
type riotAPIErrorStatus struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Region is an "enum" convenience structure
type Region int

const (
	// Na is North America
	Na Region = 0

	// Euw is Europe West
	Euw Region = 1

	// Eune is Europe Nordic East
	Eune Region = 2

	// Kr is Korea
	Kr Region = 3

	// Lan is Latin America North
	Lan Region = 4

	// Las is Latin America South
	Las Region = 5

	// Jp is Japan
	Jp Region = 6

	// Ru is Russia
	Ru Region = 7

	// Tr is Turkey
	Tr Region = 8

	// Oce is Oceania
	Oce Region = 9

	// Br is Brazil
	Br Region = 10

	// Pbe is Public Beta Environment
	Pbe Region = 11
)

// BaseURL is the base url serving the API. Override this for testing.
var BaseURL = "https://%s.api.pvp.net"

// Client is the the API client hook to query Riot's League of Legends API
type Client struct {
	sling  *sling.Sling
	APIKey string
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

// NewClient constructs a client to handle API calls to the Riot League of Legends public API
// apiKey: unique key given by registering with https://developer.riotgames.com/
// region: the region to make queries against (Na, Euw, etc.)
// httpClient: if desired, provide your own instance of an httpClient; pass nil otherwise
func NewClient(apiKey string, region Region, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	slingClient := sling.New().Client(httpClient).Base(addRegionToString(BaseURL, region))
	userAgentString := "%s/%v (https://github.com/p-ob/lolgo) " +
		"sling/1.1.0 (https://github.com/dghubble/sling) %s (%s/%s)"
	slingClient.Set(
		"User-Agent",
		fmt.Sprintf(
			userAgentString,
			Name,
			Version,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		),
	)

	c := &Client{sling: slingClient, APIKey: apiKey, region: region}
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

// Region returns the region set for this Client
func (c Client) Region() Region {
	return c.region
}

func (c *Client) getResource(ctx context.Context, pathPart string, sid string, params interface{}, v interface{}) error {
	sidPart := pathPart
	if sid != "" {
		sidPart = strings.Join([]string{pathPart, sid}, "/")
	}
	riotError := new(RiotAPIError)
	baseParams := new(baseParams)
	baseParams.APIKey = c.APIKey
	req, err := c.sling.New().Get(sidPart).QueryStruct(params).QueryStruct(baseParams).Request()
	if err != nil {
		v = nil
		return err
	}
	req.WithContext(ctx)
	resp, err := c.sling.Do(req, v, riotError)
	if err != nil {
		v = nil
		return err
	}

	// handle special headers supplied when Riot returns a 429 response
	if resp.StatusCode == http.StatusTooManyRequests {
		riotError.XRateLimitType = resp.Header.Get("X-Rate-Limit-Type")
		retryAfter, err := strconv.ParseInt(resp.Header.Get("Retry-After"), 10, 32)
		if err == nil {
			riotError.RetryAfter = int(retryAfter)
		}
		xRateLimitCount, err := strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Count"), 10, 32)
		if err == nil {
			riotError.XRateLimitCount = int(xRateLimitCount)
		}
	}

	// if the API returns an error status code (4xx or 5xx), return the error
	if riotError.Status.StatusCode >= 400 {
		v = nil
		return riotError
	}
	return nil
}

// Error formats a RiotAPIError
func (e *RiotAPIError) Error() string {
	return fmt.Sprintf("lolgo: %+v", *e)
}

// private helper methods

func addRegionToString(str string, region Region) string {
	r := "(%#?[a-zA-Z])"
	isMatch, err := regexp.MatchString(r, str)
	if !isMatch || err != nil {
		return str
	}
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
