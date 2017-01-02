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

func TestLeagueService_GetBySummoner(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	league := generateLeagueDto()
	summonerID := r1.Int63()
	pathPart := fmt.Sprintf("/%s/by-summoner/%v", addRegionToString(leaguePathPart, region), summonerID)
	getResponse := make(map[int64]LeagueDto)
	getResponse[summonerID] = league

	jsonByteArray, _ := json.Marshal(getResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeagueMap, err := client.League.GetBySummoner(ctx, summonerID)
	retrievedLeague := (*retrievedLeagueMap)[summonerID]
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(league, retrievedLeague) {
		t.Errorf("expected %+v, got %+v", league, retrievedLeague)
	}
}

func TestLeagueService_GetEntriesBySummoner(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	league := generateLeagueDto()
	summonerID := r1.Int63()
	pathPart := fmt.Sprintf("/%s/by-summoner/%v/entry", addRegionToString(leaguePathPart, region), summonerID)
	getResponse := make(map[int64]LeagueDto)
	getResponse[summonerID] = league

	jsonByteArray, _ := json.Marshal(getResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeagueMap, err := client.League.GetEntriesBySummoner(ctx, summonerID)
	retrievedLeague := (*retrievedLeagueMap)[summonerID]
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(league, retrievedLeague) {
		t.Errorf("expected %+v, got %+v", league, retrievedLeague)
	}
}

func TestLeagueService_GetBySummoner_MoreThan10(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	const nSummoners = 11
	getResponse := make(map[int64]LeagueDto, nSummoners)
	summonerIDs := make([]int64, nSummoners)
	for i := 0; i < nSummoners; i++ {
		sID := r1.Int63()
		l := generateLeagueDto()
		summonerIDs[i] = sID
		getResponse[sID] = l
	}

	pathPart := fmt.Sprintf(
		"/%s/by-summoner/%v",
		addRegionToString(leaguePathPart, region),
		int64ArrayToCommaDelimitedList(summonerIDs),
	)

	jsonByteArray, _ := json.Marshal(getResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeagueMap, err := client.League.GetBySummoner(ctx, summonerIDs...)
	if err == nil {
		t.Errorf("expected error, got %+v", err)
	}
	if retrievedLeagueMap != nil {
		t.Errorf("expected nil, got %+v", retrievedLeagueMap)
	}
}

func TestLeagueService_GetEntriesBySummoner_MoreThan10(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	const nSummoners = 11
	getResponse := make(map[int64]LeagueDto, nSummoners)
	summonerIDs := make([]int64, nSummoners)
	for i := 0; i < nSummoners; i++ {
		sID := r1.Int63()
		l := generateLeagueDto()
		summonerIDs[i] = sID
		getResponse[sID] = l
	}

	pathPart := fmt.Sprintf(
		"/%s/by-summoner/%v/entry",
		addRegionToString(leaguePathPart, region),
		int64ArrayToCommaDelimitedList(summonerIDs),
	)

	jsonByteArray, _ := json.Marshal(getResponse)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeagueMap, err := client.League.GetEntriesBySummoner(ctx, summonerIDs...)
	if err == nil {
		t.Errorf("expected error, got %+v", err)
	}
	if retrievedLeagueMap != nil {
		t.Errorf("expected nil, got %+v", retrievedLeagueMap)
	}
}

func TestLeagueService_GetChallenger(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	league := generateLeagueDto()
	pathPart := fmt.Sprintf("/%s/challenger", addRegionToString(leaguePathPart, region))

	jsonByteArray, _ := json.Marshal(league)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeague, err := client.League.GetChallenger(ctx)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(league, *retrievedLeague) {
		t.Errorf("expected %+v, got %+v", league, *retrievedLeague)
	}
}

func TestLeagueService_GetMaster(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	league := generateLeagueDto()
	pathPart := fmt.Sprintf("/%s/master", addRegionToString(leaguePathPart, region))

	jsonByteArray, _ := json.Marshal(league)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedLeague, err := client.League.GetMaster(ctx)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(league, *retrievedLeague) {
		t.Errorf("expected %+v, got %+v", league, *retrievedLeague)
	}
}

func generateLeagueDto() LeagueDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return LeagueDto{
		Entries:       []LeagueEntryDto{generateLeagueEntryDto()},
		Name:          randString(10),
		ParticipantID: r1.Int63(),
		Queue:         randString(10),
		Tier:          randString(10),
	}
}

func generateLeagueEntryDto() LeagueEntryDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return LeagueEntryDto{
		Division:         randString(10),
		IsFreshBlood:     false,
		IsHotStreak:      false,
		IsInactive:       false,
		IsVeteran:        false,
		LeaguePoints:     r1.Int(),
		MiniSeries:       generateMiniSeriesDto(),
		PlayerOrTeamID:   r1.Int63(),
		PlayerOrTeamName: randString(10),
		PlayStyle:        randString(10),
		Wins:             r1.Int(),
	}
}

func generateMiniSeriesDto() MiniSeriesDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return MiniSeriesDto{
		Losses:   r1.Int(),
		Progress: randString(10),
		Target:   r1.Int(),
		Wins:     r1.Int(),
	}
}
