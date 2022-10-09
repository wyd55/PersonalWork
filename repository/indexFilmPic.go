package repository

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//热映
func GetReYinPic() gdb.Result {
	film, _ := g.Model("film").Fields("filmId, filmName, filmPic").Where("filmSort = ?", 1).All()
	return film
}

//预售
func GetYuShou() gdb.Result {
	film, _ := g.Model("film").Fields("filmId, filmName, filmPic").Where("filmSort = ?", 2).All()
	return film
}
