package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"msg":"GET Opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"Post Opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,
	gin.H{
		"msg":"DELETE Opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"PUT Opening",
	})
}

func ListOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"GET Openings",
	})
}