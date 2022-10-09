package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/vo/response"
)

//根据影院id，电影id、和今天的日期查到对应档期
func GetMovieInfo(filmId int, cinemaId int, date string, time string) gdb.Result {
	film, _ := g.Model("movie").Where("filmId=? and cinemaId=?", filmId, cinemaId).Where("date=? and startTime>?", date, time).All()
	return film
}

//根据movieid返回选座信息
func GetSelectSeatInfo(movieId int) *response.MovieInfo {
	var res = response.MovieInfoPoint
	film, _ := g.Model("movie").Where("movieId=?", movieId).One()
	gconv.Struct(film, &res)
	return res
}

//根据movieid返回选座成功信息
func GetSelectOkInfo(movieId int) gdb.Record {
	film, _ := g.Model("movie").InnerJoin("cinema", "movie.cinemaId=cinema.cinemaId").Fields("cinema.Name,cinema.Location,movie.filmPrice,movie.date,movie.startTime,movie.room").Where("movieId=?", movieId).One()
	return film
}

//根据movieId返回档期的开始日期时间和结束日期时间
func GetStartAndEndTime(movieId int) gdb.Record {
	film, _ := g.Model("movie").Fields("date,startTime,endTime").Where("movieId=?", movieId).One()
	return film
}

//根据seatok和movieid修改退票后的座位表
func ChangeSeatInMovie(seatok []string, movieId int) bool {
	film, _ := g.Model("movie").Fields("seat").Where("movieId=?", movieId).One()
	seat := gstr.Split(gconv.String(film["seat"]), ",")
	for _, value := range seatok {
		seat[gconv.Int(value)] = "1"
	}
	StrSeat := gstr.Join(seat, ",")
	//修改完成后返回数据库
	_, err := g.Model("movie").Data("seat=?", StrSeat).Where("movieId=?", movieId).Update()
	if err != nil {
		return false
	}
	return true
}
