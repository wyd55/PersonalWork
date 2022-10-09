package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
	"golangProject/cinemaBuy/vo/request"
)

//收藏、看过接口
var CollectSeenApi = new(collectseen)

type collectseen struct{}

//添加收藏
func (b *collectseen) InsertCollect(r *ghttp.Request) {
	var req *request.ReqCollectSeen
	r.Parse(&req)
	result := repository.InsertCollect(req.FilmId, req.UserId)
	if result {
		r.Response.Write("true")
	} else {
		r.Response.Write("false")
	}

}

//添加看过
func (b *collectseen) InsertSeen(r *ghttp.Request) {
	var req *request.ReqCollectSeen
	r.Parse(&req)
	result := repository.InsertSeen(req.FilmId, req.UserId)
	if result {
		r.Response.Write("true")
	} else {
		r.Response.Write("false")
	}
}

//收藏数据
func (b *collectseen) CollectList(r *ghttp.Request) {
	var req *request.UserId
	r.Parse(&req)
	result := repository.CollectData(req.UserId)
	r.Response.Write(result)
}

//看过数据
func (b *collectseen) SeenList(r *ghttp.Request) {
	var req *request.UserId
	r.Parse(&req)
	result := repository.SeenData(req.UserId)
	r.Response.Write(result)
}

//是否收藏过
func (b *collectseen) IsCollect(r *ghttp.Request) {
	var req *request.ReqCollectSeen
	r.Parse(&req)
	r.Response.Write(repository.GetIsCollect(req.FilmId, req.UserId))
}

//是否看过
func (b *collectseen) IsSeen(r *ghttp.Request) {
	var req *request.ReqCollectSeen
	r.Parse(&req)
	r.Response.Write(repository.GetIsSeen(req.FilmId, req.UserId))
}
