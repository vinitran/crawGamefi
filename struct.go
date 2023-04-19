package main

import "time"

type DataCrawl struct {
	Context    string `json:"@context"`
	Type       string `json:"@type"`
	MainEntity []struct {
		Type           string `json:"@type"`
		Name           string `json:"name"`
		AcceptedAnswer struct {
			Type string `json:"@type"`
			Text string `json:"text"`
		} `json:"acceptedAnswer"`
	} `json:"mainEntity"`
}

type DataAll struct {
	Payload struct {
		Games []struct {
			Code        string `json:"code"`
			Name        string `json:"name"`
			Rank        int    `json:"rank"`
			ImageURL    string `json:"imageUrl"`
			Symbol      string `json:"symbol"`
			BlockChains []struct {
				Name        string `json:"name"`
				Code        string `json:"code"`
				ExtendValue string `json:"extendValue"`
			} `json:"blockChains"`
			Genres []struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"genres"`
			Platforms []struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"platforms"`
			Price                             float64     `json:"price"`
			Price24HChangePercent             float64     `json:"price24hChangePercent"`
			PriceChangePercentage7DInCurrency float64     `json:"priceChangePercentage7dInCurrency"`
			PriceChangePercentage1HInCurrency float64     `json:"priceChangePercentage1hInCurrency"`
			Volume24H                         interface{} `json:"volume24h"`
			MarketCap                         interface{} `json:"marketCap"`
			SparkLine7DURL                    interface{} `json:"sparkLine7dUrl"`
			DateAdded                         time.Time   `json:"dateAdded"`
			SparkLine7D                       interface{} `json:"sparkLine7d"`
			IsWatched                         interface{} `json:"isWatched"`
			IsInMainWatchList                 interface{} `json:"isInMainWatchList"`
			MainWatchListID                   interface{} `json:"mainWatchListId"`
		} `json:"games"`
		Summary interface{} `json:"summary"`
	} `json:"payload"`
	Code    int `json:"code"`
	Locale  any `json:"locale"`
	Message any `json:"message"`
}
