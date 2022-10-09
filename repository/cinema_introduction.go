package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/vo/response"
)

//电影信息
func FilmInfo(filmId int) {
	var res = response.CinemaInfo
	film, _ := g.Model("film").Where("filmId = ?", filmId).One()
	gconv.Struct(film, &res)
}

//电影演员
func ActorInfo(filmId int) interface{} {
	film, _ := g.Model("actor").InnerJoin("filmactor", "actor.actorId = filmactor.actorId").Fields("actorPic,actorName,roleName").Where("filmId = ?", filmId).All()
	return film
}

//获得电影的演员角色对应的真实姓名
func ActorName(filmId string) gdb.Result {
	film, _ := g.Model("filmactor").InnerJoin("actor", "actor.actorId = filmactor.actorId").Fields("actorName,roleName").Where("filmId=?", filmId).All()
	return film
}
