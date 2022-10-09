package service

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"golangProject/cinemaBuy/vo/request"
)

func GetUserOpenIdByCode(code string) *request.WXLoginResp {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appId := "wxd9ddb5c2928dbd5a"
	secret := "8d62f74ff312364768c3565e6b207291"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, appId, secret, code)
	// Struct
	var req *request.WXLoginResp
	g.Client().GetVar(url).Scan(&req)
	return req
}
