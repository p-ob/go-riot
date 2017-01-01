package lolgo

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
)

func TestGetSummoner(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	summoner := generateSummoner()
	getSummonerPathPart := fmt.Sprintf("/%s/%v", addRegionToString(summonerPathPart, region), summoner.ID)
	getSummonerResponse := make(map[int64]SummonerDto)
	getSummonerResponse[summoner.ID] = summoner

	summonerJSONByteArray, _ := json.Marshal(getSummonerResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getSummonerPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(summonerJSONByteArray)
	})

	ctx := context.Background()
	retrievedSummonerMap, err := client.Summoner.Get(ctx, summoner.ID)
	retrievedSummoner := (*retrievedSummonerMap)[summoner.ID]
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if !reflect.DeepEqual(summoner, retrievedSummoner) {
		t.Errorf("expected %v, got %v", summoner, retrievedSummoner)
	}
}

func TestGetSummonerByName(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	summoner := generateSummoner()
	getSummonerPathPart := fmt.Sprintf(
		"/%s/by-name/%s",
		addRegionToString(summonerPathPart, region),
		summoner.Name,
	)
	getSummonerResponse := make(map[string]SummonerDto)
	getSummonerResponse[summoner.Name] = summoner

	summonerJSONByteArray, _ := json.Marshal(getSummonerResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getSummonerPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(summonerJSONByteArray)
	})

	ctx := context.Background()
	retrievedSummonerMap, err := client.Summoner.GetByName(ctx, summoner.Name)
	retrievedSummoner := (*retrievedSummonerMap)[summoner.Name]
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if !reflect.DeepEqual(summoner, retrievedSummoner) {
		t.Errorf("expected %v, got %v", summoner, retrievedSummoner)
	}
}

func TestGetSummoner_MoreThan40(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	const nSummoners = 41
	getSummonerResponse := make(map[int64]SummonerDto, nSummoners)
	summoners := make([]SummonerDto, nSummoners)
	for i := 0; i < nSummoners; i++ {
		s := generateSummoner()
		summoners[i] = s
		getSummonerResponse[s.ID] = s
	}

	summonerIDs := make([]int64, nSummoners)
	for i, s := range summoners {
		summonerIDs[i] = s.ID
	}
	getSummonerPathPart := fmt.Sprintf(
		"/%s/%v",
		addRegionToString(summonerPathPart, region),
		int64ArrayToCommaDelimitedList(summonerIDs),
	)

	summonerJSONByteArray, _ := json.Marshal(getSummonerResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getSummonerPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(summonerJSONByteArray)
	})

	ctx := context.Background()
	retrievedSummonerMap, err := client.Summoner.Get(ctx, summonerIDs...)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
	if retrievedSummonerMap != nil {
		t.Errorf("expected nil, got %v", retrievedSummonerMap)
	}
}

func generateSummoner() SummonerDto {
	return SummonerDto{
		ID:            rand.Int63(),
		Name:          randString(10),
		ProfileIconID: rand.Int(),
		RevisionDate:  rand.Int63(),
		SummonerLevel: rand.Int63n(30),
	}
}
