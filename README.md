# lolgo

lolgo is a [League of Legends API](https://developer.riotgames.com/) client for Go. 

Example usage:  
```golang  
// Initiate a lolgo.Client with your API key, and the region to query against     
client := lolgo.NewClient(apiKey, lolgo.Na, httpClient)  
ctx := context.Background()  

// Drunk7Irishman's summoner id
summonerId := int64(25886496)  

// get Drunk7Irishman
s, _ := client.Summoner.Get(ctx, summonerId)  

// print Drunk7Irishman
thisSummoner := (*s)[summonerId]  
fmt.Printf("Summoner:\n%+v\n", thisSummoner)

// OUTPUT:
// >> Summoner:
//    {Id:25886496 Name:Drunk7Irishman ProfileIconId:744 SummonerLevel:30 RevisionDate:1482381110000}
```

Resources mapped:  
- [Champion v1.2](https://developer.riotgames.com/api/methods#!/1206)
- [ChampionMastery v1.0](https://developer.riotgames.com/api/methods#!/1091)
- [CurrentGame v1.0](https://developer.riotgames.com/api/methods#!/976)
- [FeaturedGames v1.0](https://developer.riotgames.com/api/methods#!/977)
- [Game v1.3](https://developer.riotgames.com/api/methods#!/1207)
- [League v2.5](https://developer.riotgames.com/api/methods#!/1215)
- [Match v2.2](https://developer.riotgames.com/api/methods#!/1224)
- [Matchlist v2.2](https://developer.riotgames.com/api/methods#!/1223)
- [StaticData v1.2](https://developer.riotgames.com/api/methods#!/1055)
- [Stats v1.3](https://developer.riotgames.com/api/methods#!/1209)
- [Status v1.0](https://developer.riotgames.com/api/methods#!/1085)
- [Summoner v1.4](https://developer.riotgames.com/api/methods#!/1208)

