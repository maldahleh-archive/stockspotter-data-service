package models

type IexBatchResponse map[string]IexOuterStruct

type IexOuterStruct struct {
	Quote IexData `json:"quote"`
}

type IexData struct {
	Symbol                string  `json:"symbol"`
	Price                 float64 `json:"latestPrice"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	Volume                float64 `json:"latestVolume"`
	MarketCap             float64 `json:"marketCap"`
	MarketOpen            bool    `json:"isUSMarketOpen"`
}

func (data IexOuterStruct) AsStockData() *StockData {
	if data.Quote.MarketOpen || data.Quote.ExtendedChange == 0 {
		return &StockData{
			Symbol:        data.Quote.Symbol,
			Price:         data.Quote.Price,
			Change:        data.Quote.Change,
			ChangePercent: data.Quote.ChangePercent,
			Volume:        data.Quote.Volume,
			MarketCap:     data.Quote.MarketCap,
		}
	}

	return &StockData{
		Symbol:        data.Quote.Symbol,
		Price:         data.Quote.ExtendedPrice,
		Change:        data.Quote.ExtendedChange,
		ChangePercent: data.Quote.ExtendedChangePercent,
		Volume:        data.Quote.Volume,
		MarketCap:     data.Quote.MarketCap,
	}
}
