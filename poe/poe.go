package poe

import (
	"encoding/json"

	"github.com/plasmakatt/poecurrencycalc/httpclient"
	"github.com/plasmakatt/poecurrencycalc/utils/constants"
)

type PoeLeague struct {
	Id         string `json:"id"`
	Realm      string `json:"realm"`
	Url        string `json:"url"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
	DelveEvent bool   `json:"delveEvent"`
}

func GetLeagues() []string {
	httpClient := httpclient.HttpClient{TargetURL: constants.LeagueURL}
	var poeLeagues []PoeLeague
	json.Unmarshal([]byte(httpClient.DoGetRequest()), &poeLeagues)
	leagues := make([]string, 8)
	for i, poeLeague := range poeLeagues {
		leagues[i] = poeLeague.Id
	}
	return leagues
}
