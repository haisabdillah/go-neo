package services

import (
	"github.com/haisabdillah/golang-auth/core/dto"
	"github.com/haisabdillah/golang-auth/core/models"
	"github.com/haisabdillah/golang-auth/pkg/errors"
	"github.com/haisabdillah/golang-auth/pkg/hash"
	"github.com/haisabdillah/golang-auth/pkg/jwt"
	"github.com/haisabdillah/golang-auth/pkg/validation"
	"gorm.io/gorm"
)

func (s *Service) AuthLogin(dto *dto.AuthLoginDto) (interface{}, error) {
	errValidate := validation.Validate(dto)
	if errValidate != nil {
		return nil, errors.Validation(errValidate)
	}
	user := models.User{}

	if err := s.db.Where("email = ?", dto.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Validation(map[string]string{"credentials": "NOT_MATCH"})
		}
		return nil, errors.InternalServer(err)
	}
	// Check password hash
	if !hash.Compare(user.Password, dto.Password) {
		return nil, errors.Validation(map[string]string{"credentials": "NOT_MATCH"})
	}
	payload := jwt.Payload{
		ID:          user.ID,
		Role:        "super_admin",
		Permissions: []string{"Oke", "oke"},
	}
	accessToken, err := jwt.GenerateJWT(payload, 15)
	if err != nil {
		return nil, errors.InternalServer(err)
	}
	refreshToken, err := jwt.GenerateJWT(payload, 48)
	if err != nil {
		return nil, errors.InternalServer(err)
	}
	token := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	return token, nil
}
func (s *Service) AuthMe(id uint) (interface{}, error) {

	user := models.User{}
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Unauthenticate("Invalid token signature")
		}
		return nil, errors.InternalServer(err)
	}
	user.Password = ""
	return user, nil
}

func (s *Service) AuthRefreshToken(id uint) (interface{}, error) {
	user := models.User{}
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Unauthenticate("Invalid token signature")
		}
		return nil, errors.InternalServer(err)
	}
	payload := jwt.Payload{
		ID:          user.ID,
		Role:        "super_admin",
		Permissions: []string{"Oke", "oke"},
	}
	accessToken, err := jwt.GenerateJWT(payload, 15)
	if err != nil {
		return nil, errors.InternalServer(err)
	}
	refreshToken, err := jwt.GenerateJWT(payload, 48)
	if err != nil {
		return nil, errors.InternalServer(err)
	}
	token := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	return token, nil
}

func (s *Service) AuthChangePassword(id uint, dto *dto.AuthChangePasswordDto) error {

	user := models.User{}
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.Unauthenticate("Invalid token signature")
		}
		return errors.InternalServer(err)
	}

	if !hash.Compare(user.Password, dto.OldPassword) {
		return errors.Validation(map[string]string{"old_password": "NOT_MATCH"})
	}

	if dto.NewPassword != dto.NewPasswordConfirmation {
		return errors.Validation(map[string]string{"new_password": "NOT_MATCH"})
	}

	hashPassword, err := hash.Generate(dto.NewPassword)
	if err != nil {
		return errors.InternalServer(err)
	}
	user.Password = hashPassword
	if err := s.db.Updates(&user).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}

func (s *Service) AuthProfile(id uint, dto *dto.AuthProfileDto) error {
	if err := validation.Validate(dto); err != nil {
		return errors.Validation(err)
	}
	user := models.User{}
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.Unauthenticate("Invalid token signature")
		}
		return errors.InternalServer(err)
	}
	// Check if the email already exists for a different user
	var count int64
	if err := s.db.Model(&models.User{}).Where("email = ? AND id != ?", dto.Email, id).Count(&count).Error; err != nil {
		return errors.InternalServer(err)
	}
	if count > 0 {
		return errors.Validation(map[string]string{"email": "email has already been used"})
	}
	user.Name = dto.Name
	user.Email = dto.Email
	if err := s.db.Updates(&user).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}
