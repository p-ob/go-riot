package main

type RankedStatsDto struct {
	Champions  []ChampionStatsDto
	ModifyDate int64
	SummonerId int64
}

type ChampionStatsDto struct {
	Id    int
	Stats AggregatedStatsDto
}

type AggregatedStatsDto struct {
	AverageAssists              int
	AverageChampionsKilled      int
	AverageCombatPlayerScore    int
	AverageNodeCapture          int
	AverageNodeCaptureAssist    int
	AverageNodeNeutralize       int
	AverageNodeNeutralizeAssist int
	AverageNumDeaths            int
	AverageObjectivePlayerScore int
	AverageTeamObjective        int
	AverageTotalPlayerScore     int
	BotGamesPlayed              int
	KillingSpree                int
	MaxAssists                  int
	MaxChampionsKilled          int
	MaxCombatPlayerScore        int
	MaxLargestCriticalStrike    int
	MaxLargestKillingSpree      int
	MaxNodeCapture              int
	MaxNodeCaptureAssist        int
	MaxNodeNeutralize           int
	MaxNodeNeutralizeAssist     int
	MaxNumDeaths                int
	MaxObjectivePlayerScore     int
	MaxTeamObjective            int
	MaxTimePlayed               int
	MaxTimeSpentLiving          int
	MaxTotalPlayerScore         int
	MostChampionKillsPerSession int
	MostSpellsCast              int
	NormalGamesPlayed           int
	RankedPremadeGamesPlayed    int
	RankedSoloGamesPlayed       int
	TotalAssists                int
	TotalChampionKills          int
	TotalDamageDealt            int
	TotalDamageTaken            int
	TotalDeathsPerSession       int
	TotalDoubleKills            int
	TotalFirstBlood             int
	TotalGoldEarned             int
	TotalHeal                   int
	TotalMagicDamageDealt       int
	TotalMinionKills            int
	TotalNeutralMinionsKilled   int
	TotalNodeCapture            int
	TotalNodeNeutralize         int
	TotalPentaKills             int
	TotalPhysicalDamageDealt    int
	TotalQuadraKills            int
	TotalSessionsLost           int
	TotalSessionsPlayed         int
	TotalSessionsWon            int
	TotalTripleKills            int
	TotalTurretsKilled          int
	TotalUnrealKills            int
}

type PlayerStatsSummaryListDto struct {
	PlayerStatSummaries []PlayerStatsSummaryDto
	SummonerId          int64
}

type PlayerStatsSummaryDto struct {
	AggregatedStats       AggregatedStatsDto
	Losses                int
	ModifyDate            int64
	PlayerStatSummaryType string
	Wins                  int
}
