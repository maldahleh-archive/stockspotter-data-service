package models

type IexData struct {
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

func (data IexData) AsStockData() *StockData {
	return &StockData{
		Price:                 data.Price,
		Change:                data.Change,
		ChangePercent:         data.ChangePercent,
		ExtendedPrice:         data.ExtendedPrice,
		ExtendedChange:        data.ExtendedChange,
		ExtendedChangePercent: data.ExtendedChangePercent,
		Volume:                data.Volume,
		MarketCap:             data.MarketCap,
		MarketOpen:            data.MarketOpen,
	}
}
