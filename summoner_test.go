package lolgo

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

func TestSummonerService_Get(t *testing.T) {
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
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(summoner, retrievedSummoner) {
		t.Errorf("expected %+v, got %+v", summoner, retrievedSummoner)
	}
}

func TestSummonerService_GetByName(t *testing.T) {
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
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(summoner, retrievedSummoner) {
		t.Errorf("expected %+v, got %+v", summoner, retrievedSummoner)
	}
}

func TestSummonerService_GetName(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	summoner := generateSummoner()
	getNamePathPart := fmt.Sprintf(
		"/%s/%v/name",
		addRegionToString(summonerPathPart, region),
		summoner.ID,
	)
	getNameResponse := make(map[int64]string)
	getNameResponse[summoner.ID] = summoner.Name
	jsonByteArray, _ := json.Marshal(getNameResponse)
	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getNamePathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedNamesMap, err := client.Summoner.GetNames(ctx, summoner.ID)
	retrievedName := (*retrievedNamesMap)[summoner.ID]
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if retrievedName != summoner.Name {
		t.Errorf("expected %+v, got %+v", summoner.Name, retrievedName)
	}
}

func TestSummonerService_GetMasteries(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	masteryPages := generateMasteryPagesDto()
	getMasteriesPathPart := fmt.Sprintf(
		"/%s/%v/masteries",
		addRegionToString(summonerPathPart, region),
		masteryPages.SummonerID,
	)
	getMasteriesResponse := make(map[int64]MasteryPagesDto)
	getMasteriesResponse[masteryPages.SummonerID] = masteryPages
	jsonByteArray, _ := json.Marshal(getMasteriesResponse)
	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getMasteriesPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedMasteryPagesMap, err := client.Summoner.GetMasteries(ctx, masteryPages.SummonerID)
	retrievedMasteryPages := (*retrievedMasteryPagesMap)[masteryPages.SummonerID]
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(masteryPages, retrievedMasteryPages) {
		t.Errorf("expected %+v, got %+v", masteryPages, retrievedMasteryPages)
	}
}

func TestSummonerService_GetRunes(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	runePages := generateRunePagesDto()
	getRunesPathPart := fmt.Sprintf(
		"/%s/%v/runes",
		addRegionToString(summonerPathPart, region),
		runePages.SummonerID,
	)
	getRunesResponse := make(map[int64]RunePagesDto)
	getRunesResponse[runePages.SummonerID] = runePages
	jsonByteArray, _ := json.Marshal(getRunesResponse)
	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(getRunesPathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedRunePagesMap, err := client.Summoner.GetRunes(ctx, runePages.SummonerID)
	retrievedRunePages := (*retrievedRunePagesMap)[runePages.SummonerID]
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(runePages, retrievedRunePages) {
		t.Errorf("expected %+v, got %+v", runePages, retrievedRunePages)
	}
}

func TestSummonerService_Get_MoreThan40(t *testing.T) {
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
		t.Errorf("expected error, got %+v", err)
	}
	if retrievedSummonerMap != nil {
		t.Errorf("expected nil, got %+v", retrievedSummonerMap)
	}
}

func generateSummoner() SummonerDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return SummonerDto{
		ID:            r1.Int63(),
		Name:          randString(10),
		ProfileIconID: r1.Int(),
		RevisionDate:  r1.Int63(),
		SummonerLevel: r1.Int63n(30),
	}
}

func generateMasteryPagesDto() MasteryPagesDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pages := make([]MasteryPageDto, r1.Intn(3))
	for i := 0; i < len(pages); i++ {
		pages[i] = generateMasteryPageDto()
	}
	return MasteryPagesDto{
		SummonerID: r1.Int63(),
		Pages:      pages,
	}
}

func generateMasteryPageDto() MasteryPageDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	masteries := make([]MasteryDto, r1.Intn(3))
	for i := 0; i < len(masteries); i++ {
		masteries[i] = generateMasteryDto()
	}
	return MasteryPageDto{
		ID:        r1.Int(),
		Name:      randString(10),
		Current:   false,
		Masteries: masteries,
	}
}

func generateMasteryDto() MasteryDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return MasteryDto{
		ID:   r1.Int(),
		Rank: r1.Intn(3),
	}
}

func generateRunePagesDto() RunePagesDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pages := make([]RunePageDto, r1.Intn(3))
	for i := 0; i < len(pages); i++ {
		pages[i] = generateRunePageDto()
	}
	return RunePagesDto{
		SummonerID: r1.Int63(),
		Pages:      pages,
	}
}

func generateRunePageDto() RunePageDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	runes := make([]RuneDto, r1.Intn(3))
	for i := 0; i < len(runes); i++ {
		runes[i] = generateRuneDto()
	}
	return RunePageDto{
		ID:      r1.Int(),
		Name:    randString(10),
		Current: false,
		Slots:   runes,
	}
}

func generateRuneDto() RuneDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return RuneDto{
		RuneID:     r1.Int(),
		RuneSlotID: r1.Intn(30),
	}
}
