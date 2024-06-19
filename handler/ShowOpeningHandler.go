package handler

import (
    "net/http"

    "github.com/danilocassiano/GoOportunities/schemas"
    "github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show a job opening
// @Description Show a job opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
        return
    }

    opening := schemas.Opening{}

    if err := db.First(&opening, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, "opening not found")
        return
    }

    sendSuccess(ctx, "show-success", opening)
}
