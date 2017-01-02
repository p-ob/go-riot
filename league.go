package lolgo

import (
	"context"
	"errors"
	"fmt"
)

// LeagueService is the endpoint to use to get league information
type LeagueService struct {
	client *Client
}

// LeagueDto is the container returned with league information
type LeagueDto struct {
	Entries       []LeagueEntryDto `json:"entries"`
	Name          string           `json:"name"`
	ParticipantID int64            `json:"participantId"`
	Queue         string           `json:"queue"`
	Tier          string           `json:"tier"`
}

// LeagueEntryDto is the container for a single entry in a league
type LeagueEntryDto struct {
	Division         string        `json:"division"`
	IsFreshBlood     bool          `json:"isFreshBlood"`
	IsHotStreak      bool          `json:"isHotStreak"`
	IsInactive       bool          `json:"isInactive"`
	IsVeteran        bool          `json:"isVeteran"`
	LeaguePoints     int           `json:"leaguePoints"`
	Losses           int           `json:"losses"`
	MiniSeries       MiniSeriesDto `json:"miniSeries"`
	PlayerOrTeamID   int64         `json:"playerOrTeamId"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	PlayStyle        string        `json:"playstyle"`
	Wins             int           `json:"wins"`
}

// MiniSeriesDto is the container for an entry's promotional series
type MiniSeriesDto struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

const leaguePathPart = "api/lol/%s/v2.5/league"

// GetBySummoner gets the league data for given summonerIDs (up to 10)
// see https://developer.riotgames.com/api/methods#!/1215/4701 for distinction from GetEntriesBySummoner
func (s *LeagueService) GetBySummoner(ctx context.Context, summonerIDs ...int64) (*map[int64]LeagueDto, error) {
	if len(summonerIDs) > 10 {
		return nil, errors.New("Cannot pass more than 10 summoners to retrieve")
	}
	leagues := new(map[int64]LeagueDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.region)+"/by-summoner",
		int64ArrayToCommaDelimitedList(summonerIDs),
		nil,
		leagues,
	)
	return leagues, err
}

// GetEntriesBySummoner gets the league data for given summonerIDs (up to 10)
// see https://developer.riotgames.com/api/methods#!/1215/4705 for distinction from GetBySummoner
func (s *LeagueService) GetEntriesBySummoner(ctx context.Context, summonerIDs ...int64) (*map[int64]LeagueDto, error) {
	if len(summonerIDs) > 10 {
		return nil, errors.New("Cannot pass more than 10 summoners to retrieve")
	}
	leagues := new(map[int64]LeagueDto)
	err := s.client.getResource(
		ctx,
		fmt.Sprintf(
			"%s/by-summoner/%s/entry",
			addRegionToString(leaguePathPart, s.client.region),
			int64ArrayToCommaDelimitedList(summonerIDs),
		),
		"",
		nil,
		leagues,
	)
	return leagues, err
}

// GetChallenger gets the challenger league information
func (s *LeagueService) GetChallenger(ctx context.Context) (*LeagueDto, error) {
	league := new(LeagueDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.region)+"/challenger",
		"",
		nil,
		league,
	)
	return league, err
}

// GetMaster gets the master league information
func (s *LeagueService) GetMaster(ctx context.Context) (*LeagueDto, error) {
	league := new(LeagueDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(leaguePathPart, s.client.region)+"/master",
		"",
		nil,
		league,
	)
	return league, err
}
