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
type industryData []*models.StockData

type dataChannel struct {
	data models.IexBatchResponse
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

	channel := make(chan dataChannel)
	stockToIndustryMap := make(map[string]string)
	stockSeparatedList := ""

	currentBatch := 0
	batchCalls := 0
	for industry, symbols := range inputData {
		for index, symbol := range symbols {
			stockToIndustryMap[symbol] = industry
			stockSeparatedList += symbol + ","
			currentBatch++

			if currentBatch == 99 || index == len(symbols) - 1 {
				currentBatch = 0
				batchCalls++

				stockSeparatedList = stockSeparatedList[:len(stockSeparatedList) - 1]
				go fetchSet(stockSeparatedList, channel)

				stockSeparatedList = ""
			}
		}
	}

	fetchedData := make(data)
	for i := 0; i < batchCalls; i++ {
		channelData := <-channel
		if channelData.data == nil {
			continue
		}

		for symbol, iexResponse := range channelData.data {
			industry := stockToIndustryMap[symbol]

			industryDataForStock := fetchedData[industry]
			if industryDataForStock == nil {
				industryDataForStock = make(industryData, 5)
			}

			industryDataForStock = append(industryDataForStock, iexResponse.AsStockData())
			fetchedData[industry] = industryDataForStock
		}
	}

	body, err := json.Marshal(fetchedData)
	if err != nil {
		return nil
	}

	return body
}

func fetchSet(symbols string, channel chan dataChannel) {
	data := downloader.GetStockData(symbols)
	channel <- dataChannel{data:data}
}
