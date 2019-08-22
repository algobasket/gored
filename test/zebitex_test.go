package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"fmt"
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/zebitex"
	"github.com/bitontop/gored/test/conf"
	// "../exchange/zebitex"
	// "./conf"
)

/********************Public API********************/
func Test_Zebitex(t *testing.T) {
	e := InitZebitex()

	pair := pair.GetPairByKey("BTC|ETH")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Trading_Sell(e, pair, 0.00000001, 667)
	// Test_Withdraw(e, pair.Base, 1.5, "1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v")
}

func InitZebitex() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.ZEBITEX, config)
	fmt.Printf("%+v\n", config)

	ex := zebitex.CreateZebitex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
