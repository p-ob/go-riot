package main

type RecentGamesDto struct {
	Games      []GameDto
	SummonerId int64
}

type GameDto struct {
	ChampionId    int
	CreateDate    int64
	FellowPlayers []PlayerDto
	GameId        int64
	GameType      string
	Invalid       bool
	IpEarned      int
	Level         int
	MapId         int
	Spell1        int
	Spell2        int
	Stats         RawStatsDto
	SubType       string
	TeamId        int
}

type PlayerDto struct {
	ChampionId int
	SummonerId int64
	TeamId     int
}

type RawStatsDto struct {
	Assists                         int
	BarracksKilled                  int
	ChampionsKilled                 int
	CombatPlayerScore               int
	ConsumablesPurchased            int
	DamageDealtPlayer               int
	DoubleKills                     int
	FirstBlood                      int
	Gold                            int
	GoldEarned                      int
	GoldSpent                       int
	Item0                           int
	Item1                           int
	Item2                           int
	Item3                           int
	Item4                           int
	Item5                           int
	Item6                           int
	ItemsPurchased                  int
	KillingSprees                   int
	LargestCriticalStrike           int
	LargestKillingSpree             int
	LargestMultiKill                int
	LegendaryItemsCreated           int
	Level                           int
	MagicDamageDealtPlayer          int
	MagicDamageDealtToChampions     int
	MagicDamageTaken                int
	MinionsDenied                   int
	MinionsKilled                   int
	NeutralMinionsKilled            int
	NeutralMinionsKilledEnemyJungle int
	NeutralMinionsKilledYourJungle  int
	NexusKilled                     bool
	NodeCapture                     int
	NodeCaptureAssist               int
	NodeNeutralize                  int
	NodeNeutralizeAssist            int
	NumDeaths                       int
	NumItemsBought                  int
	ObjectivePlayerScore            int
	PentaKills                      int
	PhysicalDamageDealtPlayer       int
	PhysicalDamageDealtToChampions  int
	PhysicalDamageTaken             int
	PlayerPosition                  int
	PlayerRole                      int
	QuadraKills                     int
	SightWardsBought                int
	Spell1Cast                      int
	Spell2Cast                      int
	Spell3Cast                      int
	Spell4Cast                      int
	SummonerSpell1Cast              int
	SummonerSpell2Cast              int
	SuperMonsterKilled              int
	Team                            int
	TeamObjective                   int
	TimePlayed                      int
	TotalDamageDealt                int
	TotalDamageDealtToChampions     int
	TotalDamageTaken                int
	TotalHeal                       int
	TotalPlayerScore                int
	TotalScoreRank                  int
	TotalTimeCrowdControlDealt      int
	TotalUnitsHealed                int
	TripleKills                     int
	TrueDamageDealtPlayer           int
	TrueDamageDealtToChampions      int
	TrueDamageTaken                 int
	TurretsKilled                   int
	UnrealKills                     int
	VictoryPointTotal               int
	VisionWardsBought               int
	WardKilled                      int
	WardPlaced                      int
	Win                             bool
}
