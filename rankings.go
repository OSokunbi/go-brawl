package brawlstars

import (
	"net/http"
	"encoding/json"
)

type CountryLeaderboardPlayers struct {
	Players []LeaderboardPlayer
}

type LeaderboardPlayer struct {
	Tag  	  string
	Name 	  string
	NameColor string
	Icon      PlayerIcon
	Trophies  int
	Rank      int
	Club      LeaderBoardPlayerClub
}

type LeaderBoardPlayerClub struct {
	Name string
}

func (c *Client) GetCountryLeaderboardPlayers(countryCode string) (*CountryLeaderboardPlayers, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/rankings/" + countryCode + "/players", nil)
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
	var countryLeaderboard CountryLeaderboardPlayers
	err = json.NewDecoder(resp.Body).Decode(&countryLeaderboard)
	if err != nil {
		return nil, err
	}
	return &countryLeaderboard, nil
}
