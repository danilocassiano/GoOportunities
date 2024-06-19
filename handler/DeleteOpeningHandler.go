package handler

import (
    "fmt"
    "net/http"

    "github.com/danilocassiano/GoOportunities/schemas"
    "github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

// DeleteOpeningHandler deletes a job opening by ID.
//	@Summary		Delete a job opening
//	@Description	Delete a job opening by ID
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Opening Identification"
//	@Success		200	{object}	DeleteOpeningResponse
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Router			/opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
        return
    }

    opening := schemas.Opening{}

    if err := db.First(&opening, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
        return
    }

    if err := db.Delete(&opening).Error; err != nil {
        sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
        return
    }

    sendSuccess(ctx, "deleting-opening", opening)
}
