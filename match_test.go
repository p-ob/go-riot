package lolgo

import (
	"math/rand"
	"time"
)

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
