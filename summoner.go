package lolgo

import (
	"context"
	"errors"
	"strings"
)

type SummonerService struct {
	client *Client
}

type SummonerDto struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
}

type MasteryPagesDto struct {
	SummonerId int              `json:"summonerId"`
	Pages      []MasteryPageDto `json:"pages"`
}

type MasteryPageDto struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Current   bool         `json:"current"`
	Masteries []MasteryDto `json:"masteries"`
}

type MasteryDto struct {
	Id   int `json:"id"`
	Rank int `json:"rank"`
}

type RunePagesDto struct {
	SummonerId int              `json:"summonerId"`
	Pages      []MasteryPageDto `json:"pages"`
}

type RunePageDto struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Current bool      `json:"current"`
	Slots   []RuneDto `json:"slots"`
}

type RuneDto struct {
	RuneId     int `json:"runeId"`
	RuneSlotId int `json:"runeSlotId"`
}

const summonerPathPart = "api/lol/%s/v1.4/summoner"

func (s *SummonerService) Get(ctx context.Context, summonerIds ...int64) (*map[int64]SummonerDto, error) {
	if len(summonerIds) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	summoners := new(map[int64]SummonerDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIds),
		nil,
		summoners)
	return summoners, err
}

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
		summoners)
	return summoners, err
}

func (s *SummonerService) GetNames(ctx context.Context, summonerIds ...int64) (*map[int64]string, error) {
	if len(summonerIds) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	names := new(map[int64]string)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIds)+"/name",
		nil,
		names)
	return names, err
}

func (s *SummonerService) GetMasteries(ctx context.Context, summonerIds ...int64) (*map[int64]MasteryPagesDto, error) {
	if len(summonerIds) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	masteries := new(map[int64]MasteryPagesDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIds)+"/masteries",
		nil,
		masteries)
	return masteries, err
}

func (s *SummonerService) GetRunes(ctx context.Context, summonerIds ...int64) (*map[int64]RunePagesDto, error) {
	if len(summonerIds) > 40 {
		return nil, errors.New("Cannot pass more than 40 summoners to retrieve")
	}
	runes := new(map[int64]RunePagesDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIds)+"/runes",
		nil,
		runes)
	return runes, err
}
