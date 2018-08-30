package main

import (
	"fmt"
	_ "github.com/cyber01/test_go_daemon2/controllers"
	"github.com/cyber01/test_go_daemon2/server"
)

func main() {
	fmt.Println("Server Started")
	server.Server.Run()
}
