package api

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func Test_Test(t *testing.T) {
	film, _ := g.Model("purchase").InnerJoin("movie", "purchase.movieId=movie.movieId").Fields("movie.filmId,movie.movieId").Where("purchase.userId=?", 15).All()
	fmt.Println(film)
}
