package lolgo

import (
	"context"
	"strconv"
)

type ChampionService struct {
	client *Client
}

type ChampionListDto struct {
	Champions []ChampionDto `json:"champions"`
}

type ChampionDto struct {
	Active            bool  `json:"active"`
	BotEnabled        bool  `json:"botEnabled"`
	BotMmEnabled      bool  `json:"botMmEnabled"`
	FreeToPlay        bool  `json:"freeToPlay"`
	Id                int64 `json:"id"`
	RankedPlayEnabled bool  `json:"rankedPlayEnabled"`
}

type GetChampionsParams struct {
	FreeToPlay bool `url:"freeToPlay,omitempty"`
}

const championPathPart = "api/lol/%s/v1.2/champion"

func (s *ChampionService) Get(ctx context.Context, championId int64) (*MatchList, error) {
	matchList := new(MatchList)
	err := s.client.GetResource(
		ctx,
		addRegionToString(championPathPart, s.client.Region),
		strconv.FormatInt(championId, 10),
		nil,
		matchList)
	return matchList, err
}

func (s *ChampionService) GetAll(ctx context.Context, params GetChampionsParams) (*MatchList, error) {
	matchList := new(MatchList)
	err := s.client.GetResource(
		ctx,
		addRegionToString(championPathPart, s.client.Region),
		"",
		params,
		matchList)
	return matchList, err
}
