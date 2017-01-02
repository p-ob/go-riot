package lol

import (
	"context"
	"strconv"
)

// MatchService is the endpoint to use to get detailed match information
type MatchService struct {
	client *Client
}

// MatchDetail is the container for a given match information
type MatchDetail struct {
	MapID                 int                   `json:"mapId"`
	MatchCreation         int64                 `json:"matchCreation"`
	MatchDuration         int64                 `json:"matchDuration"`
	MatchID               int64                 `json:"matchId"`
	MatchMode             string                `json:"matchMode"`
	MatchType             string                `json:"matchType"`
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
	Teams                 []Team                `json:"teams"`
	Timeline              Timeline              `json:"timeline"`
}

// ParticipantIdentity is the container for a given participant in a match
type ParticipantIdentity struct {
	ParticipantID int    `json:"participantId"`
	Player        Player `json:"player"`
}

// Player is the container with summoner information for a given match
type Player struct {
	MatchHistoryURI string `json:"matchHistoryUri"`
	ProfileIcon     int    `json:"profileIcon"`
	SummonerID      int64  `json:"summonerId"`
	SummonerName    string `json:"summonerName"`
}

// Participant is the container for participant information in a game
type Participant struct {
	ChampionID                int                 `json:"championId"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
	Masteries                 []Mastery           `json:"masteries"`
	ParticipantID             int                 `json:"participantId"`
	Runes                     []Rune              `json:"runes"`
	Spell1ID                  int                 `json:"spell1Id"`
	Spell2ID                  int                 `json:"spell2Id"`
	Stats                     ParticipantStats    `json:"stats"`
	TeamID                    int                 `json:"teamId"`
	Timeline                  ParticipantTimeline `json:"timeline"`
}

// Mastery is the container for basic mastery utilization information
type Mastery struct {
	MasteryID int64 `json:"masteryId"`
	Rank      int   `json:"rank"`
}

// Rune is the cotainer for basic rune utilization information
type Rune struct {
	Rank   int64 `json:"rank"`
	RuneID int64 `json:"runeId"`
}

// ParticipantStats is the detailed information for a participant
type ParticipantStats struct {
	Assists                         int64 `json:"assists"`
	ChampLevel                      int64 `json:"champLevel"`
	CombatPlayerScore               int64 `json:"combatPlayerScore"`
	Deaths                          int64 `json:"deaths"`
	DoubleKills                     int64 `json:"doubleKills"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	GoldEarned                      int64 `json:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent"`
	InhibitorKills                  int64 `json:"inhibitorKills"`
	Item0                           int64 `json:"item0"`
	Item1                           int64 `json:"item1"`
	Item2                           int64 `json:"item2"`
	Item3                           int64 `json:"item3"`
	Item4                           int64 `json:"item4"`
	Item5                           int64 `json:"item5"`
	Item6                           int64 `json:"item6"`
	KillingSprees                   int64 `json:"killingSprees"`
	Kills                           int64 `json:"kills"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike"`
	LargestKillingSpree             int64 `json:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int64 `json:"magicDamageTaken"`
	MinionsKilled                   int64 `json:"minionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     int64 `json:"nodeCapture"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist"`
	NodeNeutralize                  int64 `json:"nodeNeutralize"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore"`
	PentaKills                      int64 `json:"pentaKills"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
	QuadraKills                     int64 `json:"quadraKills"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame"`
	TeamObjective                   int64 `json:"teamObjective"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalPlayerScore                int64 `json:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed"`
	TowerKills                      int64 `json:"towerKills"`
	TripleKills                     int64 `json:"tripleKills"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	UnrealKills                     int64 `json:"unrealKills"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame"`
	WardsKilled                     int64 `json:"wardsKilled"`
	WardsPlaced                     int64 `json:"wardsPlaced"`
	Winner                          bool  `json:"winner"`
}

// ParticipantTimeline is the container for a snapshot of participant information at a given time
type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        ParticipantTimelineData `json:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          ParticipantTimelineData `json:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              ParticipantTimelineData `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              ParticipantTimelineData `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         ParticipantTimelineData `json:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       ParticipantTimelineData `json:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         ParticipantTimelineData `json:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    ParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                ParticipantTimelineData `json:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      ParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`
	Lane                            string                  `json:"lane"`
	Role                            string                  `json:"role"`
	TowerAssistsPerMinCounts        ParticipantTimelineData `json:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          ParticipantTimelineData `json:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          ParticipantTimelineData `json:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      ParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        ParticipantTimelineData `json:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               ParticipantTimelineData `json:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              ParticipantTimelineData `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  ParticipantTimelineData `json:"xpPerMinDeltas"`
}

// ParticipantTimelineData is the container to display progression of stats
type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty"`
	ThirtyToTen    float64 `json:"thirtyToTen"`
	TwentyToThirty float64 `json:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen"`
}

