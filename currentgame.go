package lolgo

import (
	"context"
	"fmt"
	"strconv"
)

// CurrentGameService is the endpoint to use to get current game data
type CurrentGameService struct {
	client *Client
}

// CurrentGameInfo is the container returned for a current game
type CurrentGameInfo struct {
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	GameID            int64                    `json:"gameId"`
	GameLength        int64                    `json:"gameLength"`
	GameMode          string                   `json:"gameMode"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	GameType          string                   `json:"gameType"`
	MapID             int64                    `json:"mapId"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	PlatformID        string                   `json:"platformId"`
}

// CurrentGameParticipant is a participant's information for the current game
type CurrentGameParticipant struct {
	Bot           bool              `json:"bot"`
	ChampionID    int64             `json:"championId"`
	Masteries     []Mastery         `json:"masteries"`
	ProfileIconID int64             `json:"profileIconId"`
	Runes         []CurrentGameRune `json:"runes"`
	Spell1ID      int64             `json:"spell1Id"`
	Spell2ID      int64             `json:"spell2Id"`
	SummonerID    int64             `json:"summonerId"`
	SummonerName  string            `json:"summonerName"`
	TeamID        int64             `json:"teamId"`
}

// CurrentGameRune is a rune used by a current game participant
type CurrentGameRune struct {
	Count  int   `json:"count"`
	RuneID int64 `json:"runeId"`
}

// Observer supplies spectate game observer information
type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

const currentGamePathPart = "observer-mode/rest/consumer/getSpectatorGameInfo/%s"

// Get gets the current game for a summonerID
func (s *CurrentGameService) Get(ctx context.Context, summonerID int64) (*CurrentGameInfo, error) {
	currentGameInfo := new(CurrentGameInfo)

	err := s.client.getResource(
		ctx,
		fmt.Sprintf(currentGamePathPart, mapRegionToLocationString(s.client.region)),
		strconv.FormatInt(summonerID, 10),
		nil,
		currentGameInfo,
	)
	return currentGameInfo, err
}
