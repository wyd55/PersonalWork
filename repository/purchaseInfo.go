package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
)

//下单成功后添加用户订单
func InsertPurchase(userId int, movieId int, seatok []string, pai []string, zuo []string, ticketNum int) error {
	seatokStr := gstr.Join(seatok, ",")
	paiStr := gstr.Join(pai, ",")
	zuoStr := gstr.Join(zuo, ",")
	_, err := g.Model("purchase").Data(g.Map{"userId": userId, "movieId": movieId, "number": ticketNum, "seatok": seatokStr, "pai": paiStr, "zuo": zuoStr, "checkIn": 1}).Insert()
	if err != nil {
		return err
	}
	return nil
}

//根据userid返回全部用户订单电影id
func GetPurFilmIdByUserId(userId int) gdb.Result {
	film, _ := g.Model("purchase").InnerJoin("movie", "purchase.movieId=movie.movieId").Fields("movie.filmId,movie.movieId,purchase.purchaseId").Where("purchase.userId=?", userId).All()
	return film
}

//根据purchaseid返回movieid
func GetMovieIdByPurchaseId(purchaseId int) gdb.Record {
	film, _ := g.Model("purchase").Fields("movieId,seatok,pai,zuo,number").Where("purchaseId=?", purchaseId).One()
	return film
}

//取票修改checkIn为0并返回true
func ChangeCheckIn(purchaseId int) bool {
	_, err := g.Model("purchase").Data("checkIn=?", 0).Where("purchaseId=?", purchaseId).Update()
	if err != nil {
		return false
	}
	return true
}

//删除指定订单
func DeletePurchase(purchaseId int) bool {
	_, err := g.Model("purchase").Where("purchaseId=?", purchaseId).Delete()
	if err != nil {
		return false
	}
	return true
}
