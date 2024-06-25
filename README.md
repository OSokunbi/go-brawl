# go-brawl

an unofficial brawl stars api wrapper for golang

## prerequisites

- a brawl stars api token, you can get one [here](https://developer.brawlstars.com/#/)

## inatallation

```bash
go get github.com/OSokunbi/go-brawl
```

## Usage

here's a simple example of how to use the library

```go
package main

import (
	"fmt"
	"github.com/OSokunbi/go-brawl"
	"os"
)

func main() {
	client := brawl.NewClient(os.Getenv("BRAWL_API_TOKEN"))
	player, err := client.GetPlayer("2Y9GJQJ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(player)

	leaderboard, err := client.GetCountryLeaderboardPlayers("12")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(leaderboard)
}
```

