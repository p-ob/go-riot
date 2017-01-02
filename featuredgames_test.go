package lolgo

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestFeaturedGamesService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	featuredGames := generateFeaturedGames()
	pathPart := "/" + featuredGamePathPart

	jsonByteArray, _ := json.Marshal(featuredGames)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	games, err := client.FeaturedGames.Get(ctx)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(featuredGames, *games) {
		t.Errorf("expected %+v, got %+v", featuredGames, *games)
	}
}

func generateFeaturedGames() FeaturedGames {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return FeaturedGames{
		ClientRefreshInterval: r1.Int63(),
		GameList:              []CurrentGameInfo{generateCurrentGameInfo()},
	}
}
