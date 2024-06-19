package handler

import (
    "net/http"

    "github.com/danilocassiano/GoOportunities/schemas"
    "github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

//	@Summary		Create opening
//	@Description	Create a new job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateOpeningRequest	true	"Request body"
//	@Success		200		{object}	CreateOpeningResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
    request := CreateOpeningRequest{}

    if err := ctx.BindJSON(&request); err != nil {
        logger.Errorf("bind json error: %v", err.Error())
        sendError(ctx, http.StatusBadRequest, "invalid request format")
        return
    }

    if err := request.Validate(); err != nil {
        logger.Errorf("validate error: %v", err.Error())
        sendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    opening := schemas.Opening{
        Role:     request.Role,
        Company:  request.Company,
        Location: request.Location,
        Link:     request.Link,
        Remote:   *request.Remote,
        Salary:   request.Salary,
    }

    if err := db.Create(&opening).Error; err != nil {
        logger.Errorf("error creating opening: %v", err.Error())
        sendError(ctx, http.StatusInternalServerError, "error creating opening in database")
        return
    }

    sendSuccess(ctx, "create-opening", opening)
}
