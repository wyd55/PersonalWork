package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/service"
	"golangProject/cinemaBuy/vo/request"
)

//影院信息接口
var CinemaApi = new(cinema)

type cinema struct{}

func (c *cinema) GetCinemaInfo(r *ghttp.Request) {
	r.Response.Write(repository.GetCinemaInfo())
}

//电影五天内有档期的影院
func (c *cinema) GetCinemaByFilm(r *ghttp.Request) {
	var req *request.FilmId
	r.Parse(&req)
	cinemaByFilm := service.GetCinemaByFilm(req.Id["data"])
	r.Response.Write(cinemaByFilm)
}


