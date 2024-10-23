package services

import (
	"github.com/haisabdillah/golang-auth/core/dto"
	"github.com/haisabdillah/golang-auth/core/models"
	"github.com/haisabdillah/golang-auth/pkg/errors"
	"github.com/haisabdillah/golang-auth/pkg/validation"
)

func (s *Service) PermissionCreate(dto *dto.PermissionDto) error {
	if err := validation.Validate(dto); err != nil {
		return errors.Validation(err)
	}
	data := models.Permission{
		Name:  dto.Name,
		Level: dto.Level,
	}
	var count int64
	if err := s.db.Model(&data).Where("name = ?", data.Name).Count(&count).Error; err != nil {
		return errors.InternalServer(err)
	}
	if count > 0 {
		return errors.Validation(map[string]string{"name": "UNIQUE"})
	}
	if err := s.db.Create(&data).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}

func (s *Service) PermissionGet() (interface{}, error) {
	var data []models.Permission
	if err := s.db.Find(&data).Error; err != nil {
		return nil, errors.InternalServer(err)
	}
	return data, nil
}

func (s *Service) PermissionFirst(id string) (interface{}, error) {
	var data models.Permission
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, errors.InternalServer(err)
	}
	return data, nil
}

func (s *Service) PermissionUpdate(id string, dto *dto.PermissionDto) error {
	if err := validation.Validate(dto); err != nil {
		return errors.Validation(err)
	}
	data := models.Permission{}
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return errors.InternalServer(err)
	}
	var count int64
	if err := s.db.Model(models.Permission{}).Where("id != ?", id).Where("name = ?", dto.Name).Count(&count).Error; err != nil {
		return errors.InternalServer(err)
	}
	if count > 0 {
		return errors.Validation(map[string]string{"name": "UNIQUE"})
	}
	data.Name = dto.Name
	data.Level = dto.Level
	if err := s.db.Save(&data).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}

func (s *Service) PermissionDelete(id string) error {
	data := models.Permission{}
	if err := s.db.Where("id = ?", id).Delete(&data).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}
