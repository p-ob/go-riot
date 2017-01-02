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

func TestChampionMasteryService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	championMastery := generateChampionMasteryDto()
	pathPart := fmt.Sprintf(
		"/%s/champion/%v",
		constructChampionMasteryPathPart(region, championMastery.PlayerID),
		championMastery.ChampionID,
	)

	jsonByteArray, _ := json.Marshal(championMastery)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedChampionMastery, err := client.ChampionMastery.Get(ctx, championMastery.PlayerID, championMastery.ChampionID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(championMastery, *retrievedChampionMastery) {
		t.Errorf("expected %+v, got %+v", championMastery, *retrievedChampionMastery)
	}
}

func TestChampionMasteryService_GetAll(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	championMasteries := []ChampionMasteryDto{generateChampionMasteryDto()}
	summonerID := championMasteries[0].PlayerID
	pathPart := fmt.Sprintf(
		"/%s/champions",
		constructChampionMasteryPathPart(region, summonerID),
	)

	jsonByteArray, _ := json.Marshal(championMasteries)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedChampionMasteries, err := client.ChampionMastery.GetAll(ctx, summonerID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(championMasteries, *retrievedChampionMasteries) {
		t.Errorf("expected %+v, got %+v", championMasteries, *retrievedChampionMasteries)
	}
}

func TestChampionMasteryService_GetScore(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	score := r1.Int()
	summonerID := r1.Int63()
	pathPart := fmt.Sprintf(
		"/%s/score",
		constructChampionMasteryPathPart(region, summonerID),
	)

	jsonByteArray, _ := json.Marshal(score)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedScore, err := client.ChampionMastery.GetScore(ctx, summonerID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(score, *retrievedScore) {
		t.Errorf("expected %+v, got %+v", score, *retrievedScore)
	}
}

func TestChampionMasteryService_GetTopChampions(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	championMasteries := []ChampionMasteryDto{generateChampionMasteryDto()}
	summonerID := championMasteries[0].PlayerID
	pathPart := fmt.Sprintf(
		"/%s/topchampions",
		constructChampionMasteryPathPart(region, summonerID),
	)

	jsonByteArray, _ := json.Marshal(championMasteries)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedChampionMasteries, err := client.ChampionMastery.GetTopChampions(ctx, summonerID, nil)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(championMasteries, *retrievedChampionMasteries) {
		t.Errorf("expected %+v, got %+v", championMasteries, *retrievedChampionMasteries)
	}
}

func generateChampionMasteryDto() ChampionMasteryDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return ChampionMasteryDto{
		ChampionPoints:               r1.Int(),
		PlayerID:                     r1.Int63(),
		ChampionPointsUntilNextLevel: r1.Int63(),
		ChestGranted:                 false,
		ChampionLevel:                r1.Int(),
		TokensEarned:                 r1.Int(),
		ChampionID:                   r1.Int63(),
		ChampionPointsSinceLastLevel: r1.Int63(),
		LastPlayTime:                 r1.Int63(),
	}
}
