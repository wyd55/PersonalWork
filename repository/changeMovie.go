package repository

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

//当成功下单后修改座位
func ChangeMovieSeat(seatok []string, movieId int) {
	film, _ := g.Model("movie").Fields("seat").Where("movieId=?", movieId).One()
	seatArray := gstr.Split(gconv.String(film["seat"]), ",")
	for _, value := range seatok {
		seatArray[gconv.Int(value)] = "3"
	}
	str := gstr.Join(seatArray, ",")
	g.Model("movie").Data("seat=?", str).Where("movieId=?", movieId).Update()
}
