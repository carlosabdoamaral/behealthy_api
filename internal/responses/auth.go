package responses

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Role     string `json:"role"`
}

type RefreshTwoFactorCodeRequest struct {
	Email string `json:"email"`
}
