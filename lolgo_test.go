package lolgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetSummoner(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	summonerId := int64(1)
	summoner := SummonerDto{Id: summonerId, Name: "Test", ProfileIconId: 1, RevisionDate: 12345678910, SummonerLevel: 30}
	getSummonerPathPart := fmt.Sprintf("/%s/%v", addRegionToString(summonerPathPart, region), summonerId)
	getSummonerResponse := make(map[int64]SummonerDto)
	getSummonerResponse[summonerId] = summoner

	summonerJsonByteArray, _ := json.Marshal(getSummonerResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getSummonerPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(summonerJsonByteArray)
	})

	ctx := context.Background()
	retrievedSummonerMap, err := client.Summoner.Get(ctx, summonerId)
	retrievedSummoner := (*retrievedSummonerMap)[summonerId]
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if retrievedSummoner.Id != summonerId {
		t.Errorf("expected %v, got %v", summonerId, retrievedSummoner.Id)
	}
}

func mockClient(region Region) (*http.Client, *http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	httpClient := &http.Client{Transport: transport}

	return httpClient, mux, server, NewClient("", region, httpClient)
}
