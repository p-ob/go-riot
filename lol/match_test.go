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

func TestMatchService_Get(t *testing.T) {
	// set up data
	BaseURL = "http://example.com"
	region := Na
	match := generateMatchDetail()
	pathPart := fmt.Sprintf("/%s/%v", addRegionToString(matchPathPart, region), match.MatchID)

	jsonByteArray, _ := json.Marshal(match)

	_, mux, server, client := mockClient(region)
	defer server.Close()

	mux.HandleFunc(pathPart, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonByteArray)
	})

	ctx := context.Background()
	retrievedMatch, err := client.Match.Get(ctx, match.MatchID, nil)
	if err != nil {
		t.Errorf("expected nil, got %+v", err)
	}
	if !reflect.DeepEqual(match, *retrievedMatch) {
		t.Errorf("expected %+v, got %+v", match, *retrievedMatch)
	}
}

func generateBannedChampion() BannedChampion {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return BannedChampion{
		ChampionID: r1.Int63(),
		PickTurn:   r1.Int(),
		TeamID:     r1.Int63(),
	}
}

func generateMastery() Mastery {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Mastery{
		MasteryID: r1.Int63(),
		Rank:      r1.Int(),
	}
}

func generateMatchDetail() MatchDetail {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return MatchDetail{
		MapID:                 r1.Int(),
		MatchCreation:         r1.Int63(),
		MatchDuration:         r1.Int63(),
		MatchID:               r1.Int63(),
		MatchMode:             randString(10),
		MatchType:             randString(10),
		MatchVersion:          randString(10),
		ParticipantIdentities: []ParticipantIdentity{generateParticipantIdentity()},
		Participants:          []Participant{generateParticipant()},
		QueueType:             randString(10),
		Region:                randString(3),
		Season:                randString(10),
		Teams:                 []Team{generateTeam()},
		Timeline:              generateTimeline(),
	}
}

func generateParticipantIdentity() ParticipantIdentity {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return ParticipantIdentity{
		ParticipantID: r1.Int(),
		Player:        generatePlayer(),
	}
}

func generateParticipant() Participant {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Participant{
		ChampionID:                r1.Int(),
		HighestAchievedSeasonTier: randString(10),
		Masteries:                 []Mastery{generateMastery()},
		ParticipantID:             r1.Int(),
		Runes:                     []Rune{generateRune()},
		Spell1ID:                  r1.Int(),
		Spell2ID:                  r1.Int(),
		Stats:                     generateParticipantStats(),
		TeamID:                    r1.Int(),
		Timeline:                  generateParticipantTimeline(),
	}
}

func generateTeam() Team {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Team{
		Bans:                 []BannedChampion{generateBannedChampion()},
		BaronKills:           r1.Int(),
		DominionVictoryScore: r1.Int63(),
		DragonKills:          r1.Int(),
		FirstBaron:           false,
		FirstBlood:           false,
		FirstDragon:          false,
		FirstInhibitor:       false,
		FirstRiftHerald:      false,
		FirstTower:           false,
		InhibitorKills:       r1.Int(),
		RiftHeraldKills:      r1.Int(),
		TeamID:               r1.Int(),
		TowerKills:           r1.Int(),
		VilemawKills:         r1.Int(),
		Winner:               false,
	}
}

func generateTimeline() Timeline {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Timeline{
		FrameInterval: r1.Int63(),
		Frames:        []Frame{generateFrame()},
	}
}

func generatePlayer() Player {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Player{
		MatchHistoryURI: randString(10),
		ProfileIcon:     r1.Int(),
		SummonerID:      r1.Int63(),
		SummonerName:    randString(10),
	}
}

func generateRune() Rune {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Rune{
		Rank:   r1.Int63(),
		RuneID: r1.Int63(),
	}
}

func generateParticipantStats() ParticipantStats {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return ParticipantStats{
		Assists:              r1.Int63(),
		ChampLevel:           r1.Int63(),
		CombatPlayerScore:    r1.Int63(),
		Deaths:               r1.Int63(),
		DoubleKills:          r1.Int63(),
		FirstBloodAssist:     false,
		FirstBloodKill:       false,
		FirstInhibitorAssist: false,
		FirstInhibitorKill:   false,
		FirstTowerAssist:     false,
		FirstTowerKill:       false,
		GoldEarned:           r1.Int63(),
		GoldSpent:            r1.Int63(),
		InhibitorKills:       r1.Int63(),
		Item0:                r1.Int63(),
		Item1:                r1.Int63(),
		Item2:                r1.Int63(),
		Item3:                r1.Int63(),
		Item4:                r1.Int63(),
		Item5:                r1.Int63(),
		Item6:                r1.Int63(),
		KillingSprees:        r1.Int63(),
		Kills:                r1.Int63(),
		LargestCriticalStrike:           r1.Int63(),
		LargestKillingSpree:             r1.Int63(),
		LargestMultiKill:                r1.Int63(),
		MagicDamageDealt:                r1.Int63(),
		MagicDamageDealtToChampions:     r1.Int63(),
		MagicDamageTaken:                r1.Int63(),
		MinionsKilled:                   r1.Int63(),
		NeutralMinionsKilled:            r1.Int63(),
		NeutralMinionsKilledEnemyJungle: r1.Int63(),
		NeutralMinionsKilledTeamJungle:  r1.Int63(),
		NodeCapture:                     r1.Int63(),
		NodeCaptureAssist:               r1.Int63(),
		NodeNeutralize:                  r1.Int63(),
		ObjectivePlayerScore:            r1.Int63(),
		PentaKills:                      r1.Int63(),
		PhysicalDamageDealt:             r1.Int63(),
		PhysicalDamageDealtToChampions:  r1.Int63(),
		PhysicalDamageTaken:             r1.Int63(),
		QuadraKills:                     r1.Int63(),
		SightWardsBoughtInGame:          r1.Int63(),
		TeamObjective:                   r1.Int63(),
		TotalDamageDealt:                r1.Int63(),
		TotalDamageDealtToChampions:     r1.Int63(),
		TotalDamageTaken:                r1.Int63(),
		TotalHeal:                       r1.Int63(),
		TotalPlayerScore:                r1.Int63(),
		TotalScoreRank:                  r1.Int63(),
		TotalTimeCrowdControlDealt:      r1.Int63(),
		TotalUnitsHealed:                r1.Int63(),
		TowerKills:                      r1.Int63(),
		TripleKills:                     r1.Int63(),
		TrueDamageDealt:                 r1.Int63(),
		TrueDamageDealtToChampions:      r1.Int63(),
		TrueDamageTaken:                 r1.Int63(),
		UnrealKills:                     r1.Int63(),
		VisionWardsBoughtInGame:         r1.Int63(),
		WardsKilled:                     r1.Int63(),
		WardsPlaced:                     r1.Int63(),
		Winner:                          false,
	}
}

