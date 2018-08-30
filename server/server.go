package server

import (
	"fmt"
	"github.com/go-martini/martini"
)

var Server *martini.ClassicMartini

func init() {
	fmt.Println("server/server.go init()")
	Server = martini.Classic()
}
