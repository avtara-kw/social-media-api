package photos

import (
	"errors"
	"github.com/avtara-kw/social-media-api/businesses"
	"github.com/avtara-kw/social-media-api/businesses/photos"
	"github.com/avtara-kw/social-media-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	photoService photos.Service
}

func NewPhotosController(uc photos.Service) *Controller {
	return &Controller{
		photoService: uc,
	}
}

func (ctrl *Controller) Post(ctx *gin.Context) {
	var req RequestPhotosPostAndUpdate
	ID := ctx.MustGet("id").(float64)
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.photoService.Post(req.ToDomain(int(ID)))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				utils.BuildErrorResponse("Internal Server Error",
					err, utils.EmptyObj{}))

		} else {
			ctx.JSON(http.StatusCreated,
				utils.BuildResponse("Successfully upload photo!",
					ResponsePhotoPostFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *Controller) GetAll(ctx *gin.Context) {
	res, err := ctrl.photoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			utils.BuildErrorResponse("Internal server error!",
				err, utils.EmptyObj{}))
	} else {
		ctx.JSON(http.StatusOK,
			utils.BuildResponse("Successfully get all photos!",
				ResponsePhotosFromDomain(res)))
	}
}

func (ctrl *Controller) Delete(ctx *gin.Context) {
	photoID := ctx.Param("id")
	userID := ctx.MustGet("id").(float64)
	intVar, err := strconv.Atoi(photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.BuildErrorResponse("Internal Server Error!",
				err, utils.EmptyObj{}))
	}
	err = ctrl.photoService.Delete(intVar, int(userID))
	if err != nil {
		if errors.Is(err, businesses.ErrInternalServer) {
			ctx.JSON(http.StatusInternalServerError,
				utils.BuildErrorResponse("Internal Server Error!",
					err, utils.EmptyObj{}))
		} else if errors.Is(err, businesses.ErrForbiddenAccess) {
			ctx.JSON(http.StatusForbidden,
				utils.BuildErrorResponse("Forbidden Access!",
					err, utils.EmptyObj{}))
		}
	} else {
		ctx.JSON(http.StatusOK,
			utils.BuildResponse("Your photo has been successfully deleted", map[string]string{"id": photoID}))
	}
}

func (ctrl *Controller) Update(ctx *gin.Context) {
	var req RequestPhotosPostAndUpdate
	ID := ctx.MustGet("id").(float64)
	photoID := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.photoService.Update(req.ToDomain(int(ID)), photoID)
		if err != nil {
			if errors.Is(err, businesses.ErrForbiddenAccess) {
				ctx.JSON(http.StatusForbidden,
					utils.BuildErrorResponse("Forbidden Access!",
						err, utils.EmptyObj{}))
			} else {
				ctx.JSON(http.StatusInternalServerError,
					utils.BuildErrorResponse("Internal Server Error",
						err, utils.EmptyObj{}))
			}
		} else {
			ctx.JSON(http.StatusOK,
				utils.BuildResponse("Successfully updated an photos!",
					ResponsePhotoUpdateFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}
