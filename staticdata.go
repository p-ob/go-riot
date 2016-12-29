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

type ItemListDto struct {
	Basic   BasicDataDto       `json:"basic"`
	Data    map[string]ItemDto `json:"data"`
	Groups  []GroupDto         `json:"groups"`
	Tree    []ItemTreeDto      `json:"tree"`
	Type    string             `json:"type"`
	Version string             `json:"version"`
}

type BasicDataDto struct {
	Colloq               string            `json:"colloq"`
	ConsumerOnFull       bool              `json:"consumerOnFull"`
	Consumed             bool              `json:"consumed"`
	Depth                int               `json:"depth"`
	Description          string            `json:"description"`
	From                 []string          `json:"from"`
	Gold                 GoldDto           `json:"gold"`
	Group                string            `json:"group"`
	HideFromAll          bool              `json:"hideFromAll"`
	Id                   int               `json:"id"`
	Image                ImageDto          `json:"image"`
	InStore              bool              `json:"inStore"`
	Into                 []string          `json:"into"`
	Maps                 map[string]bool   `json:"maps"`
	Name                 string            `json:"name"`
	PlainText            string            `json:"plaintext"`
	RequiredChampion     string            `json:"requiredChampion"`
	Rune                 MetaDataDto       `json:"rune"`
	SanitizedDescription string            `json:"sanitizedDescription"`
	SpecialRecipe        int               `json:"specialRecipe"`
	Stacks               int               `json:"stacks"`
	Stats                BasicDataStatsDto `json:"stats"`
	Tags                 []string          `json:"tags"`
}

type GroupDto struct {
	MaxGroupOwnable string `json:"MaxGroupOwnable"`
	Key             string `json:"key"`
}

type ItemDto struct {
	Colloq               string            `json:"colloq"`
	ConsumeOnFull        bool              `json:"consumeOnFull"`
	Consumed             bool              `json:"consumed"`
	Depth                int               `json:"depth"`
	Description          string            `json:"description"`
	Effect               map[string]string `json:"effect"`
	From                 []string          `json:"from"`
	Gold                 GoldDto           `json:"gold"`
	Group                string            `json:"group"`
	HideFromAll          bool              `json:"hideFromAll"`
	Id                   int               `json:"id"`
	Image                ImageDto          `json:"image"`
	InStore              bool              `json:"inStore"`
	Into                 []string          `json:"into"`
	Maps                 map[string]bool   `json:"maps"`
	Name                 string            `json:"name"`
	PlainText            string            `json:"plaintext"`
	RequiredChampion     string            `json:"requiredChampion"`
	Rune                 MetaDataDto       `json:"rune"`
	SanitizedDescription string            `json:"sanitizedDescription"`
	SpecialRecipe        int               `json:"specialRecipe"`
	Stacks               int               `json:"stacks"`
	Stats                BasicDataStatsDto `json:"stats"`
}

type ItemTreeDto struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}

