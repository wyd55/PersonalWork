package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

//根据时间、影院id获取对应的电影图片（购买页面使用）
func GetFilmPic(cinemaId int) gdb.Result {
	now := time.Now().Format("2006-01-02")
	time := time.Now().Format("15:04:05")
	film, _ := g.Model("movie").InnerJoin("film", "movie.filmId=film.filmId").Fields("film.filmPic,film.filmId").Where("movie.date=? And movie.startTime>=? or movie.date>?", now, time, now).Where("movie.cinemaId=?", cinemaId).Group("movie.filmId").All()
	return film
}

//通过movieId得到对应的电影名字
func GetFilmNameByMovieId(movieId int) string {
	film, _ := g.Model("movie").InnerJoin("film", "movie.filmId=film.filmId").Fields("film.filmName").Where("movieId=?", movieId).Value()
	return gconv.String(film)
}
