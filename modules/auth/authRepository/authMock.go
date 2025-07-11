package authrepository

import (
	"context"
	"user-manages/config"
	"user-manages/modules/auth"
	userPb "user-manages/modules/user/userPb"
	"user-manages/pkg/jwtauth"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepositoryMock struct {
	mock.Mock
}

// DeleteOneUserCredential implements AuthRepositoryService.
func (m *AuthRepositoryMock) DeleteOneUserCredential(pctx context.Context, credentialId string) (int64, error) {
	panic("unimplemented")
}

// FindOneUserCredential implements AuthRepositoryService.
func (m *AuthRepositoryMock) FindOneUserCredential(pctx context.Context, credentialId string) (*auth.Credential, error) {
	args := m.Called(pctx, credentialId)
	return args.Get(0).(*auth.Credential), args.Error(1)
}

// FindOneUserProfileToRefresh implements AuthRepositoryService.
func (m *AuthRepositoryMock) FindOneUserProfileToRefresh(pctx context.Context, grpcUrl string, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	panic("unimplemented")
}

// InsertOneUserCredential implements AuthRepositoryService.
func (m *AuthRepositoryMock) InsertOneUserCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	args := m.Called(pctx, req)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

// UpdateOneUserCredential implements AuthRepositoryService.
func (m *AuthRepositoryMock) UpdateOneUserCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error {
	panic("unimplemented")
}

func (m *AuthRepositoryMock) CredentialSearch(pctx context.Context, grpcUrl string, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	args := m.Called(pctx, grpcUrl, req)
	return args.Get(0).(*userPb.UserProfile), args.Error(1)
}

func (m *AuthRepositoryMock) AccessToken(cfg *config.Config, claims *jwtauth.Claims) string {
	args := m.Called(cfg, claims)
	return args.String(0)
}

func (m *AuthRepositoryMock) RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string {
	args := m.Called(cfg, claims)
	return args.String(0)
}

func (m *AuthRepositoryMock) FindOneAccessToken(pctx context.Context, accessToken string) (*auth.Credential, error) {
	return nil, nil
}
