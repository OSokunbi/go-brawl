package brawlstars

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BattleLogResponse struct {
	BattleLogs []Battle `json:"items"`
	Paging     struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type Battle struct {
	Time  string     `json:"battleTime"`
	Event Event      `json:"event,omitempty"`
	Info  BattleInfo `json:"battle"`
}

type Event struct {
	ID   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type BattleInfo struct {
	Mode       string        `json:"mode"`
	Type       string        `json:"type"`
	Result     string        `json:"result"`
	Duration   int           `json:"duration"`
	StarPlayer *BattlePlayer `json:"starPlayer,omitempty"`
	Teams      []Team        `json:"teams,omitempty"`
	Players    []BattlePlayer `json:"players,omitempty"` // Added for solo modes
}

type Team struct {
	Players []BattlePlayer `json:"team"`
}

type BattlePlayer struct {
	Tag     string        `json:"tag"`
	Name    string        `json:"name"`
	Brawler BattleBrawler `json:"brawler"`
}

type BattleBrawler struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}

func (c *Client) GetPlayerBattlelog(tag string) ([]Battle, error) {
	encodedTag := url.PathEscape(tag)
	url := fmt.Sprintf("https://api.brawlstars.com/v1/players/%s/battlelog", encodedTag)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var battlelogResp BattleLogResponse
	err = json.NewDecoder(resp.Body).Decode(&battlelogResp)
	if err != nil {
		return nil, err
	}
	
	return battlelogResp.BattleLogs, nil
}