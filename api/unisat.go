package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Height int           `json:"height"`
	Total  int           `json:"total"`
	Start  int           `json:"start"`
	Detail []BRC20Detail `json:"detail"`
}

type BRC20Detail struct {
	Ticker                 string `json:"ticker"`
	HoldersCount           int    `json:"holdersCount"`
	HistoryCount           int    `json:"historyCount"`
	InscriptionNumber      int    `json:"inscriptionNumber"`
	InscriptionId          string `json:"inscriptionId"`
	Max                    string `json:"max"`
	Limit                  string `json:"limit"`
	Minted                 string `json:"minted"`
	TotalMinted            string `json:"totalMinted"`
	ConfirmedMinted        string `json:"confirmedMinted"`
	ConfirmedMinted1h      string `json:"confirmedMinted1h"`
	ConfirmedMinted24h     string `json:"confirmedMinted24h"`
	MintTimes              int    `json:"mintTimes"`
	Decimal                int    `json:"decimal"`
	Creator                string `json:"creator"`
	Txid                   string `json:"txid"`
	DeployHeight           int    `json:"deployHeight"`
	DeployBlocktime        int    `json:"deployBlocktime"`
	CompleteHeight         int    `json:"completeHeight"`
	CompleteBlocktime      int    `json:"completeBlocktime"`
	InscriptionNumberStart int    `json:"inscriptionNumberStart"`
	InscriptionNumberEnd   int    `json:"inscriptionNumberEnd"`
}

func RequestBrc20() ([]BRC20Detail, error) {
	resp, err := http.Get("https://unisat.io/brc20-api-v2/brc20/status?ticker=&start=0&limit=400&complete=no&sort=minted")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, fmt.Errorf("API request failed with code: %d, message: %s", response.Code, response.Msg)
	}

	return response.Data.Detail, nil
}
