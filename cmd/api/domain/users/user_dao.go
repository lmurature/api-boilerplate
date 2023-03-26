package users

import (
	"errors"
	"fmt"
	"github.com/lmurature/api-boilerplate/cmd/api/clients"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"github.com/lmurature/api-boilerplate/cmd/api/utils/date"
	"github.com/lmurature/api-boilerplate/cmd/api/utils/logger"
	"golang.org/x/net/context"
	"time"
)

const (
	insertUser             = "INSERT INTO user(first_name, last_name, email, password, contact_phone, user_type, date_created) VALUES(?,?,?,?,?,?,?);"
	getUserHashedPass      = "SELECT u.password, u.id FROM user u WHERE u.email=?;"
	getUserByID            = "SELECT u.id, u.first_name, u.last_name, u.email, u.contact_phone, u.user_type FROM user u WHERE u.id=?;"
	getUserRefreshToken    = "SELECT u.refresh_token FROM user u WHERE u.id=?;"
	updateUserRefreshToken = "UPDATE user SET refresh_token=? WHERE id=?;"
)

type UserDao struct {
	database *clients.DbClient
}

func NewUserDao(dbClient *clients.DbClient) *UserDaoInterface {
	var dao UserDaoInterface = &UserDao{database: dbClient}
	return &dao
}

type UserDaoInterface interface {
	CreateUser(ctx context.Context, user RegisterUserRequest) (*User, apierrors.ApiError)
	GetUserHashedPassword(ctx context.Context, email string) (*string, *int64, apierrors.ApiError)
	GetUserDataByID(ctx context.Context, userID int64) (*User, apierrors.ApiError)
	UpdateUserRefreshToken(ctx context.Context, userID int64, refreshToken string) apierrors.ApiError
	GetUserRefreshToken(ctx context.Context, userID int64) (*string, apierrors.ApiError)
}

func (dao *UserDao) CreateUser(ctx context.Context, u RegisterUserRequest) (*User, apierrors.ApiError) {
	stmt, err := dao.database.Client.Prepare(insertUser)
	if err != nil {
		logger_utils.Logger.Println(err)
		return nil, apierrors.NewInternalServerApiError("error when trying to insert user", errors.New("db error"))
	}
	defer stmt.Close()

	res, saveErr := stmt.Exec(u.UserData.FirstName, u.UserData.LastName, u.UserData.Email, u.Password, u.UserData.ContactPhone, u.UserData.UserType, time.Now().UTC().Format(date.DatetimeLayout))
	if saveErr != nil {
		logger_utils.Logger.Println(saveErr)
		return nil, apierrors.NewInternalServerApiError("error when trying to save user", errors.New("db error"))
	}

	userID, _ := res.LastInsertId()
	u.UserData.UserID = userID

	return u.UserData, nil
}

func (dao *UserDao) UpdateUserRefreshToken(ctx context.Context, userID int64, refreshToken string) apierrors.ApiError {
	stmt, err := dao.database.Client.Prepare(updateUserRefreshToken)
	if err != nil {
		logger_utils.Logger.Println("error when trying to prepare update user statement", err)
		return nil
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(refreshToken, userID)
	if updateErr != nil {
		logger_utils.Logger.Println("error when trying to update user", updateErr)
		return apierrors.NewInternalServerApiError("error when trying to update user", errors.New("db error"))
	}

	logger_utils.Logger.Println(fmt.Sprintf("successfully updated user %d", userID))
	return nil
}

func (dao *UserDao) GetUserHashedPassword(ctx context.Context, email string) (*string, *int64, apierrors.ApiError) {
	stmt, err := dao.database.Client.Prepare(getUserHashedPass)
	if err != nil {
		logger_utils.Logger.Println(err)
		return nil, nil, apierrors.NewInternalServerApiError("error when trying to prepare select user", errors.New("db error"))
	}
	defer stmt.Close()

	res := stmt.QueryRow(email)
	if res.Err() != nil {
		logger_utils.Logger.Println(res.Err())
		return nil, nil, apierrors.NewInternalServerApiError("error when trying to get user hashed pass", errors.New("db error"))
	}

	var hashedPass string
	var userID int64
	if err := res.Scan(&hashedPass, &userID); err != nil {
		logger_utils.Logger.Println(err)
		return nil, nil, apierrors.NewNotFoundApiError("email not found")
	}

	return &hashedPass, &userID, nil
}

func (dao *UserDao) GetUserDataByID(ctx context.Context, userID int64) (*User, apierrors.ApiError) {
	stmt, err := dao.database.Client.Prepare(getUserByID)
	if err != nil {
		logger_utils.Logger.Println(err)
		return nil, apierrors.NewInternalServerApiError("error when trying to prepare select user", errors.New("db error"))
	}
	defer stmt.Close()

	res := stmt.QueryRow(userID)
	if res.Err() != nil {
		logger_utils.Logger.Println(res.Err())
		return nil, apierrors.NewInternalServerApiError("error when trying to get user data from email", errors.New("db error"))
	}

	var userResult User
	if err := res.Scan(&userResult.UserID, &userResult.FirstName, &userResult.LastName, &userResult.Email, &userResult.ContactPhone, &userResult.UserType); err != nil {
		logger_utils.Logger.Println(err)
		return nil, apierrors.NewNotFoundApiError("user not found")
	}

	return &userResult, nil
}

func (dao *UserDao) GetUserRefreshToken(ctx context.Context, userID int64) (*string, apierrors.ApiError) {
	stmt, err := dao.database.Client.Prepare(getUserRefreshToken)
	if err != nil {
		logger_utils.Logger.Println(err)
		return nil, apierrors.NewInternalServerApiError("error when trying to prepare select user", errors.New("db error"))
	}
	defer stmt.Close()

	res := stmt.QueryRow(userID)
	if res.Err() != nil {
		logger_utils.Logger.Println(res.Err())
		return nil, apierrors.NewInternalServerApiError("error when trying to get user refresh_token", errors.New("db error"))
	}

	var refreshToken string
	if err := res.Scan(&refreshToken); err != nil {
		logger_utils.Logger.Println(err)
		return nil, apierrors.NewNotFoundApiError("refresh token not found")
	}

	return &refreshToken, nil
}
