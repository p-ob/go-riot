package lol

import (
	"context"
)

// StatusService is the endpoint to use to get server related information
type StatusService struct {
	client *Client
}

// Shard is the container per server
type Shard struct {
	Hostname  string   `json:"hostname"`
	Locales   []string `json:"locales"`
	Name      string   `json:"name"`
	RegionTag string   `json:"region_tag"`
	Slug      string   `json:"slug"`
}

const statusPathPart = "shards"

// Get gets the shard for a given region
func (s *StatusService) Get(ctx context.Context, region Region) (*Shard, error) {
	shard := new(Shard)

	err := s.client.getResource(
		ctx,
		statusPathPart,
		mapRegionToString(region),
		nil,
		shard,
	)
	return shard, err
}

// GetAll gets all shards
func (s *StatusService) GetAll(ctx context.Context) (*[]Shard, error) {
	shards := new([]Shard)

	err := s.client.getResource(
		ctx,
		statusPathPart,
		"",
		nil,
		shards,
	)
	return shards, err
}
