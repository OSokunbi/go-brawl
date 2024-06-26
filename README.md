# go-brawl

an unofficial brawl stars api wrapper for golang

it is incomplete and still in development, feel free to contribute

## prerequisites

- a brawl stars api token, you can get one [here](https://developer.brawlstars.com/#/)

a proxy server is used to handle requests, when registering for an api token, you will need to provide the proxy server's ip address:
`45.79.218.79`

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

