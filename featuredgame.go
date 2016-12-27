package lolgo

import (
	"context"
)

type FeaturedGamesService struct {
	client *Client
}

type FeaturedGames struct {
	ClientRefreshInterval int64
	GameList []CurrentGameInfo
}

const featuredGamePathPart = "observer-mode/rest/featured"

func (s *FeaturedGamesService) Get(ctx context.Context) (*FeaturedGames, error) {
	featuredGames := new(FeaturedGames)

	err := s.client.GetResource(
		ctx,
		featuredGamePathPart,
		"",
		nil,
		featuredGames)
	return featuredGames, err
}


