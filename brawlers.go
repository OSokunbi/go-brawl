package brawlstars

import (
	"net/http"
	"encoding/json"
)

type Brawler struct {
	ID   	   int
	Name 	   string
	StarPowers []StarPower
  	Gadgets    []Gadget
}

type StarPower struct {
	ID   int
	Name string
}

type Gadget struct {
	ID 	 int
	Name string
}


func (c *Client) GetBrawlers () (*[]Brawler, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/brawlers", nil)
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
	var brawlers []Brawler
	err = json.NewDecoder(resp.Body).Decode(&brawlers)
	if err != nil {
		return nil, err
	}
	return &brawlers, nil
}

func (c *Client) GetBrawler (id string) (*Brawler, error) {
	req, err := http.NewRequest("GET", "https://api.brawlstars.com/v1/brawlers/%23" + id, nil)
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
	var brawler Brawler
	err = json.NewDecoder(resp.Body).Decode(&brawler)
	if err != nil {
		return nil, err
	}
	return &brawler, nil
}