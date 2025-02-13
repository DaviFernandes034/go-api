package main

import (
	

	"github.com/DaviFernandes034/go-api/config"
	"github.com/DaviFernandes034/go-api/router"
)

var (

	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	//iniciando as configurações
	
	err:= config.Init()
	if  err != nil {
		logger.Errorf("config init error: %v", err )
		return
	}

	//iniciando as rotas
	router.Init()



}
