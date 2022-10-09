package service

import (
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/vo/response"
)

//修改获得的数据返回给选座成功页面使用
func ChangeMovieCinemaInfo(movieId int) *response.SeatOkInfo {
	film := repository.GetSelectOkInfo(movieId)
	var req = response.SeatOkInfoPoint
	req.CinemaName = gconv.String(film["Name"])
	req.CinemaLocation = gconv.String(film["Location"])
	req.FilmPrice = gconv.Int(film["filmPrice"])
	req.Room = gconv.String(film["room"])
	str1 := gstr.SubStr(gconv.String(film["date"]), 0, 4)
	str2 := gstr.SubStr(gconv.String(film["date"]), 5, 2)
	str3 := gstr.SubStr(gconv.String(film["date"]), 8, 2)
	str4 := gstr.SubStr(gconv.String(film["startTime"]), 11, 2)
	str5 := gstr.SubStr(gconv.String(film["startTime"]), 14, 2)
	req.Date = str1 + "年" + str2 + "月" + str3 + "日" + " " + str4 + ":" + str5
	return req
}
