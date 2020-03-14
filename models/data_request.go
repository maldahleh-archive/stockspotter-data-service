package models

type DataRequest struct {
	Version string `json:"version"`
}

var DefaultRequest = DataRequest{Version: "1"}
