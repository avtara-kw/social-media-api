package users

import (
	"errors"
	"fmt"
	"github.com/avtara-kw/social-media-api/businesses"
	"github.com/avtara-kw/social-media-api/businesses/users"
	"github.com/avtara-kw/social-media-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService users.Service
}

func NewUserController(uc users.Service) *UserController {
	return &UserController{
		userService: uc,
	}
}

func (ctrl *UserController) Registration(ctx *gin.Context) {
	var req RequestUserRegistration
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.userService.Registration(req.UserRegistrationToDomain())
		if err != nil {
			if errors.Is(err, businesses.ErrEmailAccountDuplicate) || errors.Is(err, businesses.ErrEmailAccountDuplicate) {
				ctx.JSON(http.StatusConflict,
					utils.BuildErrorResponse("Data Duplicate",
						err, utils.EmptyObj{}))
			} else {
				ctx.JSON(http.StatusInternalServerError,
					utils.BuildErrorResponse("Internal Server Error",
						err, utils.EmptyObj{}))
			}
		} else {
			ctx.JSON(http.StatusCreated,
				utils.BuildResponse("Successfully created an account, please check your email to activate!",
					FromDomain(res)))
		}
		fmt.Println(res, err)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}

}
