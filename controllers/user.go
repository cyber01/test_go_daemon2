package controllers

import (
	"fmt"
	"github.com/cyber01/test_go_daemon2/server"
	"github.com/go-martini/martini"
	"encoding/json"
    "log"
    "os/exec"
    "net/http"
)

func init() {
	fmt.Println("controllers/server.go init()")
	server.Server.Group("/server", serverRouter)
}

type test_struct struct {
    Test string
    Bla string
}



func serverRouter(r martini.Router) {
	r.Post("/tests", ttttt)
	r.Get("/test/:check", func(params martini.Params) string {
    return params["check"]
  })
}


func ttttt(req *http.Request) string {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    log.Println(t.Test)
    return t.Test+t.Bla

}

func one() string {
	return "one"
}

func two() string {
	return "two"

}