// Team is the container with basic team information
type Team struct {
	Bans                 []BannedChampion `json:"bans"`
	BaronKills           int              `json:"baronKills"`
	DominionVictoryScore int64            `json:"dominionVictoryScore"`
	DragonKills          int              `json:"dragonKills"`
	FirstBaron           bool             `json:"firstBaron"`
	FirstBlood           bool             `json:"firstBlood"`
	FirstDragon          bool             `json:"firstDragon"`
	FirstInhibitor       bool             `json:"firstInhibitor"`
	FirstRiftHerald      bool             `json:"firstRiftHerald"`
	FirstTower           bool             `json:"firstTower"`
	InhibitorKills       int              `json:"inhibitorKills"`
	RiftHeraldKills      int              `json:"riftHeraldKills"`
	TeamID               int              `json:"teamId"`
	TowerKills           int              `json:"TowerKills"`
	VilemawKills         int              `json:"vilemawKills"`
	Winner               bool             `json:"winner"`
}

// BannedChampion is the container with banning phase information
type BannedChampion struct {
	ChampionID int64 `json:"championId"`
	PickTurn   int   `json:"pickTurn"`
	TeamID     int64 `json:"teamId"`
}

// Timeline is the container to show events in a match
type Timeline struct {
	FrameInterval int64   `json:"frameInterval"`
	Frames        []Frame `json:"frames"`
}

// Frame is the container for a snapshot in time
type Frame struct {
	Events            []Event                  `json:"events"`
	ParticipantFrames map[int]ParticipantFrame `json:"participantFrames"`
}

// Event is the information for a given moment
type Event struct {
	AscendedType            string   `json:"ascendedType"`
	AssistingParticipantIDs []int    `json:"assistingParticipantIds"`
	BuildingType            string   `json:"buildingType"`
	CreatorID               int      `json:"creatorId"`
	EventType               string   `json:"eventType"`
	ItemAfter               int      `json:"itemAfter"`
	ItemBefore              int      `json:"itemBefore"`
	ItemID                  int      `json:"itemId"`
	KillerID                int      `json:"killerId"`
	LaneType                string   `json:"laneType"`
	LevelUpType             string   `json:"levelUpType"`
	MonsterType             string   `json:"monsterType"`
	ParticipantID           int      `json:"participantId"`
	PointCaptured           string   `json:"pointCaptured"`
	Position                Position `json:"position"`
	SkillSlot               int      `json:"skillSlot"`
	TeamID                  int      `json:"teamId"`
	Timestamp               int64    `json:"timestamp"`
	TowerType               string   `json:"towerType"`
	VictimID                int      `json:"victimId"`
	WardType                string   `json:"wardType"`
}

// Position is the X and Y coords on the map
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ParticipantFrame is the information for a participant in a frame
type ParticipantFrame struct {
	CurrentGold         int      `json:"currentGold"`
	DominionScore       int      `json:"dominionScore"`
	JungleMinionsKilled int      `json:"jungleMinionsKilled"`
	Level               int      `json:"level"`
	MinionsKilled       int      `json:"minionsKilled"`
	ParticipantID       int      `json:"participantId"`
	Position            Position `json:"position"`
	TeamScore           int      `json:"teamScore"`
	TotalGold           int      `json:"totalGold"`
	Xp                  int      `json:"xp"`
}

// GetMatchParams are the optional query params for Get
type GetMatchParams struct {
	IncludeTimeline bool `url:"includeTimeline,omitempty"`
}

const matchPathPart = "api/lol/%s/v2.2/match"

// Get gets the match for a given matchID
func (s *MatchService) Get(ctx context.Context, matchID int64, params *GetMatchParams) (*MatchDetail, error) {
	match := new(MatchDetail)
	err := s.client.getResource(
		ctx,
		addRegionToString(matchPathPart, s.client.region),
		strconv.FormatInt(matchID, 10),
		params,
		match,
	)
	return match, err
}
