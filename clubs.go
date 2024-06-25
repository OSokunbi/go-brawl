package brawlstars

import (
	"net/http"
	"encoding/json"
)
type Club struct {
  Tag              string
  Name             string
  Description      string
  Type             string
  BadgeID          int
  RequiredTrophies int
  Trophies         int
  Members          Members
} 

type Members struct {
  ClubMembers  []ClubMember
}

type ClubMember struct {
  Tag       string
  Name      string
  NameColor string
  Role      string
  Trophies  int
  Icon      PlayerIcon
}

func (c *Client) GetClub(tag string) (*Club, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/clubs/%23" + tag, nil)
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
	var club Club
	err = json.NewDecoder(resp.Body).Decode(&club)
	if err != nil {
		return nil, err
	}
	return &club, nil
}

func (c *Client) GetClubMembers(tag string) (*Members, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/clubs/%23" + tag + "/members", nil)
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
	var members Members
	err = json.NewDecoder(resp.Body).Decode(&members)
	if err != nil {
		return nil, err
	}
	return &members, nil
}
