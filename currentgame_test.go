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

func TestCurrentGameService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	summonerID := r1.Int63()
	currentGame := generateCurrentGameInfo()
	pathPart := fmt.Sprintf(currentGamePathPart, mapRegionToLocationString(region))
	pathPart = fmt.Sprintf(
		"/%s/%v",
		pathPart,
		summonerID,
	)

	jsonByteArray, _ := json.Marshal(currentGame)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	game, err := client.CurrentGame.Get(ctx, summonerID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(currentGame, *game) {
		t.Errorf("expected %+v, got %+v", currentGame, *game)
	}
}

func generateCurrentGameInfo() CurrentGameInfo {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return CurrentGameInfo{
		BannedChampions:   []BannedChampion{generateBannedChampion()},
		GameID:            r1.Int63(),
		GameLength:        r1.Int63(),
		GameMode:          randString(10),
		GameQueueConfigID: r1.Int63(),
		GameStartTime:     r1.Int63(),
		GameType:          randString(10),
		MapID:             r1.Int63(),
		Observers:         generateObserver(),
		Participants:      []CurrentGameParticipant{generateCurrentGameParticipant()},
		PlatformID:        randString(3),
	}
}

func generateObserver() Observer {
	return Observer{
		EncryptionKey: randString(10),
	}
}

func generateCurrentGameParticipant() CurrentGameParticipant {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return CurrentGameParticipant{
		Bot:           false,
		ChampionID:    r1.Int63(),
		Masteries:     []Mastery{generateMastery()},
		ProfileIconID: r1.Int63(),
		Runes:         []CurrentGameRune{generateCurrentGameRune()},
		Spell1ID:      r1.Int63(),
		Spell2ID:      r1.Int63(),
		SummonerID:    r1.Int63(),
		SummonerName:  randString(10),
		TeamID:        r1.Int63(),
	}
}

func generateCurrentGameRune() CurrentGameRune {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return CurrentGameRune{
		Count:  r1.Int(),
		RuneID: r1.Int63(),
	}
}
