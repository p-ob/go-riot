package lolgo

import (
	"context"
	"errors"
	"strings"
)

// SummonerService is the endpoint to use to get summoner information
type SummonerService struct {
	client *Client
}

// SummonerDto is the container for a single summoner
type SummonerDto struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
}

// MasteryPagesDto is the container for all mastery pages for a summoner
type MasteryPagesDto struct {
	SummonerID int64            `json:"summonerId"`
	Pages      []MasteryPageDto `json:"pages"`
}

// MasteryPageDto is the container for a single mastery page
type MasteryPageDto struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Current   bool         `json:"current"`
	Masteries []MasteryDto `json:"masteries"`
}

// MasteryDto is the container for a single mastery in a mastery page
type MasteryDto struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}

// RunePagesDto is the container for all rune pages for a summoner
type RunePagesDto struct {
	SummonerID int64         `json:"summonerId"`
	Pages      []RunePageDto `json:"pages"`
}

// RunePageDto is the container for a single rune page
type RunePageDto struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Current bool      `json:"current"`
	Slots   []RuneDto `json:"slots"`
}

// RuneDto is the container for a single rune in a rune page
type RuneDto struct {
	RuneID     int `json:"runeId"`
	RuneSlotID int `json:"runeSlotId"`
}

const summonerPathPart = "api/lol/%s/v1.4/summoner"

// Get gets the summoner information for summonerIDs (at most 40); the map keys are the summonerIDs
func (s *SummonerService) Get(ctx context.Context, summonerIDs ...int64) (*map[int64]SummonerDto, error) {
	if len(summonerIDs) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	summoners := new(map[int64]SummonerDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIDs),
		nil,
		summoners,
	)
	return summoners, err
}

// GetByName gets the summoner information for summonerNames (at most 40); the map keys are the summonerNames
func (s *SummonerService) GetByName(ctx context.Context, summonerNames ...string) (*map[string]SummonerDto, error) {
	if len(summonerNames) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	summoners := new(map[string]SummonerDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region)+"/by-name",
		strings.Join(summonerNames, ","),
		nil,
		summoners,
	)
	return summoners, err
}

// GetNames gets the summoner names for the summonerIDs (at most 40); the map keys are the summonerIDs
func (s *SummonerService) GetNames(ctx context.Context, summonerIDs ...int64) (*map[int64]string, error) {
	if len(summonerIDs) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	names := new(map[int64]string)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIDs)+"/name",
		nil,
		names,
	)
	return names, err
}

// GetMasteries gets the masteries for the summonerIDs (at most 40); the map keys are the summonerIDs
func (s *SummonerService) GetMasteries(ctx context.Context, summonerIDs ...int64) (*map[int64]MasteryPagesDto, error) {
	if len(summonerIDs) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	masteries := new(map[int64]MasteryPagesDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIDs)+"/masteries",
		nil,
		masteries)
	return masteries, err
}

// GetRunes gets the runes for the summonerIDs (at most 40); the map keys are the summonerIDs
func (s *SummonerService) GetRunes(ctx context.Context, summonerIDs ...int64) (*map[int64]RunePagesDto, error) {
	if len(summonerIDs) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	runes := new(map[int64]RunePagesDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIDs)+"/runes",
		nil,
		runes)
	return runes, err
}
