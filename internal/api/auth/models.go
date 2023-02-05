package auth

type SignInRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequestModel struct {
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Role     string `json:"role"`
}

type RecoverPasswordRequestModel struct {
	Email string `json:"email"`
}

type RecoverPasswordValidationRequestModel struct {
	NewPassword   string `json:"new_password"`
	TwoFactorCode string `json:"two_factor_code"`
}
