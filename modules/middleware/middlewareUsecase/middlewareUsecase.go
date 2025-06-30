package middlewareusecase

import (
	"errors"
	"log"
	"user-manages/config"
	middlewareRepository "user-manages/modules/middleware/middlewareRepository"
	"user-manages/pkg/jwtauth"

	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error)
		UserIdParamValidation(c echo.Context) (echo.Context, error)
	}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}

func (u *middlewareUsecase) JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error) {
	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return nil, err
	}

	if err := u.middlewareRepository.AccessTokenSearch(ctx, cfg.Grpc.AuthUrl, accessToken); err != nil {
		return nil, err
	}

	c.Set("user_id", claims.UserId)

	return c, nil
}

func (u *middlewareUsecase) UserIdParamValidation(c echo.Context) (echo.Context, error) {
	userIdReq := c.Param("user_id")
	userIdToken := c.Get("user_id").(string)

	if userIdToken == "" {
		log.Printf("Error: user_id not found")
		return nil, errors.New("error: user_id is required")
	}

	if userIdToken != userIdReq {
		log.Printf("Error: user_id not match, user_id_req: %s, user_id_token: %s", userIdReq, userIdToken)
		return nil, errors.New("error: user_id not match")
	}

	return c, nil
}
