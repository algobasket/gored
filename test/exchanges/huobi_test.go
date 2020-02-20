package test

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/huobi"
	"github.com/bitontop/gored/test/conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Huobi(t *testing.T) {
	e := InitHuobi()

	pair := pair.GetPairByKey("BTC|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	// Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")

	// Test Withdraw
	Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")

}

func InitHuobi() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.HUOBI, config)

	ex := huobi.CreateHuobi(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}

func Test_Huobi_TradeHistory(t *testing.T) {
	e := InitHuobi()
	p := pair.GetPairByKey("BTC|ETH")

	opTradeHistory := &exchange.PublicOperation{
		Type:      exchange.TradeHistory,
		EX:        e.GetName(),
		Pair:      p,
		DebugMode: true,
	}

	err := e.LoadPublicData(opTradeHistory)
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("TradeHistory: %s::%s", opTradeHistory.EX, opTradeHistory.Pair.Name)

	for _, d := range opTradeHistory.TradeHistory {
		log.Printf(">> %+v ", d)
	}
}
