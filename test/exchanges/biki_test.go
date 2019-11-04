package test

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/biki"
	"github.com/bitontop/gored/test/conf"
	// "../exchange/biki"
	// "./conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Biki(t *testing.T) {
	e := InitBiki()

	pair := pair.GetPairByKey("BTC|ETH")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.01234567, 0.12345678)
	// Test_Trading_Sell(e, pair, 0.02642640, 0.05700000)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}

func InitBiki() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.BIKI, config)

	ex := biki.CreateBiki(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
