package services

import (
	"github.com/haisabdillah/golang-auth/core/dto"
	"github.com/haisabdillah/golang-auth/core/models"
	"github.com/haisabdillah/golang-auth/pkg/errors"
	"github.com/haisabdillah/golang-auth/pkg/hash"
)

func (s *Service) UserCreate(req *dto.UserDto) error {
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	existEmail := models.User{}
	s.db.Model(models.User{}).Where("email =?", user.Email).First(&existEmail)
	if existEmail.Email != "" {
		return errors.BadRequest(map[string]string{"email": "UNIQUE"})
	}
	hashPassword, err := hash.Generate(user.Password)
	if err != nil {
		return errors.InternalServer(err)
	}
	user.Password = hashPassword
	if err := s.db.Create(&user).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}

func (s *Service) UserFirst(id string) (interface{}, error) {

	result := models.User{}
	if err := s.db.Where("id =?", id).First(&result).Error; err != nil {
		return nil, errors.NotFound("User not found")
	}
	return result, nil
}
