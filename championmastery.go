package lolgo

import (
	"context"
	"fmt"
	"strconv"
)

// ChampionMasteryService is the endpoint to use to get champion mastery information
type ChampionMasteryService struct {
	client *Client
}

// ChampionMasteryDto is the container returned per champion mastery
type ChampionMasteryDto struct {
	ChampionPoints               int   `json:"championPoints"`
	PlayerID                     int64 `json:"playerId"`
	ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionLevel                int   `json:"championLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionID                   int64 `json:"championId"`
	ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

// GetTopChampionsParams are optional query params
type GetTopChampionsParams struct {
	Count int `url:"count,omitempty"`
}

const championMasteryPathPart = "championmastery/location/%s/player/%v"

// Get gets the champion mastery for a single championID
func (s *ChampionMasteryService) Get(ctx context.Context, summonerID int64, championID int64) (*ChampionMasteryDto, error) {
	championMastery := new(ChampionMasteryDto)
	err := s.client.getResource(
		ctx,
		constructChampionMasteryPathPart(s.client.region, summonerID)+"/champion",
		strconv.FormatInt(championID, 10),
		nil,
		championMastery,
	)
	return championMastery, err
}

// GetAll gets all champion masteries for a summonerID
func (s *ChampionMasteryService) GetAll(ctx context.Context, summonerID int64) (*[]ChampionMasteryDto, error) {
	championMasteries := new([]ChampionMasteryDto)
	err := s.client.getResource(
		ctx,
		constructChampionMasteryPathPart(s.client.region, summonerID)+"/champions",
		"",
		nil,
		championMasteries,
	)
	return championMasteries, err
}

// GetTopChampions gets the top params.Count champion masteries by score (or 3 by default)
func (s *ChampionMasteryService) GetTopChampions(ctx context.Context, summonerID int64, params *GetTopChampionsParams) (*[]ChampionMasteryDto, error) {
	championMasteries := new([]ChampionMasteryDto)
	err := s.client.getResource(
		ctx,
		constructChampionMasteryPathPart(s.client.region, summonerID)+"/topchampions",
		"",
		params,
		championMasteries,
	)
	return championMasteries, err
}

// GetScore gets the total score for a summonerID
func (s *ChampionMasteryService) GetScore(ctx context.Context, summonerID int64) (*int, error) {
	score := new(int)
	err := s.client.getResource(
		ctx,
		constructChampionMasteryPathPart(s.client.region, summonerID)+"/score",
		"",
		nil,
		score,
	)
	return score, err
}

func constructChampionMasteryPathPart(region Region, summonerID int64) string {
	return fmt.Sprintf(championMasteryPathPart, mapRegionToLocationString(region), summonerID)
}
