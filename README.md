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
```

Resources mapped:  
- [Champion](https://developer.riotgames.com/api/methods#!/1206)
- [ChampionMastery](https://developer.riotgames.com/api/methods#!/1091)
- [CurrentGame](https://developer.riotgames.com/api/methods#!/976)
- [FeaturedGame](https://developer.riotgames.com/api/methods#!/977)
- [Game](https://developer.riotgames.com/api/methods#!/1207)
- [League](https://developer.riotgames.com/api/methods#!/1215)
- [Match](https://developer.riotgames.com/api/methods#!/1224)
- [Matchlist](https://developer.riotgames.com/api/methods#!/1223)
- [StaticData](https://developer.riotgames.com/api/methods#!/1055)
- [Stats](https://developer.riotgames.com/api/methods#!/1209)
- [Status](https://developer.riotgames.com/api/methods#!/1085)
- [Summoner](https://developer.riotgames.com/api/methods#!/1208)