type BasicDataStatsDto struct {
	FlatArmorMod                        float64 `json:"FlatArmorMod"`
	FlatAttackSpeedMod                  float64 `json:"FlatAttackSpeedMod"`
	FlatBlockMod                        float64 `json:"FlatBlockMod"`
	FlatCritChanceMod                   float64 `json:"FlatCritChanceMod"`
	FlatCritDamageMod                   float64 `json:"FlatCritDamageMod"`
	FlatExpBonus                        float64 `json:"FlatEXPBonus"`
	FlatEnerygyPoolMod                  float64 `json:"FlatEnerygyPoolMod"`
	FlatEnergyRegenMod                  float64 `json:"FlatEnergyRegenMod"`
	FlatHpPoolMod                       float64 `json:"FlatHPPoolMod"`
	FlatHpRegenMod                      float64 `json:"FlatHPRegenMod"`
	FlatMpPoolMod                       float64 `json:"FlatMPPoolMod"`
	FlatMpRegenMod                      float64 `json:"FlatMPRegenMod"`
	FlatMagicDamageMod                  float64 `json:"FlatMagicDamageMod"`
	FlatMovementSpeedMod                float64 `json:"FlatMovementSpeedMod"`
	FlatPhysicalDamageMod               float64 `json:"FlatPhysicalDamageMod"`
	FlatSpellBlockMod                   float64 `json:"FlatSpellBlockMod"`
	PercentArmorMod                     float64 `json:"PercentArmorMod"`
	PercentAttackSpeedMod               float64 `json:"PercentAttackSpeedMod"`
	PercentBlockMod                     float64 `json:"PercentBlockMod"`
	PercentCritChanceMod                float64 `json:"PercentCritChanceMod"`
	PercentCritDamageMod                float64 `json:"PercentCritDamageMod"`
	PercentDodgeMod                     float64 `json:"PercentDodgeMod"`
	PercentExpBonus                     float64 `json:"PercentEXPBonus"`
	PercentHpPoolMod                    float64 `json:"PercentHPPoolMod"`
	PercentHpRegenMod                   float64 `json:"PercentHPRegenMod"`
	PercentLifeStealMod                 float64 `json:"PercentLifeStealMod"`
	PercentMpPoolMod                    float64 `json:"PercentMPPoolMod"`
	PercentMpRegendMod                  float64 `json:"PercentMPRegendMod"`
	PercentMagicDamageMod               float64 `json:"PercentMagicDamageMod"`
	PercentMovementSpeedMod             float64 `json:"PercentMovementSpeedMod"`
	PercentPhysicalDamageMod            float64 `json:"PercentPhysicalDamageMod"`
	PercentSpellBlockMod                float64 `json:"PercentSpellBlockMod"`
	PercentSpellVampMod                 float64 `json:"PercentSpellVampMod"`
	RFlatArmorModPerLevel               float64 `json:"rFlatArmorModPerLevel"`
	RFlatArmorPenetrationMod            float64 `json:"rFlatArmorPenetrationMod"`
	RFlatArmorPenetrationModPerLevel    float64 `json:"rFlatArmorPenetrationModPerLevel"`
	RFlatCritChanceModPerLevel          float64 `json:"rFlatCritChanceModPerLevel"`
	RFlatCritDamageModPerLevel          float64 `json:"rFlatCritDamageModPerLevel"`
	RFlatDodgeMod                       float64 `json:"rFlatDodgeMod"`
	RFlatDodgeModPerLevel               float64 `json:"rFlatDodgeModPerLevel"`
	RFlatEnergyModPerLevel              float64 `json:"rFlatEnergyModPerLevel"`
	RFlatEnergyRegenModPerLevel         float64 `json:"rFlatEnergyRegenModPerLevel"`
	RFlatGoldPer10Mod                   float64 `json:"rFlatGoldPer10Mod"`
	RFlatHpModPerLevel                  float64 `json:"rFlatHPModPerLevel"`
	RFlatHpRegenModPerLevel             float64 `json:"rFlatHPRegenModPerLevel"`
	RFlatMpModPerLevel                  float64 `json:"rFlatMPModPerLevel"`
	RFlatMpRegenModPerLevel             float64 `json:"rFlatMPRegenModPerLevel"`
	RFlatMagicDamageModPerLevel         float64 `json:"rFlatMagicDamageModPerLevel"`
	RFlatMagicPenetrationMod            float64 `json:"rFlatMagicPenetrationMod"`
	RFlatMagicPenetrationModPerLevel    float64 `json:"rFlatMagicPenetrationModPerLevel"`
	RFlatMovementSpeedModPerLevel       float64 `json:"rFlatMovementSpeedModPerLevel"`
	RFlatPhysicalDamageModPerLevel      float64 `json:"rFlatPhysicalDamageModPerLevel"`
	RFlatSpellBlockModPerLevel          float64 `json:"rFlatSpellBlockModPerLevel"`
	RFlatTimeDeadMod                    float64 `json:"rFlatTimeDeadMod"`
	RFlatTimeDeadModPerLevel            float64 `json:"rFlatTimeDeadModPerLevel"`
	RPercentArmorPenetrationMod         float64 `json:"rPercentArmorPenetrationMod"`
	RPercentArmorPenetrationModPerLevel float64 `json:"rPercentArmorPenetrationModPerLevel"`
	RPercentAttackSpeedModPerLevel      float64 `json:"rPercentAttackSpeedModPerLevel"`
	RPercentCooldownMod                 float64 `json:"rPercentCooldownMod"`
	RPercentCooldownModPerLevel         float64 `json:"rPercentCooldownModPerLevel"`
	RPercentMagicPenetrationMod         float64 `json:"rPercentMagicPenetrationMod"`
	RPercentMagicPenetrationModPerLevel float64 `json:"rPercentMagicPenetrationModPerLevel"`
	RPercentMovementSpeedModPerLevel    float64 `json:"rPercentMovementSpeedModPerLevel"`
	RPercentTimeDeadMod                 float64 `json:"rPercentTimeDeadMod"`
	RPercentTimeDeadModPerLevel         float64 `json:"rPercentTimeDeadModPerLevel"`
}

type GoldDto struct {
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
	Sell        int  `json:"sell"`
	Total       int  `json:"total"`
}

