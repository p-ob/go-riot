package lolgo

import (
	"context"
	"strconv"
)

type MatchListService struct {
	client *Client
}

type MatchList struct {
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
	TotalGames int              `json:"totalGames"`
	Matches    []MatchReference `json:"matches"`
}

type MatchReference struct {
	Champion   int64  `json:"champion"`
	Lane       string `json:"lane"`
	MatchId    int64  `json:"matchId"`
	PlatformId string `json:"platformId"`
	Queue      string `json:"queue"`
	Region     string `json:"region"`
	Role       string `json:"role"`
	Season     string `json:"season"`
	Timestamp  int64  `json:"timestamp"`
}

type GetMatchListParams struct {
	ChampionIds  string `url:"championIds,omitempty"`
	RankedQueues string `url:"rankedQueues,omitempty"`
	Seasons      string `url:"seasons,omitempty"`
	BeginTime    int64  `url:"beginTime,omitempty"`
	EndTime      int64  `url:"endTime,omitempty"`
	BeginIndex   int    `url:"beginIndex,omitempty"`
	EndIndex     int    `url:"endIndex,omitempty"`
}

const matchListPathPath = "api/lol/%s/v2.2/matchlist/by-summoner"

func (s *MatchListService) GetBySummoner(ctx context.Context, summonerId int64) (*MatchList, error) {
	matchList := new(MatchList)
	err := s.client.getResource(
		ctx,
		addRegionToString(matchListPathPath, s.client.region),
		strconv.FormatInt(summonerId, 10),
		nil,
		matchList)
	return matchList, err
}
