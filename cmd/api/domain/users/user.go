package users

type User struct {
	UserID       int64  `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	ContactPhone string `json:"contact_phone"`
	UserType     string `json:"user_type"`
}

type RegisterUserRequest struct {
	UserData *User  `json:"user"`
	Password []byte `json:"password"`
}

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type RefreshUserTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
}

type AuthenticateUserResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	UserId       int64  `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
