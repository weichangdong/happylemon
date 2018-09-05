package entity

const RetCodeOk = 0
const RetCodeError = 1
const MsgOk = "ok"
const MsgError = "error"
const MsgJsonError = "json error"
const MsgApiGolangError = "api golang error"
const MsgParaError = "para error"

const MsgNotExists = "data is not exists"
const MsgReExists = "data is exists"

type QueueData struct {
	Act  string       `json:"act"`
	Data QueueSonData `json:"data"`
}

type QueueSonData struct {
	Uid        string  `json:"uid"`
	Gid        float64 `json:"gid"`
	Msg        string  `json:"msg"`
	Type       float64 `json:"type"`
	Num        int     `json:"num"`
	ImgList    []UpImg `json:"imgList"`
	OldFeedsId int64   `json:"oldFeedsId"`
}

type UpImg struct {
	Img    string
	IsLink bool
	Seq    int
	Suffix string
}
