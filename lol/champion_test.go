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

func TestChampionService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	champion := generateChampionDto()
	pathPart := fmt.Sprintf("/%s/%v", addRegionToString(championPathPart, region), champion.ID)

	jsonByteArray, _ := json.Marshal(champion)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedChampion, err := client.Champion.Get(ctx, champion.ID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(champion, *retrievedChampion) {
		t.Errorf("expected %+v, got %+v", champion, *retrievedChampion)
	}
}

func TestChampionService_GetAll(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	champions := generateChampionListDto()
	pathPart := fmt.Sprintf("/%s", addRegionToString(championPathPart, region))

	jsonByteArray, _ := json.Marshal(champions)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedChampions, err := client.Champion.GetAll(ctx, nil)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(champions, *retrievedChampions) {
		t.Errorf("expected %+v, got %+v", champions, *retrievedChampions)
	}
}

func generateChampionListDto() ChampionListDto {
	return ChampionListDto{
		Champions: []ChampionDto{generateChampionDto()},
	}
}

func generateChampionDto() ChampionDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return ChampionDto{
		Active:            false,
		BotEnabled:        false,
		BotMmEnabled:      false,
		FreeToPlay:        false,
		ID:                r1.Int63(),
		RankedPlayEnabled: false,
	}
}
