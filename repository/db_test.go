package repository

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestTestWhere(t *testing.T) {
	film, _ := g.Model("collect").Where("filmId=? and userId=?", 1, 15).One()
	fmt.Println(film)
}
