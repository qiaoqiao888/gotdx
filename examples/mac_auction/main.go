package main

import (
	"log"

	"github.com/bensema/gotdx"
	"github.com/bensema/gotdx/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.NewMACClient()
	defer client.Disconnect()

	items, err := client.MACAuction(gotdx.MarketSZ, "000001", 0, 20)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range items {
		log.Printf("mac_auction time=%s price=%.2f matched=%d unmatched=%d flag=%d",
			item.Time, item.Price, item.Matched, item.Unmatched, item.Flag)
	}
}
