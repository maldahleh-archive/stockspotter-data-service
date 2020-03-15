package downloader

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/maldahleh/stockspotter-data-service/downloader/models"
)

func GetStockData(symbols string) models.IexBatchResponse {
	var iexData models.IexBatchResponse
	err := getJSON(createUrl(symbols), &iexData)
	if err != nil {
		return nil
	}

	return iexData
}

func getJSON(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}

	return nil
}

func createUrl(symbols string) string {
	var apiKey, _ = os.LookupEnv("API_KEY")

	return "https://cloud.iexapis.com/v1/stock/market/batch?types=quote&displayPercent=true&token=" + apiKey +
		"&symbols=" + symbols
}
