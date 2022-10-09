package service

import (
	"github.com/gogf/gf/text/gstr"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/vo/response"
	"time"
)

//修改查询到的movie表的信息，使其可以直接给前端使用(选座页面)
func ChangeMovieInfo(movieId int) *response.MovieInfo {
	req := repository.GetSelectSeatInfo(movieId)
	now := time.Now()
	//修改日期
	str := gstr.SubStr(req.Date, 5, 2)
	str1 := gstr.SubStr(req.Date, 8, 2)
	str2 := gstr.SubStr(req.StartTime, 11, 2)
	str3 := gstr.SubStr(req.StartTime, 14, 2)
	str4 := gstr.SubStr(req.EndTime, 11, 2)
	str5 := gstr.SubStr(req.EndTime, 14, 2)
	loc, _ := time.LoadLocation("Local")
	weekNum, _ := time.ParseInLocation("2006-01-02", req.Date, loc)
	weekHan := NumToHanzi(int(weekNum.Weekday()))
	if req.Date == now.Format("2006-01-02") {
		req.ChangeDate = "今天" + " " + str + "月" + str1 + "日" + " " + str2 + ":" + str3 + "-" + str4 + ":" + str5
	} else if req.Date == now.AddDate(0, 0, 1).Format("2006-01-02") {
		req.ChangeDate = "明天" + " " + str + "月" + str1 + "日" + " " + str2 + ":" + str3 + "-" + str4 + ":" + str5
	} else if req.Date == now.AddDate(0, 0, 2).Format("2006-01-02") {
		req.ChangeDate = "后天" + " " + str + "月" + str1 + "日" + " " + str2 + ":" + str3 + "-" + str4 + ":" + str5
	} else {
		req.ChangeDate = "周" + weekHan + " " + str + "月" + str1 + "日" + " " + str2 + ":" + str3 + "-" + str4 + ":" + str5
	}
	req.FilmName = repository.GetFilmNameByMovieId(movieId)
	req.Seatings = gstr.Split(req.Seat, ",")
	return req
}
