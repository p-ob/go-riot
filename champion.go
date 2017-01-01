package lolgo

import (
	"context"
	"strconv"
)

// ChampionService is the endpoint to use to get quick information about champions
type ChampionService struct {
	client *Client
}

// ChampionListDto is the container returned by GetAll
type ChampionListDto struct {
	Champions []ChampionDto `json:"champions"`
}

// ChampionDto is the container returned per champion
type ChampionDto struct {
	Active            bool  `json:"active"`
	BotEnabled        bool  `json:"botEnabled"`
	BotMmEnabled      bool  `json:"botMmEnabled"`
	FreeToPlay        bool  `json:"freeToPlay"`
	ID                int64 `json:"id"`
	RankedPlayEnabled bool  `json:"rankedPlayEnabled"`
}

// GetChampionsParams are optional query params
type GetChampionsParams struct {
	FreeToPlay bool `url:"freeToPlay,omitempty"`
}

const championPathPart = "api/lol/%s/v1.2/champion"

// Get gets a single champion, by championId
func (s *ChampionService) Get(ctx context.Context, championID int64) (*ChampionDto, error) {
	champion := new(ChampionDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(championPathPart, s.client.region),
		strconv.FormatInt(championID, 10),
		nil,
		champion,
	)
	return champion, err
}

// GetAll gets all champions, or pass in params to filter down
func (s *ChampionService) GetAll(ctx context.Context, params GetChampionsParams) (*ChampionListDto, error) {
	champions := new(ChampionListDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(championPathPart, s.client.region),
		"",
		params,
		champions,
	)
	return champions, err
}
