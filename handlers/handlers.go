package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateOpeningHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":" get ok",
	} )

}


func ShowOpeningHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":" get ok",
	} )

}

func DeleteOpeningHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":" get ok",
	} )

}

func UpdateOpeningHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":" get ok",
	} )

}


func ListOpeningHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":" get ok",
	} )

}