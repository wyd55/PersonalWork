package request

type Senddata struct {
	Aa int
}

//返回电影id
type FilmId struct {
	Id map[string]int
}

//返回影院id
type CinemaId struct {
	Id map[string]int
}

//返回档期id
type MovieId struct {
	Id map[string]int
}

//座位信息
type SeatArray struct {
	Seatok    []string
	Id        map[string]string
	UserId    int
	TicketNum int
	Pai       []string
	Zuo       []string
}

//code
type GetCode struct {
	Code string
}

//opneId
type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//返回userId
type UserId struct {
	UserId int
}

//确认订单界面请求数据
type ReqRecorOk struct {
	PurchaseId map[string]int
}

//收藏、看过请求
type ReqCollectSeen struct {
	FilmId int
	UserId int
}
