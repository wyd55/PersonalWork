package response

var Fanhui = new(fanhui)
var CinemaInfo = new(PiontCinemaInfo)
var ActorInfo = new(actorInfo)
var MovieInfoPoint = new(MovieInfo)
var SeatOkInfoPoint = new(SeatOkInfo)

type fanhui struct {
	Str string
}

//电影介绍信息
type PiontCinemaInfo struct {
	FilmId        int
	FilmName      string
	FilmInfo      string
	FilmTime      string
	FilmLocation  string
	FilmIntroduce string
	FilmPic       string
	//是否有取票或退票键（1可取票、2可退票、3可退可取）
	Button     int
	PurchaseId int
}

//电影演员信息
type actorInfo struct {
	ActorName string
	ActorPic  string
	RoleName  string
}

//选座信息
type MovieInfo struct {
	MovieId    int
	FilmId     int
	CinemaId   int
	Date       string
	StartTime  string
	EndTime    string
	Room       string
	Seat       string
	FilmPrice  float32
	FilmName   string
	ChangeDate string
	Seatings   []string
}

//选座成功页面
type SeatOkInfo struct {
	CinemaName     string
	CinemaLocation string
	FilmPrice      int
	Date           string
	Room           string
}

//返回userId
type UserId struct {
	UserId int
}

//返回收藏或看过数据
type CollectSeen struct {
	FilmId   int
	FilmPic  string
	FilmName string
	FilmInfo string
	DaoYan   string
	Actor    string
}
