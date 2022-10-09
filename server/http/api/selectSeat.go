package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/service"
	"golangProject/cinemaBuy/vo/request"
)

//选座信息接口
var SelectSeatApi = new(seat)

type seat struct{}

func (s *seat) GetSelectSeatInfo(r *ghttp.Request) {
	var req *request.MovieId
	r.Parse(&req)
	result := service.ChangeMovieInfo(req.Id["data"])
	r.Response.Write(&result)
}

//选座成功页面信息接口
func (s *seat) GetSelectOkInfo(r *ghttp.Request) {
	var req *request.MovieId
	r.Parse(&req)
	result := service.ChangeMovieCinemaInfo(req.Id["data"])
	r.Response.Write(&result)
}
