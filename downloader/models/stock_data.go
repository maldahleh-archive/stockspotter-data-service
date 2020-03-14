package models

type StockData struct {
	Price                 float64 `json:"price"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	Volume                float64 `json:"volume"`
	MarketCap             float64 `json:"cap"`
	MarketOpen            bool    `json:"open"`
}
