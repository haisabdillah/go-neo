package dto

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (*UserDto) Validation() error {
	return nil
}
