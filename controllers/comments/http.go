package comments

import (
	"errors"
	"github.com/avtara-kw/social-media-api/businesses"
	"github.com/avtara-kw/social-media-api/businesses/comments"
	"github.com/avtara-kw/social-media-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	commentsService comments.Service
}

func NewCommentsController(uc comments.Service) *Controller {
	return &Controller{
		commentsService: uc,
	}
}

func (ctrl *Controller) Post(ctx *gin.Context) {
	var req RequestCommentPost
	ID := ctx.MustGet("id").(float64)
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.commentsService.Post(req.ToDomain(int(ID)))
		if err != nil {
			if errors.Is(err, businesses.ErrPhotoNotFound) {
				ctx.JSON(http.StatusNotFound,
					utils.BuildErrorResponse("Photo ID not found!",
						err, utils.EmptyObj{}))
			} else {
				ctx.JSON(http.StatusInternalServerError,
					utils.BuildErrorResponse("Internal Server Error",
						err, utils.EmptyObj{}))
			}
		} else {
			ctx.JSON(http.StatusCreated,
				utils.BuildResponse("Successfully upload photo!",
					ResponseCommentPostFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *Controller) GetAll(ctx *gin.Context) {
	res, err := ctrl.commentsService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			utils.BuildErrorResponse("Internal server error!",
				err, utils.EmptyObj{}))
	} else {
		ctx.JSON(http.StatusOK,
			utils.BuildResponse("Successfully get all comments!",
				ResponseCommentsFromDomain(res)))
	}
}

func (ctrl *Controller) Delete(ctx *gin.Context) {
	commentID := ctx.Param("id")
	userID := ctx.MustGet("id").(float64)
	intVar, err := strconv.Atoi(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.BuildErrorResponse("Internal Server Error!",
				err, utils.EmptyObj{}))
	}
	err = ctrl.commentsService.Delete(intVar, int(userID))
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
			utils.BuildResponse("Your comment has been successfully deleted", map[string]string{"id": commentID}))
	}
}

func (ctrl *Controller) Update(ctx *gin.Context) {
	var req RequestCommentUpdate
	ID := ctx.MustGet("id").(float64)
	commentID := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.commentsService.Update(req.ToDomain(int(ID)), commentID)
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
				utils.BuildResponse("Successfully updated an comment!",
					ResponseCommentUpdateFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}
