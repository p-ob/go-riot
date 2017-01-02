package lol

import (
	"context"
	"strconv"
)

// MatchListService is the endpoint to use to get matchlist information
type MatchListService struct {
	client *Client
}

// MatchList is the container for the matchlist endpoint
type MatchList struct {
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
	TotalGames int              `json:"totalGames"`
	Matches    []MatchReference `json:"matches"`
}

// MatchReference is the container for basic match information
type MatchReference struct {
	Champion   int64  `json:"champion"`
	Lane       string `json:"lane"`
	MatchID    int64  `json:"matchId"`
	PlatformID string `json:"platformId"`
	Queue      string `json:"queue"`
	Region     string `json:"region"`
	Role       string `json:"role"`
	Season     string `json:"season"`
	Timestamp  int64  `json:"timestamp"`
}

// GetMatchListParams are the optional query params
type GetMatchListParams struct {
	ChampionIDs  string `url:"championIds,omitempty"`
	RankedQueues string `url:"rankedQueues,omitempty"`
	Seasons      string `url:"seasons,omitempty"`
	BeginTime    int64  `url:"beginTime,omitempty"`
	EndTime      int64  `url:"endTime,omitempty"`
	BeginIndex   int    `url:"beginIndex,omitempty"`
	EndIndex     int    `url:"endIndex,omitempty"`
}

const matchListPathPath = "api/lol/%s/v2.2/matchlist/by-summoner"

// GetBySummoner gets the matchlist for a given summonerID
func (s *MatchListService) GetBySummoner(ctx context.Context, summonerID int64) (*MatchList, error) {
	matchList := new(MatchList)
	err := s.client.getResource(
		ctx,
		addRegionToString(matchListPathPath, s.client.region),
		strconv.FormatInt(summonerID, 10),
		nil,
		matchList,
	)
	return matchList, err
}
