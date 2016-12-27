package lolgo

import (
	"context"
	"fmt"
)

type LeagueService struct {
	client *Client
}

type LeagueDto struct {
	Entries []LeagueEntryDto
	Name string
	ParticipantId string
	Queue string
	Tier string
}

type LeagueEntryDto struct {
	Division string
	IsFreshBlood bool
	IsHotStreak bool
	IsInactive bool
	IsVeteran bool
	LeaguePoints int
	Losses int
	MiniSeries MiniSeriesDto
	PlayerOrTeamId string
	PlayerOrTeamName string
	Playstyle string
	Wins int
}

type MiniSeriesDto struct {
	Losses int
	Progress string
	Target int
	Wins int
}

const leaguePathPart = "api/lol/%s/v2.5/league"

func (s *LeagueService) GetBySummoner(ctx context.Context, summonerIds ...int64) (*map[string]LeagueDto, error) {
	leagues := new(map[string]LeagueDto)
	err := s.client.GetResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.Region) + "/by-summoner",
		int64ArrayToCommaDelimitedList(summonerIds),
		nil,
		leagues)
	return leagues, err
}

func (s *LeagueService) GetEntriesBySummoner(ctx context.Context, summonerIds ...int64) (*map[string]LeagueDto, error) {
	leagues := new(map[string]LeagueDto)
	err := s.client.GetResource(
		ctx,
		fmt.Sprintf(
			"%s/by-summoner/%s/entry",
			addRegionToString(leaguePathPart, s.client.Region),
			int64ArrayToCommaDelimitedList(summonerIds)),
		"",
		nil,
		leagues)
	return leagues, err
}

func (s *LeagueService) GetChallenger(ctx context.Context) (*LeagueDto, error) {
	league := new(LeagueDto)
	err := s.client.GetResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.Region) + "/challenger",
		"",
		nil,
		league)
	return league, err
}


func (s *LeagueService) GetMaster(ctx context.Context) (*LeagueDto, error) {
	league := new(LeagueDto)
	err := s.client.GetResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.Region) + "/master",
		"",
		nil,
		league)
	return league, err
}

