package services

import (
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/users"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNewUsersService(t *testing.T) {
	var userDao users.UserDaoInterface = &users.UserDaoMock{MockGetUserDataByID: func(ctx context.Context, userID int64) (*users.User, apierrors.ApiError) {
		return &users.User{UserID: 1}, nil
	}}

	service := *NewUsersService(&userDao)
	assert.NotNil(t, service)
}

func TestUsersService_GetUserByID_NoError(t *testing.T) {
	var userDao users.UserDaoInterface = &users.UserDaoMock{MockGetUserDataByID: func(ctx context.Context, userID int64) (*users.User, apierrors.ApiError) {
		return &users.User{UserID: 1}, nil
	}}

	service := &UsersService{
		usersDao: userDao,
	}
	res, err := service.GetUserByID(context.Background(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 1, res.UserID)
}

func TestUsersService_RegisterUser_HashErrorPasswordTooLarge(t *testing.T) {
	var userDao users.UserDaoInterface = &users.UserDaoMock{}

	service := &UsersService{
		usersDao: userDao,
	}

	res, err := service.RegisterUser(context.Background(), users.RegisterUserRequest{
		UserData: &users.User{},
		Password: []byte("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$"),
	})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "password cannot be longer than 72 bytes", err.Message())
}
