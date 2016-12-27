package lolgo

import (
	"context"
	"fmt"
)

type StatsService struct {
	client *Client
}

type RankedStatsDto struct {
	Champions  []ChampionStatsDto `json:"champions"`
	ModifyDate int64              `json:"modifyDate"`
	SummonerId int64              `json:"summonerId"`
}

type ChampionStatsDto struct {
	Id    int                `json:"id"`
	Stats AggregatedStatsDto `json:"stats"`
}

type AggregatedStatsDto struct {
	AverageAssists              int `json:"averageAssists"`
	AverageChampionsKilled      int `json:"averageChampionsKilled"`
	AverageCombatPlayerScore    int `json:"averageCombatPlayerScore"`
	AverageNodeCapture          int `json:"averageNodeCapture"`
	AverageNodeCaptureAssist    int `json:"averageNodeCaptureAssist"`
	AverageNodeNeutralize       int `json:"averageNodeNeutralize"`
	AverageNodeNeutralizeAssist int `json:"averageNodeNeutralizeAssist"`
	AverageNumDeaths            int `json:"averageNumDeaths"`
	AverageObjectivePlayerScore int `json:"averageObjectivePlayerScore"`
	AverageTeamObjective        int `json:"averageTeamObjective"`
	AverageTotalPlayerScore     int `json:"averageTotalPlayerScore"`
	BotGamesPlayed              int `json:"botGamesPlayed"`
	KillingSpree                int `json:"killingSpree"`
	MaxAssists                  int `json:"maxAssists"`
	MaxChampionsKilled          int `json:"maxChampionsKilled"`
	MaxCombatPlayerScore        int `json:"maxCombatPlayerScore"`
	MaxLargestCriticalStrike    int `json:"maxLargestCriticalStrike"`
	MaxLargestKillingSpree      int `json:"maxLargestKillingSpree"`
	MaxNodeCapture              int `json:"maxNodeCapture"`
	MaxNodeCaptureAssist        int `json:"maxNodeCaptureAssist"`
	MaxNodeNeutralize           int `json:"maxNodeNeutralize"`
	MaxNodeNeutralizeAssist     int `json:"maxNodeNeutralizeAssist"`
	MaxNumDeaths                int `json:"maxNumDeaths"`
	MaxObjectivePlayerScore     int `json:"maxObjectivePlayerScore"`
	MaxTeamObjective            int `json:"maxTeamObjective"`
	MaxTimePlayed               int `json:"maxTimePlayed"`
	MaxTimeSpentLiving          int `json:"maxTimeSpentLiving"`
	MaxTotalPlayerScore         int `json:"maxTotalPlayerScore"`
	MostChampionKillsPerSession int `json:"mostChampionKillsPerSession"`
	MostSpellsCast              int `json:"mostSpellsCast"`
	NormalGamesPlayed           int `json:"normalGamesPlayed"`
	RankedPremadeGamesPlayed    int `json:"rankedPremadeGamesPlayed"`
	RankedSoloGamesPlayed       int `json:"rankedSoloGamesPlayed"`
	TotalAssists                int `json:"totalAssists"`
	TotalChampionKills          int `json:"totalChampionKills"`
	TotalDamageDealt            int `json:"totalDamageDealt"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TotalDeathsPerSession       int `json:"totalDeathsPerSession"`
	TotalDoubleKills            int `json:"totalDoubleKills"`
	TotalFirstBlood             int `json:"totalFirstBlood"`
	TotalGoldEarned             int `json:"totalGoldEarned"`
	TotalHeal                   int `json:"totalHeal"`
	TotalMagicDamageDealt       int `json:"totalMagicDamageDealt"`
	TotalMinionKills            int `json:"totalMinionKills"`
	TotalNeutralMinionsKilled   int `json:"totalNeutralMinionsKilled"`
	TotalNodeCapture            int `json:"totalNodeCapture"`
	TotalNodeNeutralize         int `json:"totalNodeNeutralize"`
	TotalPentaKills             int `json:"totalPentaKills"`
	TotalPhysicalDamageDealt    int `json:"totalPhysicalDamageDealt"`
	TotalQuadraKills            int `json:"totalQuadraKills"`
	TotalSessionsLost           int `json:"totalSessionsLost"`
	TotalSessionsPlayed         int `json:"totalSessionsPlayed"`
	TotalSessionsWon            int `json:"totalSessionsWon"`
	TotalTripleKills            int `json:"totalTripleKills"`
	TotalTurretsKilled          int `json:"totalTurretsKilled"`
	TotalUnrealKills            int `json:"totalUnrealKills"`
}

type PlayerStatsSummaryListDto struct {
	PlayerStatSummaries []PlayerStatsSummaryDto `json:"playerStatSummaries"`
	SummonerId          int64                   `json:"summonerId"`
}

type PlayerStatsSummaryDto struct {
	AggregatedStats       AggregatedStatsDto `json:"aggregatedStats"`
	Losses                int                `json:"losses"`
	ModifyDate            int64              `json:"modifyDate"`
	PlayerStatSummaryType string             `json:"playerStatSummaryType"`
	Wins                  int                `json:"wins"`
}

type GetStatsParams struct {
	Season string `url:"season,omitempty"`
}

const statsPathPart = "/api/lol/{region}/v1.3/stats/"

func (s *StatsService) GetRankedBySummoner(ctx context.Context, summonerId int64, params *GetStatsParams) (*RankedStatsDto, error) {
	stats := new(RankedStatsDto)

	err := s.client.GetResource(
		ctx,
		fmt.Sprintf(statsPathPart+"/%s/ranked", summonerId),
		"",
		params,
		stats)
	return stats, err
}

func (s *StatsService) GetSummaryBySummoner(ctx context.Context, summonerId int64, params *GetStatsParams) (*PlayerStatsSummaryListDto, error) {
	stats := new(PlayerStatsSummaryListDto)

	err := s.client.GetResource(
		ctx,
		fmt.Sprintf(statsPathPart+"/%s/summary", summonerId),
		"",
		params,
		stats)
	return stats, err
}
