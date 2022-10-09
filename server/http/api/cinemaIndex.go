package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/service"
	"golangProject/cinemaBuy/vo/request"
)

var CinemaIndexApi = new(cinemaIndex)

type cinemaIndex struct{}

//影院首页信息-影院信息
func (c *cinemaIndex) GetCinemaIndexInfo(r *ghttp.Request) {
	var req *request.CinemaId
	r.Parse(&req)
	film := repository.GetSomeCinemaInfo(req.Id["data"])
	r.Response.Write(film)
}

//返回影院页面的热映电影图片页面
func (c *cinemaIndex) GetCinemaIndexPic(r *ghttp.Request) {
	var req *request.CinemaId
	r.Parse(&req)
	film := repository.GetFilmPic(req.Id["data"])
	film2 := service.AlterStringPic(film)
	r.Response.Write(film2)
}

//返回影院页面电影档期（可购票的电影）
func (c *cinemaIndex) GetCinemaIndexMovie(r *ghttp.Request) {
	var req *request.CinemaId
	r.Parse(&req)
	film := service.GetMovieByCinemaFilm(req.Id["data"])
	r.Response.Write(film)
}
