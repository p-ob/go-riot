/*
Package lol is a Go client that wraps the Riot Games public API.

All Riot defined objects have been mapped to their Go equivalent, with equivalent property types (e.g. long -> int64).
The only exceptions are properties that were stored as strings, but are safe to parse to int64; LeagueDto.ParticipantID
is an example of this.

Usage

Instantiate a lol.Client using the provided lol.NewClient function.

	// Riot's API requires a valid API key to utilize these methods
	// grab one from https://developer.riotgames.com/.
	apiKey := "your-key-here"

	// the region to query against; to hit multiple regions, create one client per region
	// as the API rate limits you per region per key, not just per key
	// the lol.Client object enforces this, as its base url and region cannot be changed after instantiation
	region := lol.Na

	client := lol.NewClient(apiKey, region, nil)

	// start making API calls (e.g. get summoner with SummonerID = 25886496)
	client.Summoner.Get(ctx, 25886496)

Structure

All 12 endpoints have their own associated *.go file and service on the Client. For example, the summoner endpoint and
all associated objects are stored in the summoner.go file. The API methods exposed are all exposed on the
Client.Summoner service.

Optional Params

For any method that allows query params, there is an associated Params object. For valid values to use with these query
params and further documentation, please see the API docs at https://developer.riotgames.com/api/methods.
*/
package lol
