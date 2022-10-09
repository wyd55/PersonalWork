package api

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/repository"
)

//轮播图接口
var IndexFilmApi = new(indexFilm)

type indexFilm struct{}

//热映
func (i *indexFilm) GetReYin(r *ghttp.Request) {
	reYinPicUrl := repository.GetReYinPic()
	r.Response.Write(reYinPicUrl)
}

//预售
func (i *indexFilm) GetYuShou(r *ghttp.Request) {
	yuShouPicUrl := repository.GetYuShou()
	r.Response.Write(yuShouPicUrl)
}
