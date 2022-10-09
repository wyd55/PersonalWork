package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/service"
)

//更多热映/预售
var MoreFilmApi = new(moreFilm)

type moreFilm struct{}

//热映
func (m *moreFilm) MoreFilm(r *ghttp.Request) {
	var act []map[string]string
	film := repository.GetReYinPic()
	for _, value := range film {
		film2 := repository.ActorName(gconv.String(value["filmId"]))
		zhumap := service.AlterString(film2)
		act = append(act, zhumap)
	}
	r.Response.Write(act)
}

//预售
func (m *moreFilm) MoreFilm2(r *ghttp.Request) {
	var act []map[string]string
	film := repository.GetYuShou()
	for _, value := range film {
		film2 := repository.ActorName(gconv.String(value["filmId"]))
		zhumap := service.AlterString(film2)
		act = append(act, zhumap)
	}
	r.Response.Write(act)
}

//显示五天日期
func (m *moreFilm) GetFiveDate(r *ghttp.Request) {
	r.Response.Write(service.GetFiveDate())
}
