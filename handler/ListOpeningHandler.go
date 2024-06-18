package handler

import (
	"net/http"

	"github.com/danilocassiano/GoOportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context){
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error liting openings")
		return
	}

	sendSuccess(ctx, "Llinting-openings", openings)
}