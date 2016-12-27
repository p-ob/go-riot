package lolgo

import (
	"context"
)

type StatusService struct {
	client *Client
}

type Shard struct {
	Hostname  string   `json:"hostname"`
	Locales   []string `json:"locales"`
	Name      string   `json:"name"`
	RegionTag string   `json:"region_tag"`
	Slug      string   `json:"slug"`
}

const statusPathPart = "shards"

func (s *StatusService) Get(ctx context.Context, region Region) (*Shard, error) {
	shard := new(Shard)

	err := s.client.getResource(
		ctx,
		statusPathPart,
		mapRegionToString(region),
		nil,
		shard)
	return shard, err
}

func (s *StatusService) GetAll(ctx context.Context) (*[]Shard, error) {
	shards := new([]Shard)

	err := s.client.getResource(
		ctx,
		statusPathPart,
		"",
		nil,
		shards)
	return shards, err
}
