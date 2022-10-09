package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/vo/request"
	"golangProject/cinemaBuy/vo/response"
)

//电影介绍接口
var FilmInfoApi = new(filmInfo)

type filmInfo struct{}

func (f *filmInfo) GetFilmInfo(r *ghttp.Request) {
	//返回电影Id
	var req *request.FilmId
	r.Parse(&req)
	//返回电影信息在结构体
	repository.FilmInfo(req.Id["data"])
	r.Response.Write(&response.CinemaInfo)
}

//电影演员接口

func (f *filmInfo) GetActorInfo(r *ghttp.Request) {
	//返回电影演员信息
	var req *request.FilmId
	r.Parse(&req)
	actorInfo := repository.ActorInfo(req.Id["data"])
	r.Response.Write(actorInfo)
}
