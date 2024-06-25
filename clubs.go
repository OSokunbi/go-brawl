package brawlstars

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Club struct {
	Tag              string       `json:"tag"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Type             string       `json:"type"`
	BadgeID          int          `json:"badgeId"`
	RequiredTrophies int          `json:"requiredTrophies"`
	Trophies         int          `json:"trophies"`
	Members          []ClubMember `json:"members"`
}

type ClubMembersResponse struct {
	Members []ClubMember `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type ClubMember struct {
	Icon      PlayerIcon `json:"icon"`
	Tag       string     `json:"tag"`
	Name      string     `json:"name"`
	Trophies  int        `json:"trophies"`
	Role      string     `json:"role"`
	NameColor string     `json:"nameColor"`
}

func (c *Client) GetClub(tag string) (*Club, error) {
	url := fmt.Sprintf("https://api.brawlstars.com/v1/clubs/%s", url.PathEscape(tag))
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

	var club Club
	err = json.NewDecoder(resp.Body).Decode(&club)
	if err != nil {
		return nil, err
	}

	return &club, nil
}

func (c *Client) GetClubMembers(tag string) ([]ClubMember, error) {
	url := fmt.Sprintf("https://api.brawlstars.com/v1/clubs/%s/members", url.PathEscape(tag))
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

	var clubMembersResp ClubMembersResponse
	err = json.NewDecoder(resp.Body).Decode(&clubMembersResp)
	if err != nil {
		return nil, err
	}

	return clubMembersResp.Members, nil
}