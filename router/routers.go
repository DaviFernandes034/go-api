package router

import (
	"net/http"

	"github.com/DaviFernandes034/go-api/handlers"
	"github.com/gin-gonic/gin"
)


func InitRouters(router *gin.Engine){

  v1:= router.Group("/api/v1")
  {

	v1.GET("/opening", handlers.ShowOpeningHandler)
	v1.POST("/opening",handlers.CreateOpeningHandler)
	v1.DELETE("/opening",handlers.DeleteOpeningHandler)
	v1.PUT("/opening",handlers.UpdateOpeningHandler)
	v1.GET("/openings",handlers.ListOpeningHandler)
  }


}