package main

import (
	"fmt"
	"irc-20/api"
	"strconv"
)

type MyData struct {
	Value float64 `json:"value"`
}

func main() {

	brcList, err := api.RequestBrc20()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filterBrc := FilterAndMap(brcList)
	for _, detail := range filterBrc {
		fmt.Printf("%+v\n", detail)

	}
}

func FilterAndMap(details []api.BRC20Detail) []api.BRC20Detail {
	tickerMap := make(map[string]float64)

	filteredDetails := make([]api.BRC20Detail, 0)
	for _, detail := range details {
		confirmedMinted1h, _ := strconv.ParseFloat(detail.ConfirmedMinted1h, 64)
		MaxMinted, _ := strconv.ParseFloat(detail.Max, 64)
		hConfirmed := confirmedMinted1h / MaxMinted
		if detail.HoldersCount >= 20 && hConfirmed > 0.20 && hConfirmed > tickerMap[detail.Ticker] {
			tickerMap[detail.Ticker] = hConfirmed
			filteredDetails = append(filteredDetails, detail)
		}
	}

	return filteredDetails
}
