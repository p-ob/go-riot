package lolgo

import (
	"context"
	"fmt"
)

type LeagueService struct {
	client *Client
}

type LeagueDto struct {
	Entries       []LeagueEntryDto `json:"entries"`
	Name          string           `json:"name"`
	ParticipantId string           `json:"participantId"`
	Queue         string           `json:"queue"`
	Tier          string           `json:"tier"`
}

type LeagueEntryDto struct {
	Division         string        `json:"division"`
	IsFreshBlood     bool          `json:"isFreshBlood"`
	IsHotStreak      bool          `json:"isHotStreak"`
	IsInactive       bool          `json:"isInactive"`
	IsVeteran        bool          `json:"isVeteran"`
	LeaguePoints     int           `json:"leaguePoints"`
	Losses           int           `json:"losses"`
	MiniSeries       MiniSeriesDto `json:"miniSeries"`
	PlayerOrTeamId   string        `json:"playerOrTeamId"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Playstyle        string        `json:"playstyle"`
	Wins             int           `json:"wins"`
}

type MiniSeriesDto struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

const leaguePathPart = "api/lol/%s/v2.5/league"

func (s *LeagueService) GetBySummoner(ctx context.Context, summonerIds ...int64) (*map[string]LeagueDto, error) {
	leagues := new(map[string]LeagueDto)
	err := s.client.GetResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.Region)+"/by-summoner",
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
		addRegionToString(leaguePathPart, s.client.Region)+"/challenger",
		"",
		nil,
		league)
	return league, err
}

func (s *LeagueService) GetMaster(ctx context.Context) (*LeagueDto, error) {
	league := new(LeagueDto)
	err := s.client.GetResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.Region)+"/master",
		"",
		nil,
		league)
	return league, err
}
