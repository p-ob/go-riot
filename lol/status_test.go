package lol

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestStatusService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	shard := generateShard()
	shard.RegionTag = mapRegionToString(region)
	getShardsPathPart := fmt.Sprintf("/%s/%v", addRegionToString(statusPathPart, region), shard.RegionTag)
	jsonByteArray, _ := json.Marshal(shard)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getShardsPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedShard, err := client.Status.Get(ctx, region)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(shard, *retrievedShard) {
		t.Errorf("expected %+v, got %+v", shard, *retrievedShard)
	}
}

func TestStatusService_GetAll(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	const shardsCount = 7
	shards := make([]Shard, shardsCount)
	for i := 0; i < shardsCount; i++ {
		shards[i] = generateShard()
	}
	getShardsPathPart := fmt.Sprintf("/%s", addRegionToString(statusPathPart, region))
	jsonByteArray, _ := json.Marshal(shards)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getShardsPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedShards, err := client.Status.GetAll(ctx)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(shards, *retrievedShards) {
		t.Errorf("expected %+v, got %+v", shards, *retrievedShards)
	}
}

func generateShard() Shard {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	locales := make([]string, r1.Intn(3))
	for i := 0; i < len(locales); i++ {
		locales[i] = randString(10)
	}
	return Shard{
		Hostname:  randString(10),
		Locales:   locales,
		Name:      randString(10),
		RegionTag: randString(3),
		Slug:      randString(10),
	}
}
