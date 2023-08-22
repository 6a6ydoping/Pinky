package handler

import (
	"github.com/6a6ydoping/Pinky/api"
	"github.com/6a6ydoping/Pinky/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createPic(ctx *gin.Context) {
	var pic entity.Picture
	err := ctx.BindJSON(&pic)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	err = h.service.CreatePicture(ctx, &pic)
	ctx.JSON(http.StatusOK, "OK")
}
