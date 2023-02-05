package responses

import "time"

type AccountDetails struct {
	AccountId     int64     `json:"account_id"`
	FullName      string    `json:"full_name"`
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	BirthDate     time.Time `json:"birth_date"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	SoftDeleted   bool      `json:"soft_deleted"`
	IsBlocked     bool      `json:"is_blocked"`
	TwoFactorCode string    `json:"two_factor_code"`
}

type UpdateAccountRequest struct {
	AccountId     int64     `json:"account_id"`
	FullName      string    `json:"full_name"`
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	BirthDate     time.Time `json:"birth_date"`
	Role          string    `json:"role"`
	SoftDeleted   bool      `json:"soft_deleted"`
	IsBlocked     bool      `json:"is_blocked"`
	TwoFactorCode string    `json:"two_factor_code"`
}

type SoftDeleteAccountRequest struct {
	AccountId     int64  `json:"account_id"`
	TwoFactorCode string `json:"two_factor_code"`
}

type RestoreAccountRequest struct {
	AccountId     int64  `json:"account_id"`
	TwoFactorCode string `json:"two_factor_code"`
}

type DeleteAccountRequest struct {
	AccountId     int64  `json:"account_id"`
	TwoFactorCode string `json:"two_factor_code"`
}

type UpdatePasswordRequest struct {
	Email         string `json:"email"`
	NewPassword   string `json:"new_password"`
	TwoFactorCode string `json:"two_factor_code"`
}
