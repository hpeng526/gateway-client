package utils

type MsgTemplate struct {
	Id   int64   `json:"id"`
	URL  string  `json:"url"`
	Data MsgData `json:"data"`
}

type MsgData struct {
	First  MsgVC `json:"first"`
	Send   MsgVC `json:"send"`
	Text   MsgVC `json:"text"`
	Time   MsgVC `json:"time"`
	Remark MsgVC `json:"remark,omitempty"`
}

type MsgVC struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
