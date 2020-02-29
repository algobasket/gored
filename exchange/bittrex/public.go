package bittrex

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/utils"
)

func (e *Bittrex) LoadPublicData(operation *exchange.PublicOperation) error {
	switch operation.Type {

	case exchange.TradeHistory:
		return e.doTradeHistory(operation)
	case exchange.CoinChainType:
		return e.getCoinChainType(operation)

	}
	return fmt.Errorf("LoadPublicData :: Operation type invalid: %+v", operation.Type)
}

func (e *Bittrex) doTradeHistory(operation *exchange.PublicOperation) error {
	symbol := e.GetSymbolByPair(operation.Pair)

	get := &utils.HttpGet{
		URI: fmt.Sprintf("%s/v1.1/public/getmarkethistory?market=%s", API_URL, symbol),
	}

	err := utils.HttpGetRequest(get)

	if err != nil {
		return err

	} else {
		jsonResponse := JsonResponse{}
		tradeHistory := TradeHistory{}
		if err := json.Unmarshal(get.ResponseBody, &jsonResponse); err != nil {
			return err
		}

		if err := json.Unmarshal(jsonResponse.Result, &tradeHistory); err != nil {
			return err
		}

		operation.TradeHistory = []*exchange.TradeDetail{}
		// for _, d := range tradeHistory {
		for i := len(tradeHistory) - 1; i > 0; i-- {
			d := tradeHistory[i]
			td := &exchange.TradeDetail{}

			td.ID = fmt.Sprintf("%d", d.ID)
			if d.OrderType == "BUY" {
				td.Direction = exchange.Buy
			} else if d.OrderType == "SELL" {
				td.Direction = exchange.Sell
			}

			td.Quantity = d.Quantity
			td.Rate = d.Price

			layout := "2006-01-02T15:04:05.00"
			t, err := time.Parse(layout, d.TimeStamp)
			if err != nil {
				// log.Printf("%+v", err)
				layout = "2006-01-02T15:04:05.0"
				t, err = time.Parse(layout, d.TimeStamp)
				if err != nil {
					// log.Printf("%+v", err)
					layout = "2006-01-02T15:04:05"
					t, err = time.Parse(layout, d.TimeStamp)
				}
			}

			td.TimeStamp = t.UnixNano() / 1e6

			operation.TradeHistory = append(operation.TradeHistory, td)
		}
	}

	return nil
}

func (e *Bittrex) getCoinChainType(operation *exchange.PublicOperation) error {
	operation.CoinChainType = []exchange.ChainType{}
	request := &exchange.ChainTypeRequest{
		Exchange: string(operation.EX),
		CoinID:   operation.Coin.ID,
	}

	byteJson, err := json.Marshal(request)
	post := &utils.HttpPost{
		URI:         "http://127.0.0.1:52020/getchaintype",
		RequestBody: byteJson,
	}

	err = utils.HttpPostRequest(post)
	if err != nil {
		return err

	} else {
		chainType := []*exchange.ChainTypeRequest{}
		if err := json.Unmarshal(post.ResponseBody, &chainType); err != nil {
			return err
		}

		for _, data := range chainType {
			for _, ct := range data.ChainType {
				switch ct {
				case "MAINNET":
					operation.CoinChainType = append(operation.CoinChainType, exchange.MAINNET)
				case "BEP2":
					operation.CoinChainType = append(operation.CoinChainType, exchange.BEP2)
				case "ERC20":
					operation.CoinChainType = append(operation.CoinChainType, exchange.ERC20)
				case "NEP5":
					operation.CoinChainType = append(operation.CoinChainType, exchange.NEP5)
				case "OMNI":
					operation.CoinChainType = append(operation.CoinChainType, exchange.OMNI)
				case "TRC20":
					operation.CoinChainType = append(operation.CoinChainType, exchange.TRC20)
				default:
					operation.CoinChainType = append(operation.CoinChainType, exchange.OTHER)
				}
			}
		}
	}

	return nil
}
