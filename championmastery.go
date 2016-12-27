package lolgo

import (
	"context"
	"fmt"
	"strconv"
)

type ChampionMasteryService struct {
	client *Client
}

type ChampionMasteryDto struct {
	ChampionPoints               int   `json:"championPoints"`
	PlayerId                     int   `json:"playerId"`
	ChampionPointsUntilNextLevel int   `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionLevel                int   `json:"championLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionId                   int   `json:"championId"`
	ChampionPointsSinceLastLevel int   `json:"championPointsSinceLastLevel"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

type GetTopChampionsParams struct {
	Count int `url:"count,omitempty"`
}

const championMasteryPathPart = "championmastery/location/%s/player/%s"

func (s *ChampionMasteryService) Get(ctx context.Context, summonerId int64, championId int64) (*ChampionMasteryDto, error) {
	championMastery := new(ChampionMasteryDto)
	err := s.client.GetResource(
		ctx,
		constructChampionMasteryPathPart(s.client.Region, summonerId)+"/champion",
		strconv.FormatInt(championId, 10),
		nil,
		championMastery)
	return championMastery, err
}

func (s *ChampionMasteryService) GetAll(ctx context.Context, summonerId int64) (*[]ChampionMasteryDto, error) {
	championMasteries := new([]ChampionMasteryDto)
	err := s.client.GetResource(
		ctx,
		constructChampionMasteryPathPart(s.client.Region, summonerId)+"/champions",
		"",
		nil,
		championMasteries)
	return championMasteries, err
}

func (s *ChampionMasteryService) GetTopChampions(ctx context.Context, summonerId int64, params *GetTopChampionsParams) (*[]ChampionMasteryDto, error) {
	championMasteries := new([]ChampionMasteryDto)
	err := s.client.GetResource(
		ctx,
		constructChampionMasteryPathPart(s.client.Region, summonerId)+"/topchampions",
		"",
		params,
		championMasteries)
	return championMasteries, err
}

func (s *ChampionMasteryService) GetScore(ctx context.Context, summonerId int64) (*int, error) {
	score := new(int)
	err := s.client.GetResource(
		ctx,
		constructChampionMasteryPathPart(s.client.Region, summonerId)+"/score",
		"",
		nil,
		score)
	return score, err
}

func constructChampionMasteryPathPart(region Region, summonerId int64) string {
	return fmt.Sprintf(championMasteryPathPart, mapRegionToLocationString(region), summonerId)
}
