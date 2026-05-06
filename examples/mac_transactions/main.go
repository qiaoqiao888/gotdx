package main

import (
	"log"

	"github.com/bensema/gotdx"
	"github.com/bensema/gotdx/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.NewMACClient()
	defer client.Disconnect()

	items, err := client.MACTransactions(gotdx.MarketSZ, "000001", 0, 20)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range items {
		log.Printf("mac_transactions time=%s price=%.2f vol=%d trade_count=%d action=%d",
			item.Time, item.Price, item.Vol, item.TradeCount, item.BuyOrSell)
	}
}
