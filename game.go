package lolgo

import (
	"context"
	"fmt"
)

type GameService struct {
	client *Client
}

type RecentGamesDto struct {
	Games      []GameDto `json:"games"`
	SummonerId int64     `json:"summonerId"`
}

type GameDto struct {
	ChampionId    int         `json:"championId"`
	CreateDate    int64       `json:"createDate"`
	FellowPlayers []PlayerDto `json:"fellowPlayers"`
	GameId        int64       `json:"gameId"`
	GameMode      string      `json:"gameMode"`
	GameType      string      `json:"gameType"`
	Invalid       bool        `json:"invalid"`
	IpEarned      int         `json:"ipEarned"`
	Level         int         `json:"level"`
	MapId         int         `json:"mapId"`
	Spell1        int         `json:"spell1"`
	Spell2        int         `json:"spell2"`
	Stats         RawStatsDto `json:"stats"`
	SubType       string      `json:"subType"`
	TeamId        int         `json:"teamId"`
}

type PlayerDto struct {
	ChampionId int   `json:"championId"`
	SummonerId int64 `json:"summonerId"`
	TeamId     int   `json:"teamId"`
}

type RawStatsDto struct {
	Assists                         int  `json:"assists"`
	BarracksKilled                  int  `json:"barracksKilled"`
	ChampionsKilled                 int  `json:"championsKilled"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	ConsumablesPurchased            int  `json:"consumablesPurchased"`
	DamageDealtPlayer               int  `json:"damageDealtPlayer"`
	DoubleKills                     int  `json:"doubleKills"`
	FirstBlood                      int  `json:"firstBlood"`
	Gold                            int  `json:"gold"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	ItemsPurchased                  int  `json:"itemsPurchased"`
	KillingSprees                   int  `json:"killingSprees"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LegendaryItemsCreated           int  `json:"legendaryItemsCreated"`
	Level                           int  `json:"level"`
	MagicDamageDealtPlayer          int  `json:"magicDamageDealtPlayer"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int  `json:"magicDamageTaken"`
	MinionsDenied                   int  `json:"minionsDenied"`
	MinionsKilled                   int  `json:"minionsKilled"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledYourJungle  int  `json:"neutralMinionsKilledYourJungle"`
	NexusKilled                     bool `json:"nexusKilled"`
	NodeCapture                     int  `json:"nodeCapture"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	NumDeaths                       int  `json:"numDeaths"`
	NumItemsBought                  int  `json:"numItemsBought"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	PentaKills                      int  `json:"pentaKills"`
	PhysicalDamageDealtPlayer       int  `json:"physicalDamageDealtPlayer"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerPosition                  int  `json:"playerPosition"`
	PlayerRole                      int  `json:"playerRole"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBought                int  `json:"sightWardsBought"`
	Spell1Cast                      int  `json:"spell1Cast"`
	Spell2Cast                      int  `json:"spell2Cast"`
	Spell3Cast                      int  `json:"spell3Cast"`
	Spell4Cast                      int  `json:"spell4Cast"`
	SummonerSpell1Cast              int  `json:"summonerSpell1Cast"`
	SummonerSpell2Cast              int  `json:"summonerSpell2Cast"`
	SuperMonsterKilled              int  `json:"superMonsterKilled"`
	Team                            int  `json:"team"`
	TeamObjective                   int  `json:"teamObjective"`
	TimePlayed                      int  `json:"timePlayed"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealtPlayer           int  `json:"trueDamageDealtPlayer"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretsKilled                   int  `json:"turretsKilled"`
	UnrealKills                     int  `json:"unrealKills"`
	VictoryPointTotal               int  `json:"victoryPointTotal"`
	VisionWardsBought               int  `json:"visionWardsBought"`
	WardKilled                      int  `json:"wardKilled"`
	WardPlaced                      int  `json:"wardPlaced"`
	Win                             bool `json:"win"`
}

const gamePathPart = "api/lol/{region}/v1.3/game"

func (s *GameService) GetRecent(ctx context.Context, summonerId int64) (*FeaturedGames, error) {
	featuredGames := new(FeaturedGames)

	err := s.client.GetResource(
		ctx,
		fmt.Sprintf(gamePathPart+"/%s/recent", summonerId),
		"",
		nil,
		featuredGames)
	return featuredGames, err
}
