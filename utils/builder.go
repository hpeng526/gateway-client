package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"log"
)

func BuildMsg(userId int64, url string, first string, send string, text string, time string, remark string) (msg *MsgTemplate) {
	msg = &MsgTemplate{
		URL:    url,
		UserId: userId,
		Data: MsgData{
			First: MsgVC{Value: first, Color: "#172177"},
			Send:  MsgVC{Value: send, Color: "#172177"},
			Text:  MsgVC{Value: text, Color: "#172177"},
			Time:  MsgVC{Value: time, Color: "#172177"},
		},
	}
	return
}

func SendToGateway(url string, msg *MsgTemplate) (resp *http.Response, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(&msg)
	if err != nil {
		log.Fatalf("encode json error: %v\n", err)
	}
	log.Printf("%s\n", b)
	resp, err = http.Post(url, "application/json; charset=utf-8", b)
	log.Printf("resp is %v\n", resp)
	return
}
