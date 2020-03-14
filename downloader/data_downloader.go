package downloader

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maldahleh/stockspotter-data-service/downloader/models"
)

const root = "https://cloud.iexapis.com/v1/stock/"

func GetStockData(symbol string) *models.StockData {
	var iexData models.IexData
	err := getJSON(createUrl(symbol), &iexData)
	if err != nil {
		return nil
	}

	return iexData.AsStockData()
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

func createUrl(symbol string) string {
	return root + symbol + "/quote?displayPercent=true&token=sk_cf7d9c2fdd9447fc8204eec400bd5b9b"
}
