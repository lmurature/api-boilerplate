package services

import (
	"context"
	"errors"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/users"
	"github.com/lmurature/api-boilerplate/cmd/api/utils/auth"
	"github.com/lmurature/api-boilerplate/cmd/api/utils/crypto"
	"github.com/lmurature/api-boilerplate/cmd/api/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

type UsersServiceInterface interface {
	GetUserByID(ctx context.Context, userID int64) (*users.User, apierrors.ApiError)
	RegisterUser(ctx context.Context, request users.RegisterUserRequest) (*users.User, apierrors.ApiError)
	AuthenticateUser(ctx context.Context, request users.AuthenticateUserRequest) (*users.AuthenticateUserResponse, apierrors.ApiError)
	RefreshUserToken(ctx context.Context, request users.RefreshUserTokenRequest) (*users.AuthenticateUserResponse, apierrors.ApiError)
}

type UsersService struct {
	usersDao users.UserDaoInterface
}

func NewUsersService(dao *users.UserDaoInterface) *UsersServiceInterface {
	var s UsersServiceInterface = &UsersService{
		usersDao: *dao,
	}
	return &s
}

func (s *UsersService) GetUserByID(ctx context.Context, userID int64) (*users.User, apierrors.ApiError) {
	return s.usersDao.GetUserDataByID(ctx, userID)
}

func (s *UsersService) RegisterUser(ctx context.Context, request users.RegisterUserRequest) (*users.User, apierrors.ApiError) {
	if request.UserData == nil {
		return nil, apierrors.NewBadRequestApiError("user data cannot be null")
	}

	hashedPassword, hashErr := crypto_utils.HashPassword(request.Password)
	request.Password = nil
	if hashErr != nil {
		if hashErr == bcrypt.ErrPasswordTooLong {
			return nil, apierrors.NewBadRequestApiError("password cannot be longer than 72 bytes")
		}
		return nil, apierrors.NewInternalServerApiError("Error handling password", errors.New("hash error"))
	}

	request.Password = hashedPassword
	result, err := s.usersDao.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth_utils.GenerateToken(result.UserID, auth_utils.TokenTypeRefresh)
	if err != nil {
		return nil, err
	}
	if err := s.usersDao.UpdateUserRefreshToken(ctx, result.UserID, *refreshToken); err != nil {
		logger_utils.Logger.Println("Successfully created refresh token for user ", result.UserID)
		return nil, err
	}

	return result, nil
}

func (s *UsersService) AuthenticateUser(ctx context.Context, request users.AuthenticateUserRequest) (*users.AuthenticateUserResponse, apierrors.ApiError) {
	hashedPass, userID, err := s.usersDao.GetUserHashedPassword(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if !crypto_utils.CheckPasswordHash(request.Password, []byte(*hashedPass)) {
		return nil, apierrors.NewUnauthorizedApiError("invalid credentials")
	}

	userData, err := s.usersDao.GetUserDataByID(ctx, *userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := auth_utils.GenerateToken(userData.UserID, auth_utils.TokenTypeAccess)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.usersDao.GetUserRefreshToken(ctx, userData.UserID)
	if err != nil {
		return nil, err
	}

	return &users.AuthenticateUserResponse{
		AccessToken:  *accessToken,
		TokenType:    auth_utils.Bearer,
		ExpiresIn:    config.AccessTokenExpirationMs,
		UserId:       userData.UserID,
		RefreshToken: *refreshToken,
	}, nil
}

func (s *UsersService) RefreshUserToken(ctx context.Context, request users.RefreshUserTokenRequest) (*users.AuthenticateUserResponse, apierrors.ApiError) {
	userRefreshToken, err := s.usersDao.GetUserRefreshToken(ctx, request.UserID)
	if err != nil {
		return nil, err
	}
	if *userRefreshToken != request.RefreshToken {
		return nil, apierrors.NewUnauthorizedApiError("invalid refresh token")
	}

	accessToken, err := auth_utils.GenerateToken(request.UserID, auth_utils.TokenTypeAccess)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := auth_utils.GenerateToken(request.UserID, auth_utils.TokenTypeRefresh)
	if err != nil {
		return nil, err
	}

	if err := s.usersDao.UpdateUserRefreshToken(ctx, request.UserID, *newRefreshToken); err != nil {
		logger_utils.Logger.Println("Successfully rotated refresh token for user ", request.UserID)
		return nil, err
	}

	return &users.AuthenticateUserResponse{
		AccessToken:  *accessToken,
		TokenType:    auth_utils.Bearer,
		ExpiresIn:    config.AccessTokenExpirationMs,
		UserId:       request.UserID,
		RefreshToken: *newRefreshToken,
	}, nil
}
