package main

import (
	"fmt"
	_ "github.com/dabfleming/go-martini-example/controllers"
	"github.com/dabfleming/go-martini-example/server"
)

func main() {
	fmt.Println("Server Started")
	server.Server.Run()
}
