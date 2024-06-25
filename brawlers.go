package brawlstars

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Brawler struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	StarPowers []StarPower `json:"starPowers"`
	Gadgets    []Gadget    `json:"gadgets"`
}

type StarPower struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Gadget struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BrawlersResponse struct {
	Brawlers []Brawler `json:"items"`
	Paging struct {
		Cursors struct {
			After string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) GetBrawlers() ([]Brawler, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/brawlers", nil)
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

	var brawlersResp BrawlersResponse
	err = json.NewDecoder(resp.Body).Decode(&brawlersResp)
	if err != nil {
		return nil, err
	}

	return brawlersResp.Brawlers, nil
}

func (c *Client) GetBrawler(id int) (*Brawler, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.brawlstars.com/v1/brawlers/%d", id), nil)
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

	var brawler Brawler
	err = json.NewDecoder(resp.Body).Decode(&brawler)
	if err != nil {
		return nil, err
	}

	return &brawler, nil
}