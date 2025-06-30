package user

import (
	"time"
)

type (
	UserProfile struct {
		Id        string    `json:"_id"`
		Email     string    `json:"email"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	UserClaims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}

	CreateUserReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
		Name     string `json:"name" form:"name" validate:"required,max=64"`
	}

	UpdateUserReq struct {
		Name  string `json:"name,omitempty" form:"name"`
		Email string `json:"email,omitempty" form:"email" validate:"omitempty,email,max=255"`
	}
)
