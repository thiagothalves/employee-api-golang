package main

import "github.com/thiagothalves/employee-api-golang/server"

func main() {

	server := server.NewServer()

	server.Run()

}
