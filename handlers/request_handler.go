package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/maldahleh/stockspotter-data-service/downloader"
	"github.com/maldahleh/stockspotter-data-service/downloader/models"
	"github.com/maldahleh/stockspotter-data-service/utils"
)

type stockInputFileStructure map[string][]string

type data map[string]industryData
type industryData map[string]*models.StockData

type industryDataChannel struct {
	industry string
	data     industryData
}

type stockDataChannel struct {
	symbol string
	data   *models.StockData
}

func FetchStocks(version string) []byte {
	filePath := "./versions/" + version + ".json"
	if !utils.FileExists(filePath) {
		filePath = "./versions/1.json"
	}

	file, _ := ioutil.ReadFile(filePath)
	inputData := make(stockInputFileStructure)

	err := json.Unmarshal(file, &inputData)
	if err != nil {
		return nil
	}

	fetchedData := make(data)
	channel := make(chan industryDataChannel)

	channelCalls := 0
	for industry, symbols := range inputData {
		channelCalls++
		go fetchIndustryData(industry, symbols, channel)
	}

	for i := 0; i < channelCalls; i++ {
		channelData := <-channel
		if channelData.data == nil {
			continue
		}

		fetchedData[channelData.industry] = channelData.data
	}

	body, err := json.Marshal(fetchedData)
	if err != nil {
		return nil
	}

	return body
}

func fetchIndustryData(industry string, symbols []string, industryChannel chan industryDataChannel) {
	data := make(industryData)
	channel := make(chan stockDataChannel)

	channelCalls := 0
	for _, symbol := range symbols {
		channelCalls++
		go fetchStock(symbol, channel)
	}

	for i := 0; i < channelCalls; i++ {
		channelData := <-channel
		if channelData.data == nil {
			continue
		}

		data[channelData.symbol] = channelData.data
	}

	industryChannel <- industryDataChannel{
		industry: industry,
		data:     data,
	}
}

func fetchStock(symbol string, channel chan stockDataChannel) {
	data := downloader.GetStockData(symbol)

	channel <- stockDataChannel{
		symbol: symbol,
		data:   data,
	}
}
