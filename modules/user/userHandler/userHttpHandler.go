package userhandler

import (
	"context"
	"net/http"
	"strings"
	"user-manages/config"
	"user-manages/pkg/request"
	"user-manages/pkg/response"

	"user-manages/modules/user"
	userUsecase "user-manages/modules/user/userUsecase"

	"github.com/labstack/echo/v4"
)

type (
	UserHttpHandlerService interface {
		CreateUser(c echo.Context) error
		FindOneUserProfile(c echo.Context) error
		ListUsers(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{userUsecase: userUsecase}
}

func (h *userHttpHandler) CreateUser(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.CreateUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) FindOneUserProfile(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Param("user_id"), "user:")

	res, err := h.userUsecase.FindOneUserProfile(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *userHttpHandler) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.userUsecase.ListAllUsers(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, users)
}

func (h *userHttpHandler) UpdateUser(c echo.Context) error {
	userId := strings.TrimPrefix(c.Param("user_id"), "user:")

	req := new(user.UpdateUserReq)
	if err := c.Bind(req); err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.userUsecase.UpdateUser(c.Request().Context(), userId, req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, echo.Map{"message": "user updated successfully"})
}

func (h *userHttpHandler) DeleteUser(c echo.Context) error {
	userId := strings.TrimPrefix(c.Param("user_id"), "user:")

	err := h.userUsecase.DeleteUser(c.Request().Context(), userId)
	if err != nil {
		if err.Error() == "user not found" {
			return response.ErrResponse(c, http.StatusNotFound, err.Error())
		}
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, echo.Map{"message": "user deleted successfully"})
}