type MetaDataDto struct {
	IsRune bool   `json:"isRune"`
	Tier   string `json:"tier"`
	Type   string `json:"type"`
}

type LanguageStringsDto struct {
	Data    map[string]string `json:"data"`
	Type    string            `json:"type"`
	Version string            `json:"version"`
}

type MasteryListDto struct {
	Data    map[string]StaticDataMasteryDto `json:"data"`
	Tree    MasteryTreeDto                  `json:"tree"`
	Type    string                          `json:"type"`
	Version string                          `json:"version"`
}

type StaticDataMasteryDto struct {
	Description          []string `json:"description"`
	Id                   int      `json:"id"`
	Image                ImageDto `json:"image"`
	MasteryTree          string   `json:"masteryTree"`
	Name                 string   `json:"name"`
	Prereq               string   `json:"prereq"`
	Ranks                int      `json:"ranks"`
	SanitizedDescription []string `json:"sanitizedDescription"`
}

type MasteryTreeDto struct {
	Cunning  []MasteryTreeListDto `json:"Cunning"`
	Ferocity []MasteryTreeListDto `json:"Ferocity"`
	Resolve  []MasteryTreeListDto `json:"Resolve"`
}

type MasteryTreeListDto struct {
	MasteryTreeItems []MasteryTreeItemDto `json:"masteryTreeItems"`
}

type MasteryTreeItemDto struct {
	MasteryId int    `json:"masteryId"`
	Prereq    string `json:"prereq"`
}

type RealmDto struct {
	Cdn            string            `json:"cdn"`
	Css            string            `json:"css"`
	Dd             string            `json:"dd"`
	L              string            `json:"l"`
	Lg             string            `json:"lg"`
	N              map[string]string `json:"n"`
	ProfileIconMax int               `json:"profileiconmax"`
	Store          string            `json:"store"`
	V              string            `json:"v"`
}

type RuneListDto struct {
	Basic   BasicDataDto                 `json:"basic"`
	Data    map[string]StaticDataRuneDto `json:"data"`
	Type    string                       `json:"type"`
	Version string                       `json:"version"`
}

type StaticDataRuneDto struct {
	Colloq               string            `json:"colloq"`
	ConsumeOnFull        bool              `json:"consumeOnFull"`
	Consumed             bool              `json:"consumed"`
	Depth                int               `json:"depth"`
	Description          string            `json:"description"`
	From                 []string          `json:"from"`
	Group                string            `json:"group"`
	HideFromAll          bool              `json:"hideFromAll"`
	Id                   int               `json:"id"`
	Image                ImageDto          `json:"image"`
	InStore              bool              `json:"inStore"`
	Into                 []string          `json:"into"`
	Maps                 map[string]bool   `json:"maps"`
	Name                 string            `json:"name"`
	PlainText            string            `json:"plaintext"`
	RequiredChampion     string            `json:"requiredChampion"`
	Rune                 MetaDataDto       `json:"rune"`
	SanitizedDescription string            `json:"sanitizedDescription"`
	SpecialRecipe        int               `json:"specialRecipe"`
	Stacks               int               `json:"stacks"`
	Stats                BasicDataStatsDto `json:"stats"`
	Tags                 []string          `json:"tags"`
}

type SummonerSpellListDto struct {
	Data    map[string]SummonerSpellDto `json:"data"`
	Type    string                      `json:"type"`
	Version string                      `json:"version"`
}

