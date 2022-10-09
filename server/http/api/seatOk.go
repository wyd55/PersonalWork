package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/vo/request"
)

//下单接口
var SeatOkApi = new(seatok)

type seatok struct{}

func (s *seatok) ChangeSeat(r *ghttp.Request) {
	var req *request.SeatArray
	r.Parse(&req)
	//修改座位
	repository.ChangeMovieSeat(req.Seatok, gconv.Int(req.Id["data"]))
	//给用户添加订单
	repository.InsertPurchase(req.UserId, gconv.Int(req.Id["data"]), req.Seatok, req.Pai, req.Zuo, req.TicketNum)
}
