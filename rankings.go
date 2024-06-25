package brawlstars

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CountryLeaderboardResponse struct {
	Leaderboard  Leaderboard `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type Leaderboard []LeaderboardPlayer

type LeaderboardPlayer struct {
	Tag       string           `json:"tag"`
	Name      string           `json:"name"`
	NameColor string           `json:"nameColor"`
	Icon      PlayerIcon       `json:"icon"`
	Trophies  int              `json:"trophies"`
	Rank      int              `json:"rank"`
	Club      *LeaderboardClub `json:"club,omitempty"`
}

type LeaderboardClub struct {
	Name string `json:"name"`
}

func (c *Client) GetCountryLeaderboardPlayers(countryCode string) (Leaderboard, error) {
	url := fmt.Sprintf("https://bsproxy.royaleapi.dev/v1/rankings/%s/players", countryCode)
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

	var leaderboardResp CountryLeaderboardResponse
	err = json.NewDecoder(resp.Body).Decode(&leaderboardResp)
	if err != nil {
		return nil, err
	}

	return leaderboardResp.Leaderboard, nil
}

func (lb *Leaderboard) GetPlayerByTag(tag string) (*LeaderboardPlayer, error) {
	for _, player := range *lb {
		if player.Tag == tag {
			return &player, nil
		}
	}
	return nil, fmt.Errorf("player with tag %s not found", tag)
}
