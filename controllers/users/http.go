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
		res, err := ctrl.userService.Registration(req.ToDomain())
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
				utils.BuildResponse("Successfully created an account!",
					ResponseUserRegistrationFromDomain(res)))
		}
		fmt.Println(res, err)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *UserController) Login(ctx *gin.Context) {
	var req RequestUserLogin
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.userService.Login(req.ToDomain())
		if err != nil {
			if errors.Is(err, businesses.ErrInvalidCredential) {
				ctx.JSON(http.StatusForbidden,
					utils.BuildErrorResponse("Invalid Credential!",
						err, utils.EmptyObj{}))
			} else if errors.Is(err, businesses.ErrAccountNotFound) {
				ctx.JSON(http.StatusNotFound,
					utils.BuildErrorResponse("Account not found",
						err, utils.EmptyObj{}))
			}
		} else {
			ctx.JSON(http.StatusOK,
				utils.BuildResponse("Successfully login an account!",
					ResponseUserLoginFromDomain(res)))
		}
		fmt.Println(res, err)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}

func (ctrl *UserController) Delete(ctx *gin.Context) {
	ID := fmt.Sprintf("%.0f", ctx.MustGet("id").(float64))
	err := ctrl.userService.Delete(ID)
	if err != nil {
		if errors.Is(err, businesses.ErrInternalServer) {
			ctx.JSON(http.StatusInternalServerError,
				utils.BuildErrorResponse("Internal Server Error!",
					err, utils.EmptyObj{}))
		}
	} else {
		ctx.JSON(http.StatusOK,
			utils.BuildResponse("Your account has been successfully deleted", map[string]string{"id": ID}))
	}
}

func (ctrl *UserController) Update(ctx *gin.Context) {
	var req RequestUserUpdate
	ID := fmt.Sprintf("%.0f", ctx.MustGet("id").(float64))
	if err := ctx.ShouldBindJSON(&req); err == nil {
		res, err := ctrl.userService.Update(ID, req.ToDomain())
		if err != nil {
			if errors.Is(err, businesses.ErrEmailAccountDuplicate) || errors.Is(err, businesses.ErrEmailAccountDuplicate) {
				ctx.JSON(http.StatusConflict,
					utils.BuildErrorResponse("Data Duplicate",
						err, utils.EmptyObj{}))
			} else if errors.Is(err, businesses.ErrAccountNotFound) {
				ctx.JSON(http.StatusNotFound,
					utils.BuildErrorResponse("Account not found",
						err, utils.EmptyObj{}))
			} else {
				ctx.JSON(http.StatusInternalServerError,
					utils.BuildErrorResponse("Internal Server Error",
						err, utils.EmptyObj{}))
			}
		} else {
			ctx.JSON(http.StatusOK,
				utils.BuildResponse("Successfully updated an account!",
					ResponseUserUpdateFromDomain(res)))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
}