func generateParticipantTimeline() ParticipantTimeline {
	return ParticipantTimeline{
		AncientGolemAssistsPerMinCounts: generateParticipantTimelineData(),
		AncientGolemKillsPerMinCounts:   generateParticipantTimelineData(),
		AssistedLaneDeathsPerMinDeltas:  generateParticipantTimelineData(),
		AssistedLaneKillsPerMinDeltas:   generateParticipantTimelineData(),
		BaronAssistsPerMinCounts:        generateParticipantTimelineData(),
		BaronKillsPerMinCounts:          generateParticipantTimelineData(),
		CreepsPerMinDeltas:              generateParticipantTimelineData(),
		CsDiffPerMinDeltas:              generateParticipantTimelineData(),
		DamageTakenDiffPerMinDeltas:     generateParticipantTimelineData(),
		DamageTakenPerMinDeltas:         generateParticipantTimelineData(),
		DragonAssistsPerMinCounts:       generateParticipantTimelineData(),
		DragonKillsPerMinCounts:         generateParticipantTimelineData(),
		ElderLizardAssistsPerMinCounts:  generateParticipantTimelineData(),
		ElderLizardKillsPerMinCounts:    generateParticipantTimelineData(),
		GoldPerMinDeltas:                generateParticipantTimelineData(),
		InhibitorAssistsPerMinCounts:    generateParticipantTimelineData(),
		InhibitorKillsPerMinCounts:      generateParticipantTimelineData(),
		Lane: randString(10),
		Role: randString(10),
		TowerAssistsPerMinCounts:   generateParticipantTimelineData(),
		TowerKillsPerMinCounts:     generateParticipantTimelineData(),
		TowerKillsPerMinDeltas:     generateParticipantTimelineData(),
		VilemawAssistsPerMinCounts: generateParticipantTimelineData(),
		VilemawKillsPerMinCounts:   generateParticipantTimelineData(),
		WardsPerMinDeltas:          generateParticipantTimelineData(),
		XpDiffPerMinDeltas:         generateParticipantTimelineData(),
		XpPerMinDeltas:             generateParticipantTimelineData(),
	}
}

func generateParticipantTimelineData() ParticipantTimelineData {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return ParticipantTimelineData{
		TenToTwenty:    r1.Float64(),
		ThirtyToTen:    r1.Float64(),
		TwentyToThirty: r1.Float64(),
		ZeroToTen:      r1.Float64(),
	}
}

func generateFrame() Frame {
	pFrame := generateParticipantFrame()

	return Frame{
		Events:            []Event{generateEvent()},
		ParticipantFrames: map[int]ParticipantFrame{pFrame.ParticipantID: pFrame},
	}
}

func generateEvent() Event {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Event{
		AscendedType:            randString(10),
		AssistingParticipantIDs: []int{r1.Int()},
		BuildingType:            randString(10),
		CreatorID:               r1.Int(),
		EventType:               randString(10),
		ItemAfter:               r1.Int(),
		ItemBefore:              r1.Int(),
		ItemID:                  r1.Int(),
		KillerID:                r1.Int(),
		LaneType:                randString(10),
		LevelUpType:             randString(10),
		MonsterType:             randString(10),
		ParticipantID:           r1.Int(),
		PointCaptured:           randString(10),
		Position:                generatePosition(),
		SkillSlot:               r1.Int(),
		TeamID:                  r1.Int(),
		Timestamp:               r1.Int63(),
		TowerType:               randString(10),
		VictimID:                r1.Int(),
		WardType:                randString(10),
	}
}

func generateParticipantFrame() ParticipantFrame {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return ParticipantFrame{
		CurrentGold:         r1.Int(),
		DominionScore:       r1.Int(),
		JungleMinionsKilled: r1.Int(),
		Level:               r1.Int(),
		MinionsKilled:       r1.Int(),
		ParticipantID:       r1.Int(),
		Position:            generatePosition(),
		TeamScore:           r1.Int(),
		TotalGold:           r1.Int(),
		Xp:                  r1.Int(),
	}
}

func generatePosition() Position {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Position{
		X: r1.Int(),
		Y: r1.Int(),
	}
}
