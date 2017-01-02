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

func TestGameService_GetRecent(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	games := generateRecentGamesDto()
	pathPart := fmt.Sprintf("/%s/by-summoner/%v/recent", addRegionToString(gamePathPart, region), games.SummonerID)

	jsonByteArray, _ := json.Marshal(games)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedGames, err := client.Game.GetRecent(ctx, games.SummonerID)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(games, *retrievedGames) {
		t.Errorf("expected %+v, got %+v", games, *retrievedGames)
	}
}

func generateRecentGamesDto() RecentGamesDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return RecentGamesDto{
		SummonerID: r1.Int63(),
		Games:      []GameDto{generateGameDto()},
	}
}

func generateGameDto() GameDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return GameDto{
		ChampionID:    r1.Int(),
		CreateDate:    r1.Int63(),
		FellowPlayers: []PlayerDto{generatePlayerDto()},
		GameID:        r1.Int63(),
		GameMode:      randString(10),
		GameType:      randString(10),
		Invalid:       false,
		IPEarned:      r1.Int(),
		Level:         r1.Int(),
		MapID:         r1.Int(),
		Spell1:        r1.Int(),
		Spell2:        r1.Int(),
		Stats:         generateRawStatsDto(),
		SubType:       randString(10),
		TeamID:        r1.Int(),
	}
}

func generatePlayerDto() PlayerDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return PlayerDto{
		ChampionID: r1.Int(),
		SummonerID: r1.Int63(),
		TeamID:     r1.Int(),
	}
}

func generateRawStatsDto() RawStatsDto {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return RawStatsDto{
		Assists:               r1.Int(),
		BarracksKilled:        r1.Int(),
		ChampionsKilled:       r1.Int(),
		CombatPlayerScore:     r1.Int(),
		ConsumablesPurchased:  r1.Int(),
		DamageDealtPlayer:     r1.Int(),
		DoubleKills:           r1.Int(),
		FirstBlood:            r1.Int(),
		Gold:                  r1.Int(),
		GoldEarned:            r1.Int(),
		GoldSpent:             r1.Int(),
		Item0:                 r1.Int(),
		Item1:                 r1.Int(),
		Item2:                 r1.Int(),
		Item3:                 r1.Int(),
		Item4:                 r1.Int(),
		Item5:                 r1.Int(),
		Item6:                 r1.Int(),
		ItemsPurchased:        r1.Int(),
		KillingSprees:         r1.Int(),
		LargestCriticalStrike: r1.Int(),
		LargestKillingSpree:   r1.Int(),
		LargestMultiKill:      r1.Int(),
		LegendaryItemsCreated: r1.Int(),
		Level: r1.Int(),
		MagicDamageDealtPlayer:          r1.Int(),
		MagicDamageDealtToChampions:     r1.Int(),
		MagicDamageTaken:                r1.Int(),
		MinionsDenied:                   r1.Int(),
		MinionsKilled:                   r1.Int(),
		NeutralMinionsKilled:            r1.Int(),
		NeutralMinionsKilledEnemyJungle: r1.Int(),
		NeutralMinionsKilledYourJungle:  r1.Int(),
		NexusKilled:                     false,
		NodeCapture:                     r1.Int(),
		NodeCaptureAssist:               r1.Int(),
		NodeNeutralize:                  r1.Int(),
		NodeNeutralizeAssist:            r1.Int(),
		NumDeaths:                       r1.Int(),
		NumItemsBought:                  r1.Int(),
		ObjectivePlayerScore:            r1.Int(),
		PentaKills:                      r1.Int(),
		PhysicalDamageDealtPlayer:       r1.Int(),
		PhysicalDamageDealtToChampions:  r1.Int(),
		PhysicalDamageTaken:             r1.Int(),
		PlayerPosition:                  r1.Int(),
		PlayerRole:                      r1.Int(),
		QuadraKills:                     r1.Int(),
		SightWardsBought:                r1.Int(),
		Spell1Cast:                      r1.Int(),
		Spell2Cast:                      r1.Int(),
		Spell3Cast:                      r1.Int(),
		Spell4Cast:                      r1.Int(),
		SummonerSpell1Cast:              r1.Int(),
		SummonerSpell2Cast:              r1.Int(),
		SuperMonsterKilled:              r1.Int(),
		Team:                            r1.Int(),
		TeamObjective:                   r1.Int(),
		TimePlayed:                      r1.Int(),
		TotalDamageDealt:                r1.Int(),
		TotalDamageDealtToChampions:     r1.Int(),
		TotalDamageTaken:                r1.Int(),
		TotalHeal:                       r1.Int(),
		TotalPlayerScore:                r1.Int(),
		TotalScoreRank:                  r1.Int(),
		TotalTimeCrowdControlDealt:      r1.Int(),
		TotalUnitsHealed:                r1.Int(),
		TripleKills:                     r1.Int(),
		TrueDamageDealtPlayer:           r1.Int(),
		TrueDamageDealtToChampions:      r1.Int(),
		TrueDamageTaken:                 r1.Int(),
		TurretsKilled:                   r1.Int(),
		UnrealKills:                     r1.Int(),
		VictoryPointTotal:               r1.Int(),
		VisionWardsBought:               r1.Int(),
		WardKilled:                      r1.Int(),
		WardPlaced:                      r1.Int(),
		Win:                             false,
	}
}
