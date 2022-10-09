package route

import (
	"github.com/gogf/gf/net/ghttp"
	"golangProject/cinemaBuy/server/http/api"
)

func RegisterRouter(s *ghttp.Server) {

	//轮播
	BannerApi := api.BannerApi
	//首页热映/预售
	IndexFilmApi := api.IndexFilmApi
	//电影介绍
	FilmInfoApi := api.FilmInfoApi
	//更多电影
	MoreFilmApi := api.MoreFilmApi
	//购票信息
	CinemaBuyApi := api.CinemaApi
	//影院首页信息
	CinemaIndexApi := api.CinemaIndexApi
	//选座信息
	SelectSeatApi := api.SelectSeatApi
	//影院信息
	CinemaApi := api.CinemaApi
	//修改seat
	SeatOkApi := api.SeatOkApi
	//获取用户openid
	GetOpenIdApi := api.GetOpenIdApi
	//用户订单接口
	PurchaseApi := api.PurchaseApi
	//收藏、看过接口
	CollectSeenApi := api.CollectSeenApi

	//轮播图url地址
	s.Group("/", func(group *ghttp.RouterGroup) {
		//轮播图url地址
		group.ALL("/banner", BannerApi.GetBannerPic)
		//热映
		group.ALL("/reyin", IndexFilmApi.GetReYin)
		//预售
		group.ALL("/yushou", IndexFilmApi.GetYuShou)
		//根据电影介绍
		group.ALL("/filminfo", FilmInfoApi.GetFilmInfo)
		//演员介绍
		group.ALL("/actor", FilmInfoApi.GetActorInfo)
		//更多电影
		group.ALL("/morefilm", MoreFilmApi.MoreFilm)
		group.ALL("/morefilm2", MoreFilmApi.MoreFilm2)
		group.ALL("/fivedate", MoreFilmApi.GetFiveDate)
		//影院信息
		group.ALL("/cinema", CinemaApi.GetCinemaInfo)
		//购票信息
		group.ALL("/filmofcinema", CinemaBuyApi.GetCinemaByFilm)
		//影院首页信息
		group.ALL("/cinemaindexInfo", CinemaIndexApi.GetCinemaIndexInfo)
		group.ALL("/cinemaindexPic", CinemaIndexApi.GetCinemaIndexPic)
		group.ALL("/cinemaindexMovie", CinemaIndexApi.GetCinemaIndexMovie)
		//选座信息
		group.ALL("/selectseat", SelectSeatApi.GetSelectSeatInfo)
		//选座成功页
		group.ALL("/selectok", SelectSeatApi.GetSelectOkInfo)
		//前端传来成功选座位的数据和添加用户订单
		group.ALL("/changeseat", SeatOkApi.ChangeSeat)
		//code
		group.ALL("/getcode", GetOpenIdApi.GetUserOpenId)
		//我的订单
		group.ALL("/myrecord", PurchaseApi.GetRecordList)
		//确认订单
		group.ALL("/recordok", PurchaseApi.GetRecordOkInfo)
		//取票
		group.ALL("/qupiao", PurchaseApi.QuPiao)
		//退票
		group.ALL("/tuipiao", PurchaseApi.TuiPiao)
		//收藏
		group.ALL("/shoucang", CollectSeenApi.InsertCollect)
		//看过
		group.ALL("/kanguo", CollectSeenApi.InsertSeen)
		//收藏数据
		group.ALL("/collectList", CollectSeenApi.CollectList)
		//看过数据
		group.ALL("/seenList", CollectSeenApi.SeenList)
		//是否收藏过
		group.ALL("/Iscollect", CollectSeenApi.IsCollect)
		//是否收藏过
		group.ALL("/Isseen", CollectSeenApi.IsSeen)
	})

	//预售
}
