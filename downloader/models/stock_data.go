package models

type BatchStockResponse map[string]StockData

type StockData struct {
	Price                 float64 `json:"price"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	Volume                float64 `json:"volume"`
	MarketCap             float64 `json:"cap"`
}
