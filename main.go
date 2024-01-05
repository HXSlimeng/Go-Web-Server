package main

import (
	"flag"

	"github.com/HXSlimeng/Go-Web-Server/server"
	"github.com/HXSlimeng/Go-Web-Server/tools"
)

func main() {
	geStr := flag.String("ge","","generate a template")

	flag.Parse()

	if *geStr != ""{
		tools.GeTemplate(*geStr)
	}else{
		server.StartServer()
	}

}
