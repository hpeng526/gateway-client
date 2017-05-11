package main

type MZWeatherResp struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Value   []MZWeatherValue `json:"value"`
}

type MZWeatherValue struct {
	City               string            `json:"city"`
	ProvinceName       string            `json:"provinceName"`
	RealTime           MZWeatherRealTime `json:"realtime"`
	WeatherDetailsInfo MZWeatherDetail   `json:"weatherDetailsInfo"`
}

type MZWeatherDetail struct {
	PublishTime         string        `json:"publishTime"`
	Weather3HoursDetail []interface{} `json:"weather3HoursDetailsInfos"`
}

type MZWeatherRealTime struct {
	Temp  string `json:"temp"`
	Wind  string `json:"wD"`
	WindS string `json:"wS"`
}
