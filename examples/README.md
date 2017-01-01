#Example API calls

Requires an API key; get one by registering with the [Riot developer website](https://developer.riotgames.com/). Create 
a key.txt file in the examples folder and paste your API key there.

To run the demo:
```bash
cd examples
go run example.go
```

The example script will prompt for a summoner name to pull data. The region is hard coded to Na; to query different 
regions modify the `region` const defined at the top of `examples\example.go`.