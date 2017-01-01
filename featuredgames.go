package lolgo

import (
	"context"
)

// FeaturedGamesService maps the resource to use to get the featured games shown by the League of Legends client
type FeaturedGamesService struct {
	client *Client
}

// FeaturedGames is the container returned to list the featured games
type FeaturedGames struct {
	ClientRefreshInterval int64             `json:"clientRefreshInterval"`
	GameList              []CurrentGameInfo `json:"gameList"`
}

const featuredGamePathPart = "observer-mode/rest/featured"

// Get gets the featured games
func (s *FeaturedGamesService) Get(ctx context.Context) (*FeaturedGames, error) {
	featuredGames := new(FeaturedGames)

	err := s.client.getResource(
		ctx,
		featuredGamePathPart,
		"",
		nil,
		featuredGames,
	)
	return featuredGames, err
}
