package main

type MatchDetail struct {
	MapId                 int
	MatchCreation         int64
	MatchDuration         int64
	MatchId               int64
	MatchMode             string
	MatchType             string
	MatchVersion          string
	ParticipantIdentities []ParticipantIdentity
	Participants          []Participant
	QueueType             string
	Region                string
	Season                string
	Teams                 []Team
	Timeline              Timeline
}

type ParticipantIdentity struct {
	ParticipantId int
	Player        Player
}

type Player struct {
	MatchHistoryUri string
	ProfileIcon     int
	SummonerId      int64
	SummonerName    string
}

type Participant struct {
	ChampionId                int
	HighestAchievedSeasonTier string
	Masteries                 []Mastery
	ParticipantId             int
	Runes                     []Rune
	Spell1Id                  int
	Spell2Id                  int
	Stats                     ParticipantStats
	TeamId                    int
	Timeline                  ParticipantTimeline
}

type Mastery struct {
	MasteryId int64
	Rank      int64
}

type Rune struct {
	Rank   int64
	RuneId int64
}

type ParticipantStats struct {
	Assists                         int64
	ChampLevel                      int64
	CombatPlayerScore               int64
	Deaths                          int64
	DoubleKills                     int64
	FirstBloodAssist                bool
	FirstBloodKill                  bool
	FirstInhibitorAssist            bool
	FirstInhibitorKill              bool
	FirstTowerAssist                bool
	FirstTowerKill                  bool
	GoldEarned                      int64
	GoldSpent                       int64
	InhibitorKills                  int64
	Item0                           int64
	Item1                           int64
	Item2                           int64
	Item3                           int64
	Item4                           int64
	Item5                           int64
	Item6                           int64
	KillingSprees                   int64
	Kills                           int64
	LargestCriticalStrike           int64
	LargestKillingSpree             int64
	LargestMultiKill                int64
	MagicDamageDealt                int64
	MagicDamageDealtToChampions     int64
	MagicDamageTaken                int64
	MinionsKilled                   int64
	NeutralMinionsKilled            int64
	NeutralMinionsKilledEnemyJungle int64
	NeutralMinionsKilledTeamJungle  int64
	NodeCapture                     int64
	NodeCaptureAssist               int64
	NodeNeutralize                  int64
	ObjectivePlayerScore            int64
	PentaKills                      int64
	PhysicalDamageDealt             int64
	PhysicalDamageDealtToChampions  int64
	PhysicalDamageTaken             int64
	QuadraKills                     int64
	SightWardsBoughtInGame          int64
	TeamObjective                   int64
	TotalDamageDealt                int64
	TotalDamageDealtToChampions     int64
	TotalDamageTaken                int64
	TotalHeal                       int64
	TotalPlayerScore                int64
	TotalScoreRank                  int64
	TotalTimeCrowdControlDealt      int64
	TotalUnitsHealed                int64
	TowerKills                      int64
	TripleKills                     int64
	TrueDamageDealt                 int64
	TrueDamageDealtToChampions      int64
	TrueDamageTaken                 int64
	UnrealKills                     int64
	VisionWardsBoughtInGame         int64
	WardsKilled                     int64
	WardsPlaced                     int64
	Winner                          bool
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData
	AncientGolemKillsPerMinCounts   ParticipantTimelineData
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData
	BaronAssistsPerMinCounts        ParticipantTimelineData
	BaronKillsPerMinCounts          ParticipantTimelineData
	CreepsPerMinDeltas              ParticipantTimelineData
	CsDiffPerMinDeltas              ParticipantTimelineData
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData
	DamageTakenPerMinDeltas         ParticipantTimelineData
	DragonAssistsPerMinCounts       ParticipantTimelineData
	DragonKillsPerMinCounts         ParticipantTimelineData
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData
	ElderLizardKillsPerMinCounts    ParticipantTimelineData
	GoldPerMinDeltas                ParticipantTimelineData
	InhibitorAssistsPerMinCounts    ParticipantTimelineData
	InhibitorKillsPerMinCounts      ParticipantTimelineData
	Lane                            string
	Role                            string
	TowerAssistsPerMinCounts        ParticipantTimelineData
	TowerKillsPerMinCounts          ParticipantTimelineData
	TowerKillsPerMinDeltas          ParticipantTimelineData
	VilemawAssistsPerMinCounts      ParticipantTimelineData
	VilemawKillsPerMinCounts        ParticipantTimelineData
	WardsPerMinDeltas               ParticipantTimelineData
	XpDiffPerMinDeltas              ParticipantTimelineData
	XpPerMinDeltas                  ParticipantTimelineData
}

type ParticipantTimelineData struct {
	TenToTwenty    float64
	ThirtyToTen    float64
	TwentyToThirty float64
	ZeroToTen      float64
}

type Team struct {
	Bans                 []BannedChampion
	BaronKills           int
	DominionVictoryScore int64
	DragonKills          int
	FirstBaron           bool
	FirstBlood           bool
	FirstDragon          bool
	FirstInhibitor       bool
	FirstRiftHerald      bool
	FirstTower           bool
	InhibitorKills       int
	RiftHeraldKills      int
	TeamId               int
	TowerKills           int
	VilemawKills         int
	Winner               bool
}

type BannedChampion struct {
	ChampionId int
	PickTurn   int
}

type Timeline struct {
	FrameInterval int64
	Frames        []Frame
}

type Frame struct {
	Events            []Event
	ParticipantFrames map[string]ParticipantFrame
}

type Event struct {
	AscendedType            string
	AssistingParticipantIds []int
	BuildingType            string
	CreatorId               int
	EventType               string
	ItemAfter               int
	ItemBefore              int
	ItemId                  int
	KillerId                int
	LaneType                string
	LevelUpType             string
	MonsterType             string
	ParticipantId           int
	PointCaptured           string
	Position                Position
	SkillSlot               int
	TeamId                  int
	Timestamp               int64
	TowerType               string
	VictimId                int
	WardType                string
}

type Position struct {
	X int
	Y int
}

type ParticipantFrame struct {
	CurrentGold         int
	DominionScore       int
	JungleMinionsKilled int
	Level               int
	MinionsKilled       int
	ParticipantId       int
	Position            Position
	TeamScore           int
	TotalGold           int
	Xp                  int
}
