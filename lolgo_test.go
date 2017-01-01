package lolgo

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestGetSummoner(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	summoner := generateSummoner()
	summonerId := summoner.Id
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

func generateSummoner() SummonerDto {
	return SummonerDto{
		Id:            rand.Int63(),
		Name:          randStringBytesMaskImprSrc(10),
		ProfileIconId: rand.Int(),
		RevisionDate:  rand.Int63(),
		SummonerLevel: rand.Int63n(30),
	}
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randStringBytesMaskImprSrc(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
