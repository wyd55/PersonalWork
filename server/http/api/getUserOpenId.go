package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/service"
	"golangProject/cinemaBuy/vo/request"
	"golangProject/cinemaBuy/vo/response"
)

var GetOpenIdApi = new(openId)

type openId struct{}

//获取登录用户openid
func (o *openId) GetUserOpenId(r *ghttp.Request) {
	var req *request.GetCode
	r.Parse(&req)
	result := service.GetUserOpenIdByCode(req.Code)
	repository.InsertUserInfo(result.OpenId)
	//根据openId返回userId
	var userId = new(response.UserId)
	userId.UserId = repository.GetUserIdByOpenId(result.OpenId)
	r.Response.Writeln(&userId)
}
