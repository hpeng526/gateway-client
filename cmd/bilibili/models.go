package main

type BanGuMi struct {
	New          bool   `json:"new"`
	Title        string `json:"title"`
	WeekDay      int64  `json:"weekday"`
	LastUpdateAt string `json:"lastupdate_at"`
	LastUpdate   int64  `json:"lastupdate"`
	SeasonId     int64  `json:"season_id"`
	URL          string `json:"url"`
}

type BGMResp struct {
	Code    int64     `json:"code"`
	Count   string    `json:"count"`
	List    []BanGuMi `json:"list"`
	Message string    `json:"message"`
}
