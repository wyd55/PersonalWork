package api

import "github.com/gogf/gf/net/ghttp"

//轮播图接口
var BannerApi = new(banner)

type banner struct{}

func (b *banner) GetBannerPic(r *ghttp.Request) {
	Pic := []string{
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/5.jpg",
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/2.jpg",
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/1.jpg",
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/5.jpg",
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/3.jpg",
		"http://r8xk3nk6q.hn-bkt.clouddn.com/image/lunbo/4.jpg",
	}
	r.Response.Write(Pic)
}
