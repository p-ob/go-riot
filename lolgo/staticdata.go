package main

// ChampionListDto types
type StaticChampionListDto struct {
	Data    map[string]StaticChampionDto
	Format  string
	Keys    map[string]string
	Type    string
	Version string
}

type StaticChampionDto struct {
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

// ItemListDto types
type ItemListDto struct {
	Basic   BasicDataDto
	Data    map[string]ItemDto
	Groups  []GroupDto
	Tree    []ItemTreeDto
	Type    string
	Version string
}

type BasicDataDto struct {
	Colloq               string
	ConsumeOnFull        bool
	Consumed             bool
	Depth                int
	Description          string
	From                 []string
	Gold                 GoldDto
	Group                string
	HideFromAll          bool
	Id                   int
	Image                ImageDto
	InStore              bool
	Into                 []string
	Maps                 map[string]bool
	Name                 string
	Plaintext            string
	RequiredChampion     string
	Rune                 MetaDataDto
	SanitizedDescription string
	SpecialRecipe        int
	Stacks               int
	Stats                BasicDataStatsDto
	Tags                 []string
}

type GroupDto struct {
	MaxGroupOwnable string
	Key             string
}

type ItemDto struct {
	Coolq                string
	ConsumeOnFull        bool
	Consumed             bool
	Depth                int
	Description          string
	Effect               map[string]string
	From                 []string
	Gold                 GoldDto
	Group                string
	HideFromAll          bool
	Id                   int
	Image                ImageDto
	InStore              bool
	Into                 []string
	Maps                 map[string]bool
	Name                 string
	Plaintext            string
	RequiredChampion     string
	Rune                 MetaDataDto
	SanitizedDescription string
	SepcialRecipe        int
	Stacks               int
	Stats                BasicDataStatsDto
	Tags                 []string
}

type ItemTreeDto struct {
	Header string
	Tags   []string
}

type BasicDataStatsDto struct {
	FlatArmorMod                        float64
	FlatAttackSpeedMod                  float64
	FlatBlockMod                        float64
	FlatCritChanceMod                   float64
	FlatCritDamageMod                   float64
	FlatEXPBonus                        float64
	FlatEnergyPoolMod                   float64
	FlatEnergyRegenMod                  float64
	FlatHPPoolMod                       float64
	FlatHPRegenMod                      float64
	FlatMPPoolMod                       float64
	FlatMPRegenMod                      float64
	FlatMagicDamageMod                  float64
	FlatMovementSpeedMod                float64
	FlatPhysicalDamageMod               float64
	FlatSpellBlockMod                   float64
	PercentArmorMod                     float64
	PercentAttackSpeedMod               float64
	PercentBlockMod                     float64
	PercentCritChanceMod                float64
	PercentCritDamageMod                float64
	PercentDodgeMod                     float64
	PercentEXPBonus                     float64
	PercentHPPoolMod                    float64
	PercentHPRegenMod                   float64
	PercentLifeStealMod                 float64
	PercentMPPoolMod                    float64
	PercentMPRegenMod                   float64
	PercentMagicDamageMod               float64
	PercentMovementSpeedMod             float64
	PercentPhysicalDamageMod            float64
	PercentSpellBlockMod                float64
	PercentSpellVampMod                 float64
	RFlatArmorModPerLevel               float64
	RFlatArmorPenetrationMod            float64
	RFlatArmorPenetrationModPerLevel    float64
	RFlatCritChanceModPerLevel          float64
	RFlatCritDamageModPerLevel          float64
	RFlatDodgeMod                       float64
	RFlatDodgeModPerLevel               float64
	RFlatEnergyModPerLevel              float64
	RFlatEnergyRegenModPerLevel         float64
	RFlatGoldPer10Mod                   float64
	RFlatHPModPerLevel                  float64
	RFlatHPRegenModPerLevel             float64
	RFlatMPModPerLevel                  float64
	RFlatMPRegenModPerLevel             float64
	RFlatMagicDamageModPerLevel         float64
	RFlatMagicPenetrationMod            float64
	RFlatMagicPenetrationModPerLevel    float64
	RFlatMovementSpeedModPerLevel       float64
	RFlatPhysicalDamageModPerLevel      float64
	RFlatSpellBlockModPerLevel          float64
	RFlatTimeDeadMod                    float64
	RFlatTimeDeadModPerLevel            float64
	RPercentArmorPenetrationMod         float64
	RPercentArmorPenetrationModPerLevel float64
	RPercentAttackSpeedModPerLevel      float64
	RPercentCooldownMod                 float64
	RPercentCooldownModPerLevel         float64
	RPercentMagicPenetrationMod         float64
	RPercentMagicPenetrationModPerLevel float64
	RPercentMovementSpeedModPerLevel    float64
	RPercentTimeDeadMod                 float64
	RPercentTimeDeadModPerLevel         float64
}

type GoldDto struct {
	Base        int
	Purchasable bool
	Sell        int
	Total       int
}

type MetaDataDto struct {
	IsRune bool
	Tier   string
	Type   string
}

// MasteryListDto types
type MasteryListDto struct {
	Data    map[string]MasteryDto
	Tree    MasteryTreeDto
	Type    string
	Version string
}

type MasteryDto struct {
	Description          []string
	Id                   int
	Image                ImageDto
	MasteryTree          string
	Name                 string
	Prereq               string
	Ranks                int
	SanitizedDescription []string
}

type MasteryTreeDto struct {
	Cunning  []MasteryTreeListDto
	Ferocity []MasteryTreeListDto
	Resolve  []MasteryTreeListDto
}

type MasteryTreeListDto struct {
	MasteryTreeItems []MasteryTreeItemDto
}

type MasteryTreeItemDto struct {
	MasteryId int
	Prereq    string
}

// RuneListDto types
type RuneListDto struct {
	Basic   BasicDataDto
	Data    map[string]RuneDto
	Type    string
	Version string
}

type RuneDto struct {
	Colloq               string
	ConsumeOnFull        bool
	Consumed             bool
	Depth                int
	Description          string
	From                 []string
	Group                string
	HideFromAll          bool
	Id                   int
	Image                ImageDto
	InStore              bool
	Into                 []string
	Maps                 map[string]bool
	Name                 string
	Plaintext            string
	RequiredChampion     string
	Rune                 MetaDataDto
	SanitizedDescription string
	SpecialRecipe        int
	Stacks               int
	Stats                BasicDataStatsDto
	Tags                 []string
}

// SummonerSpellListDto types
type SummonerSpellListDto struct {
	Data    map[string]SummonerSpellDto
	Type    string
	Version string
}

type SummonerSpellDto struct {
	Cooldown             float64
	CooldownBurn         string
	Cost                 []int
	CostBurn             string
	CostType             string
	Description          string
	Effect               [][]float64
	EffectBurn           []string
	Id                   int
	Image                ImageDto
	Key                  string
	Leveltip             LevelTipDto
	Maxrank              int
	Modes                []string
	Name                 string
	Range                interface{}
	RangeBurn            string
	Resource             string
	SanitizedDescription string
	SanitizedTooltip     string
	SummonerLevel        int
	Tooltip              string
	Vars                 []SpellVarsDto
}
