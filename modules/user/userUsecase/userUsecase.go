package userusecase

import (
	"context"
	"errors"
	"log"
	"time"
	"user-manages/modules/user"
	userPb "user-manages/modules/user/userPb"
	userRepository "user-manages/modules/user/userRepository"
	"user-manages/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseService interface {
		CreateUser(pctx context.Context, req *user.CreateUserReq) (*user.UserProfile, error)
		FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error)
		FindOneUserCredential(pctx context.Context, password, email string) (*userPb.UserProfile, error)
		FindOneUserProfileToRefresh(pctx context.Context, userId string) (*userPb.UserProfile, error)
	}

	userUsecase struct {
		userRepository userRepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userRepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) CreateUser(pctx context.Context, req *user.CreateUserReq) (*user.UserProfile, error) {
	if !u.userRepository.IsUniqueUser(pctx, req.Email, req.Name) {
		return nil, errors.New("error: email or username already exist")
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: failed to hash password")
	}

	// Insert one user
	userId, err := u.userRepository.InsertOneUser(pctx, &user.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Name:      req.Name,
		CreatedAt: utils.LocalTime(),
	})

	return u.FindOneUserProfile(pctx, userId.Hex())
}

func (u *userUsecase) FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error) {
	result, err := u.userRepository.FindOneUserProfile(pctx, userId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &user.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt.In(loc),
	}, nil
}

func (u *userUsecase) FindOneUserCredential(pctx context.Context, password, email string) (*userPb.UserProfile, error) {
	result, err := u.userRepository.FindOneUserCredential(pctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		log.Printf("Error: FindOneUserCredential: %s", err.Error())
		return nil, errors.New("error: password is invalid")
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &userPb.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt.In(loc).String(),
	}, nil
}

func (u *userUsecase) FindOneUserProfileToRefresh(pctx context.Context, userId string) (*userPb.UserProfile, error) {
	result, err := u.userRepository.FindOneUserProfileToRefresh(pctx, userId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &userPb.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt.In(loc).String(),
	}, nil
}
