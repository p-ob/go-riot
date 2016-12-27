package lolgo

import (
	"context"
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
	Masteries []MasteryDto `json:"masteries,omitempty"`
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
	slots   []RuneDto `json:"slots,omitempty"`
}

type RuneDto struct {
	RuneId     int `json:"runeId"`
	RuneSlotId int `json:"runeSlotId"`
}

const summonerPathPart = "api/lol/%s/v1.4/summoner"

func (s *SummonerService) Get(ctx context.Context, summonerIds ...int64) (*map[int64]SummonerDto, error) {
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
	runes := new(map[int64]RunePagesDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(summonerPathPart, s.client.region),
		int64ArrayToCommaDelimitedList(summonerIds)+"/runes",
		nil,
		runes)
	return runes, err
}
