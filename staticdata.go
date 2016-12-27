package lolgo

import (
	"context"
	"strconv"
)

type StaticDataService struct {
	client *Client
}

type StaticDataChampionListDto struct {
	Data    map[string]StaticDataChampionDto `json:"data"`
	Format  string                           `json:"format"`
	Keys    map[string]string                `json:"keys"`
	Type    string                           `json:"type"`
	Version string                           `json:"version"`
}

type StaticDataChampionDto struct {
	AllyTips    []string           `json:"allytips"`
	Blurb       string             `json:"blurb"`
	EnemyTips   []string           `json:"enemytips"`
	Id          int                `json:"id"`
	Image       ImageDto           `json:"image"`
	Info        InfoDto            `json:"info"`
	Key         string             `json:"key"`
	Lore        string             `json:"lore"`
	Name        string             `json:"name"`
	ParType     string             `json:"partype"`
	Passive     PassiveDto         `json:"passive"`
	Recommended []RecommendedDto   `json:"recommended"`
	Skins       []SkinDto          `json:"skins"`
	Spells      []ChampionSpellDto `json:"spells"`
	Stats       StatsDto           `json:"stats"`
	Tags        []string           `json:"tags"`
	Title       string             `json:"title"`
}

type ChampionSpellDto struct {
	AltImages            []ImageDto     `json:"altimages"`
	Cooldown             []float64      `json:"cooldown"`
	CooldownBurn         string         `json:"cooldownBurn"`
	Cost                 []int          `json:"cost"`
	CostBurn             string         `json:"costBurn"`
	CostType             string         `json:"costType"`
	Description          string         `json:"description"`
	Effect               [][]float64    `json:"effect"`
	EffectBurn           []string       `json:"effectBurn"`
	Image                ImageDto       `json:"image"`
	Key                  string         `json:"key"`
	LevelTip             LevelTipDto    `json:"leveltip"`
	Maxrank              int            `json:"maxrank"`
	Name                 string         `json:"name"`
	Range                interface{}    `json:"range"`
	RangeBurn            string         `json:"rangeBurn"`
	Resource             string         `json:"resource"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	Tooltip              string         `json:"tooltip"`
	Vars                 []SpellVarsDto `json:"vars"`
}

type ImageDto struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	H      int    `json:"h"`
	Sprint string `json:"sprint"`
	W      int    `json:"w"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type InfoDto struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Difficulty int `json:"difficulty"`
	Magic      int `json:"magic"`
}

type PassiveDto struct {
	Description          string   `json:"description"`
	Image                ImageDto `json:"image"`
	Name                 string   `json:"name"`
	SanitizedDescription string   `json:"sanitizedDescription"`
}

type RecommendedDto struct {
	Blocks   []BlockDto `json:"blocks"`
	Champion string     `json:"champion"`
	Map      string     `json:"map"`
	Mode     string     `json:"mode"`
	Priority bool       `json:"priority"`
	Title    string     `json:"title"`
	Type     string     `json:"type"`
}

type SkinDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type StatsDto struct {
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackRange          float64 `json:"attackrange"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	Crit                 float64 `json:"crit"`
	CritPerLevel         float64 `json:"critperlevel"`
	Hp                   float64 `json:"hp"`
	HpPerLevel           float64 `json:"hpperlevel"`
	HpRegen              float64 `json:"hpregen"`
	HpRegenPerLevel      float64 `json:"hpregenperlevel"`
	Movespeed            float64 `json:"movespeed"`
	Mp                   float64 `json:"mp"`
	MpPerLevel           float64 `json:"mpperlevel"`
	MpRegen              float64 `json:"mpregen"`
	MpRegenPerLevel      float64 `json:"mpregenperlevel"`
	Spellblock           float64 `json:"spellblock"`
	SpellblockPerLevel   float64 `json:"spellblockperlevel"`
}

type LevelTipDto struct {
	Effect []string `json:"effect"`
	Label  []string `json:"label"`
}

type SpellVarsDto struct {
	Coeff     []float64 `json:"coeff"`
	Dyn       string    `json:"dyn"`
	Key       string    `json:"key"`
	Link      string    `json:"link"`
	RanksWith string    `json:"ranksWith"`
}

type BlockDto struct {
	Items   []BlockItemDto `json:"items"`
	RecMath bool           `json:"recMath"`
	Type    string         `json:"type"`
}

type BlockItemDto struct {
	Count int `json:"count"`
	Id    int `json:"id"`
}

type GetChampionStaticDataParams struct {
	Locale    string `url:"locale,omitempty"`
	Version   string `url:"version,omitempty"`
	DataById  bool   `url:"dataById,omitempty"`
	ChampData string `url:"champData,omitempty"`
}

const staticDataPathPart = "api/lol/static-data/%s/v1.2"

func (s *StaticDataService) GetChampions(ctx context.Context, params *GetChampionStaticDataParams) (*StaticDataChampionListDto, error) {
	champions := new(StaticDataChampionListDto)

	err := s.client.GetResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.Region)+"/champion",
		"",
		params,
		champions)
	return champions, err
}

func (s *StaticDataService) GetChampion(ctx context.Context, championId int, params *GetChampionStaticDataParams) (*StaticDataChampionListDto, error) {
	champions := new(StaticDataChampionListDto)

	err := s.client.GetResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.Region)+"/champion",
		strconv.Itoa(championId),
		params,
		champions)
	return champions, err
}
