package main

type TeamDto struct {
	CreateDate                    int64
	FullId                        string
	LastGameDate                  int64
	LastJoinDate                  int64
	LastJoinedRankedTeamQueueDate int64
	MatchHistory                  []MatchHistorySummaryDto
	ModifyDate                    int64
	Name                          string
	Roster                        RosterDto
	SecondLastJoinDate            int64
	Status                        string
	Tag                           string
	TeamStatDetails               []TeamStatDetailDto
	ThirdLastJoinDate             int64
}

type MatchHistorySummaryDto struct {
	Assists           int
	Date              int64
	Deaths            int
	GameId            int64
	GameMode          string
	Invalid           bool
	Kills             int
	MapId             int
	OpposingTeamKills int
	OpposingTeamName  string
	Win               bool
}

type RosterDto struct {
	MemberList []TeamMemberInfoDto
	OwnerId    int64
}

type TeamStatDetailDto struct {
	AverageGamesPlayed int
	Losses             int
	TeamStatType       string
	Wins               int
}

type TeamMemberInfoDto struct {
	InviteDate int64
	JoinDate   int64
	PlayerId   int64
	Status     string
}
