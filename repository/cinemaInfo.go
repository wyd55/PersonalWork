package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//影院信息
func GetCinemaInfo() gdb.Result {
	film, _ := g.Model("cinema").All()
	return film
}

//获得指定影院的信息
func GetSomeCinemaInfo(cinemaid int) gdb.Result {
	film, _ := g.Model("cinema").Where("cinemaId=?", cinemaid).All()
	return film
}

//根据电影id和日期查找今天开始五天内有该电影档期的影院
func GetCinemaByFilm(filmId int, date string, time string) gdb.Result {
	film, _ := g.Model("movie").InnerJoin("cinema", "cinema.cinemaId=movie.cinemaId").Fields("cinema.*,movie.cinemaId").Where("movie.filmId=?", filmId).Where("movie.date=? and movie.startTime>?", date, time).Group("movie.cinemaId").All()
	return film
}
