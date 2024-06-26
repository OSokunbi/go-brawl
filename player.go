package brawlstars

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
type Player struct {
	Tag             	 string 		`json:"tag"`
	Name            	 string 		`json:"name"`
  	NameColor      	 	 string 		`json:"nameColor"`
   	Icon            	 PlayerIcon 	`json:"icon"`
   	Trophies             int 			`json:"trophies"`
    HighestTrophies      int 			`json:"highestTrophies"`
    ExpLevel             int 			`json:"expLevel"`
    ExpPoints            int 			`json:"expPoints"`
    TeamVictories        int 			`json:"3vs3Victories"`
    SoloVictories        int 			`json:"soloVictories"`
   	DuoVictories         int 			`json:"duoVictories"`
	BestRumbleTime       int 			`json:"bestRoboRumbleTime"`
	BestTimeAsBigBrawler int 			`json:"bestTimeAsBigBrawler"`
    Club            	 PlayerClub 	`json:"club"`
    Brawlers        	 []OwnedBrawler `json:"brawlers"`
}

type PlayerResponse struct {
	Player Player `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type OwnedBrawler struct {
  	ID       		int 	  	`json:"id"`
   	Name     		string 	    `json:"name"`
   	Power           int 		`json:"power"`
    Rank            int 		`json:"rank"`
    Trophies        int 		`json:"trophies"`
    HighestTrophies int 		`json:"highestTrophies"`
    Gears           []Gear 		`json:"gears"`
    StarPowers      []StarPower `json:"starPowers"`
    Gadgets         []Gadget 	`json:"gadgets"`
}

type PlayerClub struct {
  	Tag  string `json:"tag"`
   	Name string `json:"name"`
}


type Gear struct {
  	ID    int 	 `json:"id"`
   	Name  string `json:"name"`
    Level int 	 `json:"level"`
}
type PlayerIcon struct {
	ID int `json:"id"`
}


func (c *Client) GetPlayer(tag string) (*Player, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://bsproxy.royaleapi.dev/v1/players/%s", url.PathEscape(tag)), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + c.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var player Player
	err = json.NewDecoder(resp.Body).Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}


func (c *Client) GetPlayerBattlelog(tag string) ([]Battle, error) {
	encodedTag := url.PathEscape(tag)
	url := fmt.Sprintf("https://bsproxy.royaleapi.dev/v1/players/%s/battlelog", encodedTag)
	
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

func (p *Player) GetBrawler(id int) *OwnedBrawler {
	for _, brawler := range p.Brawlers {
		if brawler.ID == id {
			return &brawler
		}
	}
	return nil
}

func (p *Player) GetBrawlerByName(name string) *OwnedBrawler {
	for _, brawler := range p.Brawlers {
		if strings.ToLower(name) == strings.ToLower(brawler.Name) {
			return &brawler
		}
	}
	return nil
}
