package lolgo

import (
	"context"
	"fmt"
	"strconv"
)

type CurrentGameService struct {
	client *Client
}

type CurrentGameInfo struct {
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	GameId            int64                    `json:"gameId"`
	GameLength        int64                    `json:"gameLength"`
	GameMode          string                   `json:"gameMode"`
	GameQueueConfigId int64                    `json:"gameQueueConfigId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	GameType          string                   `json:"gameType"`
	MapId             int64                    `json:"mapId"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	PlatformId        string                   `json:"platformId"`
}

type CurrentGameParticipant struct {
	Bot           bool              `json:"bot"`
	ChampionId    int64             `json:"championId"`
	Masteries     []Mastery         `json:"masteries"`
	ProfileIconId int64             `json:"profileIconId"`
	Runes         []CurrentGameRune `json:"runes"`
	Spell1Id      int64             `json:"spell1Id"`
	Spell2Id      int64             `json:"spell2Id"`
	SummonerId    int64             `json:"summonerId"`
	SummonerName  string            `json:"summonerName"`
	TeamId        int64             `json:"teamId"`
}

type CurrentGameRune struct {
	Count  int   `json:"count"`
	RuneId int64 `json:"runeId"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

const currentGamePathPart = "observer-mode/rest/consumer/getSpectatorGameInfo/%s"

func (s *CurrentGameService) Get(ctx context.Context, summonerId int64) (*CurrentGameInfo, error) {
	currentGameInfo := new(CurrentGameInfo)

	err := s.client.getResource(
		ctx,
		fmt.Sprintf(currentGamePathPart, mapRegionToLocationString(s.client.Region)),
		strconv.FormatInt(summonerId, 10),
		nil,
		currentGameInfo)
	return currentGameInfo, err
}
