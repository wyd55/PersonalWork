package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

//(xx,xx,xx,xx)的字符串拼接
func AlterString(film2 gdb.Result) map[string]string {
	zhumap := make(map[string]string, 2)
	var str1 string
	var str2 string
	for _, value := range film2 {
		if gconv.String(value["roleName"]) == "导演" {
			if str1 == "" {
				str1 = gconv.String(value["actorName"])
			} else {
				str1 = str1 + "," + gconv.String(value["actorName"])
			}
		} else {
			if str2 == "" {
				str2 = gconv.String(value["roleName"])
			} else {
				str2 = str2 + "," + gconv.String(value["roleName"])
			}
		}
	}
	zhumap["daoyan"] = str1
	zhumap["zhuyan"] = str2
	return zhumap
}

//拼接影院首页图片
func AlterStringPic(film gdb.Result) []string {
	var shuzu []string
	shuzu = append(shuzu, "http://r8xk3nk6q.hn-bkt.clouddn.com/image/film/null.png")
	for _, value := range film {
		str := "http://r8xk3nk6q.hn-bkt.clouddn.com/image/film/" + gconv.String(value["filmPic"])
		shuzu = append(shuzu, str)
	}
	shuzu = append(shuzu, "http://r8xk3nk6q.hn-bkt.clouddn.com/image/film/null.png")
	return shuzu
}
