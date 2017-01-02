# go-riot [![Build Status](https://travis-ci.org/p-ob/go-riot.svg?branch=master)](https://travis-ci.org/p-ob/go-riot)

go-riot is a [Riot Games API](https://developer.riotgames.com/) client for Go.

## Install

    go get github.com/p-ob/go-riot/lol
    
## Example usage:  

```golang  
// Initiate a lol.Client with your API key, and the region to query against     
client := lol.NewClient(apiKey, lol.Na, httpClient)  
ctx := context.Background()  

// Drunk7Irishman's summoner id
summonerID := int64(25886496)  

// get Drunk7Irishman
s, _ := client.Summoner.Get(ctx, summonerID)  

// print Drunk7Irishman
thisSummoner := (*s)[summonerID]  
fmt.Printf("Summoner:\n%+v\n", thisSummoner)

// OUTPUT:
// >> Summoner:
//    {ID:25886496 Name:Drunk7Irishman ProfileIconID:744 SummonerLevel:30 RevisionDate:1482381110000}
```

## Resources mapped:  

- [Champion v1.2](https://developer.riotgames.com/api/methods#!/1206)
- [ChampionMastery v1.0](https://developer.riotgames.com/api/methods#!/1091)
- [CurrentGame v1.0](https://developer.riotgames.com/api/methods#!/976)
- [FeaturedGames v1.0](https://developer.riotgames.com/api/methods#!/977)
- [Game v1.3](https://developer.riotgames.com/api/methods#!/1207)
- [League v2.5](https://developer.riotgames.com/api/methods#!/1215)
- [Match v2.2](https://developer.riotgames.com/api/methods#!/1224)
- [MatchList v2.2](https://developer.riotgames.com/api/methods#!/1223)
- [StaticData v1.2](https://developer.riotgames.com/api/methods#!/1055)
- [Stats v1.3](https://developer.riotgames.com/api/methods#!/1209)
- [Status v1.0](https://developer.riotgames.com/api/methods#!/1085)
- [Summoner v1.4](https://developer.riotgames.com/api/methods#!/1208)

