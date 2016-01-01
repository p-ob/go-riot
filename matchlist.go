package main

type MatchList struct {
	StartIndex int
	EndIndex   int
	TotalGames int
	Matches    []MatchReference
}

type MatchReference struct {
	Champion   int64
	Lane       string
	MatchId    int64
	PlatformId string
	Queue      string
	Region     string
	Role       string
	Season     string
	Timestamp  int64
}
