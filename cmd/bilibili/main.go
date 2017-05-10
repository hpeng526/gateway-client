package main

import (
	"encoding/json"
	"fmt"
	"gateway-client/utils"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	CACHE_FILE  = "./cache.data"
	GATEWAY_URL = "http://gateway/u"
)

func main() {
	resp, err := http.Get("https://www.bilibili.com/api_proxy?app=bangumi&action=timeline_v2")
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("get error: %v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var bgmResp BGMResp
	err = json.Unmarshal(body, &bgmResp)
	if err != nil {
		log.Fatalf("json error: %v", err)
	}
	tmpMap := make(map[interface{}]interface{})
	oldMap, _ := utils.LoadFromFile(CACHE_FILE)
	for _, bgm := range bgmResp.List {
		if bgm.New {
			if oldMap[bgm.SeasonId] != bgm.LastUpdate {
				msg := utils.BuildMsg(
					100,
					"http://www.bilibili.com"+bgm.URL,
					bgm.Title,
					"bilibili-new",
					fmt.Sprintf("day-%d", bgm.WeekDay),
					bgm.LastUpdateAt,
					"",
				)
				utils.SendToGateway(GATEWAY_URL, msg)
			}
			tmpMap[bgm.SeasonId] = bgm.LastUpdate
		}
	}
	utils.SaveToFile(CACHE_FILE, tmpMap)
}
