package repository

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/vo/response"
)

//添加收藏
func InsertCollect(filmId int, userId int) bool {
	//先查找对应的filmId和userId是否存在
	film, _ := g.Model("collect").Where("filmId=? and userId=?", filmId, userId).One()
	if film == nil {
		//添加收藏
		g.Model("collect").Data(g.Map{"filmId": filmId, "userId": userId}).Insert()
		return true
	} else {
		//取消收藏
		g.Model("collect").Where("collectId=?", film["collectId"]).Delete()
		return false
	}
	return true
}

//添加看过
func InsertSeen(filmId int, userId int) bool {
	//先查找对应的filmId和userId是否存在
	film, _ := g.Model("seen").Where("filmId=? and userId=?", filmId, userId).One()
	if film == nil {
		//添加收藏
		g.Model("seen").Data(g.Map{"filmId": filmId, "userId": userId}).Insert()
		return true
	} else {
		//取消收藏
		g.Model("seen").Where("seenId=?", film["seenId"]).Delete()
		return false
	}
	return true
}

//收藏数据
func CollectData(userId int) []response.CollectSeen {
	var res []response.CollectSeen
	film, _ := g.Model("collect").Fields("filmId").Where("userId=?", userId).All()
	for _, value := range film {
		actor := ""
		daoYan := ""
		film1, _ := g.Model("film").Fields("filmName,filmInfo,filmPic").Where("filmId=?", value["filmId"]).One()
		film2, _ := g.Model("filmactor").InnerJoin("actor", "filmactor.actorId=actor.actorId").Fields("filmactor.roleName,actor.actorName").Where("filmactor.filmId=?", value["filmId"]).All()
		for _, value2 := range film2 {
			if gconv.String(value2["roleName"]) == "导演" {
				if daoYan == "" {
					daoYan = gconv.String(value2["actorName"])
				} else {
					daoYan = daoYan + "," + gconv.String(value2["actorName"])
				}
			} else {
				if actor == "" {
					actor = gconv.String(value2["actorName"])
				} else {
					actor = actor + "," + gconv.String(value2["actorName"])
				}
			}
		}
		var indata = new(response.CollectSeen)
		indata.FilmId = gconv.Int(value["filmId"])
		indata.FilmName = gconv.String(film1["filmName"])
		indata.FilmInfo = gconv.String(film1["filmInfo"])
		indata.DaoYan = daoYan
		indata.Actor = actor
		indata.FilmPic = gconv.String(film1["filmPic"])
		res = append(res, *indata)

	}
	return res
}

//看过数据
func SeenData(userId int) []response.CollectSeen {
	var res []response.CollectSeen
	film, _ := g.Model("seen").Fields("filmId").Where("userId=?", userId).All()
	for _, value := range film {
		actor := ""
		daoYan := ""
		film1, _ := g.Model("film").Fields("filmName,filmInfo,filmPic").Where("filmId=?", value["filmId"]).One()
		film2, _ := g.Model("filmactor").InnerJoin("actor", "filmactor.actorId=actor.actorId").Fields("filmactor.roleName,actor.actorName").Where("filmactor.filmId=?", value["filmId"]).All()
		for _, value2 := range film2 {
			if gconv.String(value2["roleName"]) == "导演" {
				if daoYan == "" {
					daoYan = gconv.String(value2["actorName"])
				} else {
					daoYan = daoYan + "," + gconv.String(value2["actorName"])
				}
			} else {
				if actor == "" {
					actor = gconv.String(value2["actorName"])
				} else {
					actor = actor + "," + gconv.String(value2["actorName"])
				}
			}
		}
		var indata = new(response.CollectSeen)
		indata.FilmId = gconv.Int(value["filmId"])
		indata.FilmName = gconv.String(film1["filmName"])
		indata.FilmInfo = gconv.String(film1["filmInfo"])
		indata.DaoYan = daoYan
		indata.Actor = actor
		indata.FilmPic = gconv.String(film1["filmPic"])
		res = append(res, *indata)
	}
	return res
}

//是否收藏过
func GetIsCollect(filmId int, userId int) bool {
	//先查找对应的filmId和userId是否存在
	film, _ := g.Model("collect").Where("filmId=? and userId=?", filmId, userId).One()
	if film == nil {
		return true
	} else {
		return false
	}
	return true
}

//是否看过
func GetIsSeen(filmId int, userId int) bool {
	//先查找对应的filmId和userId是否存在
	film, _ := g.Model("seen").Where("filmId=? and userId=?", filmId, userId).One()
	if film == nil {
		return true
	} else {
		return false
	}
	return true
}
