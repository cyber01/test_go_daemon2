package controllers

import (
	"fmt"
	"github.com/dabfleming/go-martini-example/server"
	"github.com/go-martini/martini"
)

func init() {
	fmt.Println("controllers/user.go init()")
	server.Server.Group("/user", userRouter)
}

func userRouter(r martini.Router) {
	r.Get("/one", one)
	r.Get("/two", two)
}

func one() string {
	return "one"
}

func two() string {
	return "two"
}
