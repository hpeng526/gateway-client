package main

import (
	"encoding/json"
	"fmt"
	"gateway-client/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	CITY_ID     = "101280601"
	GATEWAY_URL = "http://gateway/u"
)

func main() {
	weatherURL := "http://aider.meizu.com/app/weather/listWeather?cityIds=" + CITY_ID
	resp, err := http.Get(weatherURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("get error: %v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonData MZWeatherResp
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Fatalf("json error: %v", err)
	}
	fmt.Println(jsonData)
	weatherBytes, err := json.Marshal(jsonData.Value[0].WeatherDetailsInfo.Weather3HoursDetail)
	weatherStr := string(weatherBytes)
	title := jsonData.Value[0].City + "," + jsonData.Value[0].ProvinceName + " " + jsonData.Value[0].RealTime.Temp + "°C " + jsonData.Value[0].RealTime.Wind + " " + jsonData.Value[0].RealTime.WindS
	content := "不用带遮或收衫"
	if strings.Contains(weatherStr, "雨") {
		content = "落雨收衫啊"
	}

	msg := utils.BuildMsg(
		100,
		weatherURL,
		title,
		"weather-notify",
		content,
		jsonData.Value[0].WeatherDetailsInfo.PublishTime,
		"",
	)
	utils.SendToGateway(GATEWAY_URL, msg)
}
