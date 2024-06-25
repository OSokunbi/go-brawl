package brawlstars

type BattleLogResponse struct {
	BattleLogs []Battle `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type Battle struct {
  Time  string     `json:"battleTime"`
  Event *Event      `json:"event,omitempty"`
  Info  BattleInfo `json:"battle"`
}

type Event struct {
  ID   string `json:"id"`
  Mode string `json:"mode"`
  Map  string `json:"map"`
}

type BattleInfo struct {
  Mode       string 	   `json:"mode"`
  Type       string 	   `json:"type"`
  Result     string 	   `json:"result"`
  Duration   int    	   `json:"duration"`
  StarPlayer *BattlePlayer `json:"starPlayer"`
  Teams      []Team 	   `json:"teams"`
}

type Team struct {
  Team []BattlePlayer `json:"team"`
}

type BattlePlayer struct {
  Tag     string 		`json:"tag"`
  Name    string 		`json:"name"`
  Brawler BattleBrawler `json:"brawler"`
}


type BattleBrawler struct {
  Tag      string `json:"id"`
  Name     string `json:"name"`
  Power    int    `json:"power"`
  Trophies int	  `json:"trophies"`
}
