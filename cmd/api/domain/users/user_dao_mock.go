package users

import (
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"golang.org/x/net/context"
)

type UserDaoMock struct {
	MockCreateUser             func(ctx context.Context, user RegisterUserRequest) (*User, apierrors.ApiError)
	MockGetUserHashedPassword  func(ctx context.Context, email string) (*string, *int64, apierrors.ApiError)
	MockGetUserDataByID        func(ctx context.Context, userID int64) (*User, apierrors.ApiError)
	MockUpdateUserRefreshToken func(ctx context.Context, userID int64, refreshToken string) apierrors.ApiError
	MockGetUserRefreshToken    func(ctx context.Context, userID int64) (*string, apierrors.ApiError)
}

func (m *UserDaoMock) CreateUser(ctx context.Context, user RegisterUserRequest) (*User, apierrors.ApiError) {
	return m.MockCreateUser(ctx, user)
}

func (m *UserDaoMock) GetUserHashedPassword(ctx context.Context, email string) (*string, *int64, apierrors.ApiError) {
	return m.MockGetUserHashedPassword(ctx, email)
}

func (m *UserDaoMock) GetUserDataByID(ctx context.Context, userID int64) (*User, apierrors.ApiError) {
	return m.MockGetUserDataByID(ctx, userID)
}

func (m *UserDaoMock) UpdateUserRefreshToken(ctx context.Context, userID int64, refreshToken string) apierrors.ApiError {
	return m.MockUpdateUserRefreshToken(ctx, userID, refreshToken)
}

func (m *UserDaoMock) GetUserRefreshToken(ctx context.Context, userID int64) (*string, apierrors.ApiError) {
	return m.MockGetUserRefreshToken(ctx, userID)
}
