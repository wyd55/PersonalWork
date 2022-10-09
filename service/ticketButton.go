package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"time"
)

//在我的订单中是否有0：不可取不可退，1：取票键，2：取票退票键
func GetTicketButton(movieId int, purchaseId int) int {
	film, _ := g.Model("movie").Fields("date,startTime,endTime").Where("movieId=?", movieId).One()
	str := gstr.SubStr(gconv.String(film["startTime"]), 11, 19)
	str1 := gconv.String(film["date"])
	str2 := str1 + " " + str
	str3 := gstr.SubStr(gconv.String(film["endTime"]), 11, 19)
	str4 := str1 + " " + str3
	//判断是否已经取票,1为待取票，0为已取票
	film1, _ := g.Model("purchase").Fields("checkIn").Where("purchaseId=?", purchaseId).One()
	if gconv.Int(film1["checkIn"]) == 1 {
		//如果30分钟之后还没开始那就是3
		if time.Now().Add(time.Minute * 30).Before(gconv.Time(str2)) {
			return 2
		} else {
			//如果在30分钟之后就开始的前提下，在结束时间之前就是1
			if time.Now().Add(time.Minute * 30).Before(gconv.Time(str4)) {
				return 1
			} else {
				//超时未领取
				g.Model("purchase").Data("checkIn=?", -1).Where("purchaseId=?", purchaseId).Update()
				return 0
			}
		}
	}
	return -1
}
