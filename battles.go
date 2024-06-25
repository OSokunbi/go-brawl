package brawlstars

type BattleLog struct {
  Battles []Battle
}
type Battle struct {
  Time  string
  Event Event
  Info  BattleInfo
}

type Event struct {
  ID   string
  Mode string
  Map  string
}

type BattleInfo struct {
  Mode       string
  Type       string
  Result     string
  Duration   int
  StarPlayer *BattlePlayer
  Teams      []Team
}

type Team struct {
  Team []BattlePlayer
}

type BattlePlayer struct {
  Tag     string
  Name    string
  Brawler BattleBrawler
}


type BattleBrawler struct {
  Tag      string
  Name     string
  Power    int
  Trophies int
}
