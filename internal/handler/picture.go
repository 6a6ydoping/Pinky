package handler

import (
	"github.com/6a6ydoping/Pinky/api"
	"github.com/6a6ydoping/Pinky/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) getPicturesByRange(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")

	uintPageValue, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		log.Printf("Error convert any to string in getPictures handler\n")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "Type error",
		})
		return
	}

	perPage := ctx.DefaultQuery("perPage", "10")
	uintPerPageValue, err := strconv.ParseUint(perPage, 10, 64)
	if err != nil {
		log.Printf("Error convert any to string in getPictures handler\n")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "Type error",
		})
		return
	}

	pictures, err := h.service.GetPictureByPage(ctx, uint(uintPageValue), uint(uintPerPageValue))
	if err != nil {
		log.Printf("Error get picture from database\n")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, pictures)
}
