package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
	"user-manages/config"
	"user-manages/modules/auth"
	authrepository "user-manages/modules/auth/authRepository"
	authusecase "user-manages/modules/auth/authUsecase"
	"user-manages/modules/user"
	userPb "user-manages/modules/user/userPb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	testLogin struct {
		ctx      context.Context
		cfg      *config.Config
		req      *auth.UserLoginReq
		expected *auth.ProfileIntercepter
		isErr    bool
	}
)

func TestLogin(t *testing.T) {
	repoMock := new(authrepository.AuthRepositoryMock)
	usecase := authusecase.NewAuthUsecase(repoMock)

	cfg := NewTestConfig()
	ctx := context.Background()

	credentialIdSuccess := primitive.NewObjectID()
	credentialIdFailed := primitive.NewObjectID()

	tests := []testLogin{
		{
			ctx: ctx,
			cfg: cfg,
			req: &auth.UserLoginReq{
				Email:    "success@sekai.com",
				Password: "123456",
			},
			expected: &auth.ProfileIntercepter{
				UserProfile: &user.UserProfile{
					Id:        "player:001",
					Email:     "success@sekai.com",
					Name:      "player001",
					CreatedAt: time.Time{},
				},
				Credential: &auth.CredentialRes{
					Id:           credentialIdSuccess.Hex(),
					UserId:       "player:001",
					AccessToken:  "xxx",
					RefreshToken: "xxx",
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
				},
			},
			isErr: false,
		},
		{
			ctx: ctx,
			cfg: cfg,
			req: &auth.UserLoginReq{
				Email:    "failed2@sekai.com",
				Password: "123456",
			},
			expected: nil,
			isErr:    true,
		},
		{
			ctx: ctx,
			cfg: cfg,
			req: &auth.UserLoginReq{
				Email:    "failed3@sekai.com",
				Password: "123456",
			},
			expected: nil,
			isErr:    true,
		},
	}

	// CredentialSearch
	repoMock.On("CredentialSearch", ctx, cfg.Grpc.UserUrl, &userPb.CredentialSearchReq{
		Email:    "success@sekai.com",
		Password: "123456",
	}).Return(&userPb.UserProfile{
		Id:        "001",
		Email:     "success@sekai.com",
		Name:      "player001",
		CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
	}, nil)

	repoMock.On("CredentialSearch", ctx, cfg.Grpc.UserUrl, &userPb.CredentialSearchReq{
		Email:    "failed2@sekai.com",
		Password: "123456",
	}).Return(&userPb.UserProfile{}, errors.New("error: email or password is invalid"))

	repoMock.On("CredentialSearch", ctx, cfg.Grpc.UserUrl, &userPb.CredentialSearchReq{
		Email:    "failed3@sekai.com",
		Password: "123456",
	}).Return(&userPb.UserProfile{
		Id:        "003",
		Email:     "failed3@sekai.com",
		Name:      "player003",
		CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
	}, nil)

	// Access Token
	repoMock.On("AccessToken", cfg, mock.AnythingOfType("*jwtauth.Claims")).Return("xxx")

	// Refresh Token
	repoMock.On("RefreshToken", cfg, mock.AnythingOfType("*jwtauth.Claims")).Return("xxx")

	// InsertOneUserCredential
	repoMock.On("InsertOneUserCredential", ctx, &auth.Credential{
		UserId:       "player:001",
		AccessToken:  "xxx",
		RefreshToken: "xxx",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}).Return(credentialIdSuccess, nil)

	repoMock.On("InsertOneUserCredential", ctx, &auth.Credential{
		UserId:       "player:003",
		AccessToken:  "xxx",
		RefreshToken: "xxx",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}).Return(credentialIdFailed, nil)

	// FindOneUserCredential
	repoMock.On("FindOneUserCredential", ctx, credentialIdSuccess.Hex()).Return(&auth.Credential{
		Id:           credentialIdSuccess,
		UserId:       "player:001",
		AccessToken:  "xxx",
		RefreshToken: "xxx",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}, nil)

	repoMock.On("FindOneUserCredential", ctx, credentialIdFailed.Hex()).Return(&auth.Credential{}, errors.New("error: user credential not found"))

	for i, test := range tests {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.Login(test.ctx, test.cfg, test.req)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			result.CreatedAt = time.Time{}
			result.Credential.CreatedAt = time.Time{}
			result.Credential.UpdatedAt = time.Time{}

			assert.Equal(t, test.expected, result)
		}
	}
}
