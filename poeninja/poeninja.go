package poeninja

import (
	"encoding/json"

	"github.com/plasmakatt/poecurrencycalc/httpclient"
	"github.com/plasmakatt/poecurrencycalc/utils/constants"
)

type PoeNinja struct {
	Items []PoeNinjaItem `json:"lines"`
}

type PoeNinjaItem struct {
	Id                     int               `json:"id"`
	Name                   string            `json:"name"`
	Icon                   string            `json:"icon"`
	MapTier                int               `json:"mapTier"`
	LevelRequired          int               `json:"levelRequired"`
	BaseType               string            `json:"baseType"`
	StackSize              int               `json:"stackSize"`
	Variant                string            `json:"variant"`
	ProphecyText           string            `json:"prophecyText"`
	ArtFilename            string            `json:"artFilename"`
	Links                  int               `json:"links"`
	ItemClass              int               `json:"itemClass"`
	FlavourText            string            `json:"flavourText"`
	Corrupted              bool              `json:"corrupted"`
	GemLevel               int               `json:"gemLevel"`
	GemQuality             int               `json:"gemQuality"`
	ItemType               string            `json:"itemType"`
	ChaosValue             float64           `json:"chaosValue"`
	ExaltedValue           float64           `json:"exaltedValue"`
	Count                  int               `json:"count"`
	DetailsId              string            `json:"detailsId"`
	Sparkline              sparkline         `json:"sparkline"`
	LowConfidenceSparkline sparkline         `json:"lowConfidenceSparkline"`
	ImplModifiers          []string          `json:"implicitModifiers"`
	ExplModifiers          explicitModifiers `json:"explicitModifiers"`
}

type sparkline struct {
	data        []float32
	totalChange float32
}

type explicitModifiers struct {
	text     string
	optional bool
}

func (poeNinja *PoeNinja) LoadData() {
	var httpClients [12]httpclient.HttpClient

	httpClients[0] = httpclient.HttpClient{TargetURL: constants.DivinationCardsURL}
	httpClients[1] = httpclient.HttpClient{TargetURL: constants.EssenceURL}
	httpClients[2] = httpclient.HttpClient{TargetURL: constants.FragmentURL}
	httpClients[3] = httpclient.HttpClient{TargetURL: constants.CurrencyURL}
	httpClients[4] = httpclient.HttpClient{TargetURL: constants.ProphecyURL}
	httpClients[5] = httpclient.HttpClient{TargetURL: constants.MapURL}
	httpClients[6] = httpclient.HttpClient{TargetURL: constants.UniqueMapURL}
	httpClients[7] = httpclient.HttpClient{TargetURL: constants.UniqueArmorURL}
	httpClients[8] = httpclient.HttpClient{TargetURL: constants.UniqueFlaskURL}
	httpClients[9] = httpclient.HttpClient{TargetURL: constants.UniqueWeaponURL}
	httpClients[10] = httpclient.HttpClient{TargetURL: constants.UniqueAccessoryURL}
	httpClients[11] = httpclient.HttpClient{TargetURL: constants.UniqueJewelURL}

	var tempPoeNinja PoeNinja
	for _, client := range httpClients {
		json.Unmarshal([]byte(client.DoGetRequest()), &tempPoeNinja)
		poeNinja.Items = append(poeNinja.Items, tempPoeNinja.Items...)
	}
}
