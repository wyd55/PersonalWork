package repository

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

//每当用户登录后检查是否为新用户，新用户就给表添加用户数据
func InsertUserInfo(openId string) error {
	_, err := g.Model("user").Data(g.Map{"openId": openId}).Insert()
	if err != nil {
		return err
	}
	return nil
}

//通过openId返回userId
func GetUserIdByOpenId(openId string) int {
	result, _ := g.Model("user").Fields("userId").Where("openId=?", openId).One()
	return gconv.Int(result["userId"])
}
