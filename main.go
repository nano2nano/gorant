package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type ContentItemDto struct {
	Name           string            `json:"name"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
	AssetPath      string            `json:"assetPath,omitempty"`
	LocalizedNames LocalizedNamesDto `json:"localizedNames,omitempty"`
}

type ActDto struct {
	ID             string            `json:"id"`
	ParentID       string            `json:"parentId"`
	Type           string            `json:"type"`
	Name           string            `json:"name"`
	IsActive       bool              `json:"isActive"`
	LocalizedNames LocalizedNamesDto `json:"localizedNames,omitempty"`
}

type ContentDto struct {
	Version      string           `json:"version"`
	Characters   []ContentItemDto `json:"characters"`
	Maps         []ContentItemDto `json:"maps"`
	Chromas      []ContentItemDto `json:"chromas"`
	Skins        []ContentItemDto `json:"skins"`
	SkinLevels   []ContentItemDto `json:"skinLevels"`
	Equips       []ContentItemDto `json:"equips"`
	GameModes    []ContentItemDto `json:"gameModes"`
	Sprays       []ContentItemDto `json:"sprays"`
	SprayLevels  []ContentItemDto `json:"sprayLevels"`
	Charms       []ContentItemDto `json:"charms"`
	CharmLevels  []ContentItemDto `json:"charmLevels"`
	PlayerCards  []ContentItemDto `json:"playerCards"`
	PlayerTitles []ContentItemDto `json:"playerTitles"`
	Acts         []ActDto         `json:"acts"`
}

type LocalizedNamesDto struct {
	ArAE string `json:"ar-AE"`
	DeDE string `json:"de-DE"`
	EnUS string `json:"en-US"`
	EsES string `json:"es-ES"`
	EsMX string `json:"es-MX"`
	FrFR string `json:"fr-FR"`
	IDID string `json:"id-ID"`
	ItIT string `json:"it-IT"`
	JaJP string `json:"ja-JP"`
	KoKR string `json:"ko-KR"`
	PlPL string `json:"pl-PL"`
	PtBR string `json:"pt-BR"`
	RuRU string `json:"ru-RU"`
	ThTH string `json:"th-TH"`
	TrTR string `json:"tr-TR"`
	ViVN string `json:"vi-VN"`
	ZhCN string `json:"zh-CN"`
	ZhTW string `json:"zh-TW"`
}

func getContent(apiKey string, content *ContentDto) {
	url := "https://na.api.riotgames.com/val/content/v1/contents?locale=ja-JP"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Riot-Token", apiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &content); err != nil {
		panic(err)
	}
}

func main() {
	apiKey := os.Getenv("RIOT_API_KEY")
	content := new(ContentDto)
	getContent(apiKey, content)
}
