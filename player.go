package brawlstars

import (
	"net/http"
	"encoding/json"
)
type Player struct {
	Tag             string
	Name            string
  	NameColor       string
   	Icon            PlayerIcon
    Stats           Stats
    Club            PlayerClub
    Brawlers        []OwnedBrawler
}

type OwnedBrawler struct {
  	ID       string
   	Name     string  
    Stats    BrawlerStats
}

type PlayerClub struct {
  	Tag  string
   	Name string
}

type BrawlerStats struct {
  	Power           int
   	Rank            int
    Trophies        int
    HighestTrophies int
    Gears           []Gear
    StarPowers      []StarPower
    Gadgets         []Gadget
}

type Gear struct {
  	ID    int
   	Name  string
    Level int
}
type PlayerIcon struct {
	ID string
}

type Stats struct {
  	Trophies             int
   	HighestTrophies      int
    ExpLevel             int
    ExpPoints            int
    TeamVictories        int
    SoloVictories        int
  	DuoVictories         int
	BestRumbleTime       int
	BestTimeAsBigBrawler int
}

func (c *Client) GetPlayer(tag string) (*Player, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/players/%23" + tag, nil)
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

func (c *Client) GetPlayerBattlelog(tag string) (*BattleLog, error){
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/players/%23" + tag + "/battlelog", nil)
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
	var battlelog BattleLog
	err = json.NewDecoder(resp.Body).Decode(&battlelog)
	if err != nil {
		return nil, err
	}
	return &battlelog, nil
}
