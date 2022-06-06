package social_medias

import (
	"errors"
	"fmt"
	"github.com/avtara-kw/social-media-api/businesses"
	"github.com/avtara-kw/social-media-api/businesses/social_medias"
	"github.com/avtara-kw/social-media-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	socialMediasService social_medias.Service
}

func NewSocialMediasController(uc social_medias.Service) *Controller {
	return &Controller{
		socialMediasService: uc,
	}
}

func (ctrl *Controller) Post(ctx *gin.Context) {
	var req RequestSocialMediasPostAndUpdate
	ID := ctx.MustGet("id").(float64)
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.socialMediasService.Post(req.ToDomain(int(ID)))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				utils.BuildErrorResponse("Internal Server Error",
					err, utils.EmptyObj{}))

		} else {
			ctx.JSON(http.StatusCreated,
				utils.BuildResponse("Successfully post social medias!",
					ResponseSocialMediasPostFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *Controller) GetAll(ctx *gin.Context) {
	ID := fmt.Sprintf("%.0f", ctx.MustGet("id").(float64))
	res, err := ctrl.socialMediasService.GetAll(ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			utils.BuildErrorResponse("Internal server error!",
				err, utils.EmptyObj{}))
	} else {
		ctx.JSON(http.StatusOK,
			utils.BuildResponse("Successfully get all social medias!",
				ResponseSocialMediasFromDomain(res)))
	}
}

func (ctrl *Controller) Update(ctx *gin.Context) {
	var req RequestSocialMediasPostAndUpdate
	ID := ctx.MustGet("id").(float64)
	photoID := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.socialMediasService.Update(req.ToDomain(int(ID)), photoID)
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
				utils.BuildResponse("Successfully updated an social media!",
					ResponseSocialMediaUpdateFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *Controller) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("id")
	userID := ctx.MustGet("id").(float64)
	intVar, err := strconv.Atoi(socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.BuildErrorResponse("Internal Server Error!",
				err, utils.EmptyObj{}))
	}
	err = ctrl.socialMediasService.Delete(intVar, int(userID))
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
			utils.BuildResponse("Your photo has been successfully deleted", map[string]string{"id": socialMediaID}))
	}
}
