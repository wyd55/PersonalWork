package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/service"
	"golangProject/cinemaBuy/vo/request"
	"golangProject/cinemaBuy/vo/response"
)

//订单接口
var PurchaseApi = new(purchase)

type purchase struct{}

func (b *purchase) GetRecordList(r *ghttp.Request) {
	var req *request.UserId
	r.Parse(&req)
	var res []response.PiontCinemaInfo
	//通过用户id获取用户全部订单中的电影id
	film := repository.GetPurFilmIdByUserId(req.UserId)
	for _, value := range film {
		//返回电影信息在结构体
		repository.FilmInfo(gconv.Int(value["filmId"]))
		//是否可以退票或者取票
		response.CinemaInfo.Button = service.GetTicketButton(gconv.Int(value["movieId"]), gconv.Int(value["purchaseId"]))
		response.CinemaInfo.PurchaseId = gconv.Int(value["purchaseId"])
		res = append(res, *response.CinemaInfo)
	}
	r.Response.Write(res)
}

func (b *purchase) GetRecordOkInfo(r *ghttp.Request) {
	var req *request.ReqRecorOk
	r.Parse(&req)
	film := repository.GetMovieIdByPurchaseId(req.PurchaseId["data"])
	result := service.ChangeMovieCinemaInfo(gconv.Int(film["movieId"]))
	//把返回的结构体换成map
	res := gconv.Map(result)
	seatOkSli := gstr.Split(gconv.String(film["seatok"]), ",")
	paiSli := gstr.Split(gconv.String(film["pai"]), ",")
	zuoSli := gstr.Split(gconv.String(film["zuo"]), ",")
	//往map中添加number(购票数量),seatok,pai,zuo
	res["number"] = film["number"]
	res["seatok"] = seatOkSli
	res["pai"] = paiSli
	res["zuo"] = zuoSli
	r.Response.Write(res)
}

//取票
func (b *purchase) QuPiao(r *ghttp.Request) {
	var req *request.ReqRecorOk
	r.Parse(&req)
	result := repository.ChangeCheckIn(req.PurchaseId["data"])
	r.Response.Write(result)
}

//退票
func (b *purchase) TuiPiao(r *ghttp.Request) {
	var req *request.ReqRecorOk
	r.Parse(&req)
	//修改对应档期的座位表
	film := repository.GetMovieIdByPurchaseId(req.PurchaseId["data"])
	seatok := gstr.Split(gconv.String(film["seatok"]), ",")
	movieId := gconv.Int(film["movieId"])
	//根据位置信息修改座位
	result := repository.ChangeSeatInMovie(seatok, movieId)
	if !result {
		r.Response.Write("false")
	}
	//删除用户订单
	result2 := repository.DeletePurchase(req.PurchaseId["data"])
	if !result2 {
		r.Response.Write("false")
	}
	r.Response.Write("true")
}
