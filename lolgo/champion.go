package main

type ChampionListDto struct {
	Champions []ChampionDto
}

type ChampionDto struct {
	Active            bool
	BotEnabled        bool
	BotMmEnabled      bool
	FreeToPlay        bool
	Id                int64
	RankedPlayEnabled bool
}
