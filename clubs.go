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
	url := fmt.Sprintf("https://bsproxy.royaleapi.dev/v1/clubs/%s", url.PathEscape(tag))
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
	url := fmt.Sprintf("https://bsproxy.royaleapi.dev/v1/clubs/%s/members", url.PathEscape(tag))
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

func (c *Club) GetMembers() []ClubMember {
	return c.Members
}

func (c *Club) GetMember(tag string) *ClubMember {
	for _, member := range c.Members {
		if member.Tag == tag {
			return &member
		}
	}
	return nil
}

func (c *Club) GetMemberByIndex(index int) *ClubMember {
	if index >= 0 && index < len(c.Members) {
		return &c.Members[index]
	}
	return nil
}

func (c *Club) GetMemberCount() int {
	return len(c.Members)
}

func (c *Club) GetMemberTags() []string {
	tags := make([]string, len(c.Members))
	for i, member := range c.Members {
		tags[i] = member.Tag
	}
	return tags
}

func (c *Club) GetMemberNames() []string {
	names := make([]string, len(c.Members))
	for i, member := range c.Members {
		names[i] = member.Name
	}
	return names
}

func (c *Club) GetMemberByName(name string) *ClubMember {
	for _, member := range c.Members {
		if member.Name == name {
			return &member
		}
	}
	return nil
}

func (c *Club) GetMemberByRole(role string) []ClubMember {
	members := make([]ClubMember, 0)
	for _, member := range c.Members {
		if member.Role == role {
			members = append(members, member)
		}
	}
	return members
}	

func (c *Club) GetMemberByTrophies(trophies int) []ClubMember {
	members := make([]ClubMember, 0)
	for _, member := range c.Members {
		if member.Trophies >= trophies {
			members = append(members, member)
		}
	}
	return members
}
