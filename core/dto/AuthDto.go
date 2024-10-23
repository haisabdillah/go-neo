package dto

type AuthLoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthChangePasswordDto struct {
	OldPassword             string `json:"old_password"`
	NewPassword             string `json:"new_password"`
	NewPasswordConfirmation string `json:"new_password_confirmation"`
}

type AuthProfileDto struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email"`
}
