package router

import "github.com/gin-gonic/gin"

func Init() {

	//iniciando as rotas
	router := gin.Default()

	InitRouters(router)

	//rodando o servidor
	router.Run()
}
