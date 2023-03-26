package services

import (
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/users"
	"golang.org/x/net/context"
)

type UsersServiceMock struct {
}

func (u UsersServiceMock) GetUserByID(ctx context.Context, userID int64) (*users.User, apierrors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (u UsersServiceMock) RegisterUser(ctx context.Context, request users.RegisterUserRequest) (*users.User, apierrors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (u UsersServiceMock) AuthenticateUser(ctx context.Context, request users.AuthenticateUserRequest) (*users.AuthenticateUserResponse, apierrors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (u UsersServiceMock) RefreshUserToken(ctx context.Context, request users.RefreshUserTokenRequest) (*users.AuthenticateUserResponse, apierrors.ApiError) {
	//TODO implement me
	panic("implement me")
}
