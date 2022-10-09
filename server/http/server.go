package http

import (
	"github.com/gogf/gf/frame/g"
	"golangProject/cinemaBuy/server/http/route"
)

func Start() {
	s := g.Server()
	route.RegisterRouter(s)
	s.SetPort(8100)
	s.Run()
}