type SummonerSpellDto struct {
	Cooldown             []float64      `json:"cooldown"`
	CooldownBurn         string         `json:"cooldownBurn"`
	Cost                 []int          `json:"cost"`
	CostBurn             string         `json:"costBurn"`
	CostType             string         `json:"costType"`
	Description          string         `json:"description"`
	Effect               [][]float64    `json:"effect"`
	EffectBurn           []string       `json:"effectBurn"`
	Id                   int            `json:"id"`
	Image                ImageDto       `json:"image"`
	Key                  string         `json:"key"`
	LevelTip             LevelTipDto    `json:"levelTip"`
	Maxrank              int            `json:"maxrank"`
	Modes                []string       `json:"modes"`
	Name                 string         `json:"name"`
	Range                interface{}    `json:"range"`
	RangeBurn            string         `json:"rangeBurn"`
	Resource             string         `json:"resource"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	SummonerLevel        int            `json:"summonerLevel"`
	Tooltip              string         `json:"tooltip"`
	Vars                 []SpellVarsDto `json:"vars"`
}

type GetStaticDataBaseParams struct {
	Locale  string `url:"locale,omitempty"`
	Version string `url:"version,omitempty"`
}

type GetChampionStaticDataParams struct {
	GetStaticDataBaseParams
	DataById  bool   `url:"dataById,omitempty"`
	ChampData string `url:"champData,omitempty"`
}

type GetItemStaticDataParams struct {
	GetStaticDataBaseParams
	ItemListData string `url:"itemListData,omitempty"`
}

type GetMasteryStaticDataParams struct {
	GetStaticDataBaseParams
	MasteryListData string `url:"masteryListData,omitempty"`
}

type GetRuneStaticDataParams struct {
	GetStaticDataBaseParams
	RuneListData string `url:"runeListData,omitempty"`
}

type GetSummonerSpellStaticDataParams struct {
	GetStaticDataBaseParams
	DataById  bool   `url:"dataById,omitempty"`
	SpellData string `url:"spellData,omitempty"`
}

const staticDataPathPart = "api/lol/static-data/%s/v1.2"

func (s *StaticDataService) GetChampions(ctx context.Context, params *GetChampionStaticDataParams) (*StaticDataChampionListDto, error) {
	champions := new(StaticDataChampionListDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/champion",
		"",
		params,
		champions)
	return champions, err
}

func (s *StaticDataService) GetChampion(ctx context.Context, championId int, params *GetChampionStaticDataParams) (*StaticDataChampionDto, error) {
	champion := new(StaticDataChampionDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/champion",
		strconv.Itoa(championId),
		params,
		champion)
	return champion, err
}

func (s *StaticDataService) GetItems(ctx context.Context, params *GetItemStaticDataParams) (*ItemListDto, error) {
	items := new(ItemListDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/item",
		"",
		params,
		items)
	return items, err
}

func (s *StaticDataService) GetItem(ctx context.Context, itemId int, params *GetItemStaticDataParams) (*ItemDto, error) {
	item := new(ItemDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/item",
		strconv.Itoa(itemId),
		params,
		item)
	return item, err
}

func (s *StaticDataService) GetLanguageStrings(ctx context.Context) (*LanguageStringsDto, error) {
	languageStrings := new(LanguageStringsDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/language-strings",
		"",
		nil,
		languageStrings)
	return languageStrings, err
}

func (s *StaticDataService) GetLanguages(ctx context.Context) (*[]string, error) {
	languages := new([]string)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/languages",
		"",
		nil,
		languages)
	return languages, err
}

func (s *StaticDataService) GetMasteries(ctx context.Context, params *GetMasteryStaticDataParams) (*MasteryListDto, error) {
	masteries := new(MasteryListDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/mastery",
		"",
		params,
		masteries)
	return masteries, err
}

func (s *StaticDataService) GetMastery(ctx context.Context, masteryId int, params *GetMasteryStaticDataParams) (*StaticDataMasteryDto, error) {
	mastery := new(StaticDataMasteryDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/mastery",
		strconv.Itoa(masteryId),
		params,
		mastery)
	return mastery, err
}

func (s *StaticDataService) GetRealm(ctx context.Context) (*RealmDto, error) {
	realm := new(RealmDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/realm",
		"",
		nil,
		realm)
	return realm, err
}

func (s *StaticDataService) GetRunes(ctx context.Context, params *GetRuneStaticDataParams) (*RuneListDto, error) {
	runes := new(RuneListDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/rune",
		"",
		params,
		runes)
	return runes, err
}

func (s *StaticDataService) GetRune(ctx context.Context, runeId int, params *GetRuneStaticDataParams) (*StaticDataRuneDto, error) {
	r := new(StaticDataRuneDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/rune",
		strconv.Itoa(runeId),
		params,
		r)
	return r, err
}

func (s *StaticDataService) GetSummonerSpells(ctx context.Context, params *GetSummonerSpellStaticDataParams) (*SummonerSpellListDto, error) {
	spells := new(SummonerSpellListDto)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/summoner-spell",
		"",
		params,
		spells)
	return spells, err
}

func (s *StaticDataService) GetSummonerSpell(ctx context.Context, runeId int, params *GetSummonerSpellStaticDataParams) (*SummonerSpellDto, error) {
	spell := new(SummonerSpellDto)
	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/summoner-spell",
		strconv.Itoa(runeId),
		params,
		spell)
	return spell, err
}

func (s *StaticDataService) GetVersions(ctx context.Context) (*[]string, error) {
	versions := new([]string)

	err := s.client.getResource(
		ctx,
		addRegionToString(staticDataPathPart, s.client.region)+"/versions",
		"",
		nil,
		versions)
	return versions, err
}
