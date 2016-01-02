package main

// ChampionListDto types
type ChampionListDto struct {
	Data    map[string]ChampionDto
	Format  string
	Keys    map[string]string
	Type    string
	Version string
}

type ChampionDto struct {
	AllyTips    []string
	Blurb       string
	EnemyTips   []string
	Id          int
	Image       ImageDto
	Info        InfoDto
	Key         string
	Lore        string
	Name        string
	Partype     string
	Passive     PassiveDto
	Recommended []RecommendedDto
	Skins       []SkinDto
	Spells      []ChampionSpellDto
	Stats       StatsDto
	Tags        []string
	Title       string
}

type ChampionSpellDto struct {
	AltImages            []ImageDto
	Cooldown             []float64
	CooldownBurn         string
	Cost                 []int
	CostBurn             string
	CostType             string
	Description          string
	Effect               [][]float64
	EffectBurn           []string
	Image                ImageDto
	Key                  string
	Leveltip             LevelTipDto
	Maxrank              int
	Name                 string
	Range                interface{}
	RangeBurn            string
	Resource             string
	SanitizedDescription string
	SanitizedTooltip     string
	Tooltip              string
	Vars                 []SpellVarsDto
}

type ImageDto struct {
	Full   string
	Group  string
	H      int
	Sprite string
	W      int
	X      int
	Y      int
}

type InfoDto struct {
	Attack     int
	Defense    int
	Difficulty int
	Magic      int
}

type PassiveDto struct {
	Description          string
	Image                ImageDto
	Name                 string
	SanitizedDescription string
}

type RecommendedDto struct {
	Blocks   []BlockDto
	Champion string
	Map      string
	Mode     string
	Priority bool
	Title    string
	Type     string
}

type SkinDto struct {
	Id   int
	Name string
	Num  int
}

type StatsDto struct {
	Armor                float64
	Armorperlevel        float64
	Attackdamage         float64
	Attackdamageperlevel float64
	Attackrange          float64
	Attackspeedoffset    float64
	Attackspeedperlevel  float64
	Crit                 float64
	Critperlevel         float64
	Hp                   float64
	Hpperlevel           float64
	Hpregen              float64
	Hpregenperlevel      float64
	Movespeed            float64
	Mp                   float64
	Mpperlevel           float64
	Mpregen              float64
	Mpregenperlevel      float64
	Spellblock           float64
	Spellblockperlevel   float64
}

type LevelTipDto struct {
	Effect []string
	Label  []string
}

type SpellVarsDto struct {
	Coeff     []float64
	Dyn       string
	Key       string
	Link      string
	RanksWith string
}

type BlockDto struct {
	Items   []BlockItemDto
	RecMath bool
	Type    string
}

type BlockItemDto struct {
	Count int
	Id    int
}

// ...
