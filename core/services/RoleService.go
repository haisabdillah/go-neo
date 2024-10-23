package services

import (
	"github.com/haisabdillah/golang-auth/core/dto"
	"github.com/haisabdillah/golang-auth/core/models"
	"github.com/haisabdillah/golang-auth/pkg/errors"
	"github.com/haisabdillah/golang-auth/pkg/validation"
)

func (s *Service) RoleCreate(dto *dto.RoleDto) error {
	if err := validation.Validate(dto); err != nil {
		return errors.Validation(err)
	}
	// Prepare the permissions slice
	var permissions []models.Permission
	if err := s.db.Where("id IN ?", dto.Permissions).Find(&permissions).Error; err != nil {
		return errors.InternalServer(err)
	}
	data := models.Role{
		Name:        dto.Name,
		Permissions: permissions,
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

func (s *Service) RoleUpdate(id string, dto *dto.RoleDto) error {
	if err := validation.Validate(dto); err != nil {
		return errors.Validation(err)
	}

	data := models.Role{}
	if err := s.db.Preload("Permissions").First(&data, id).Error; err != nil {
		return errors.InternalServer(err)
	}

	// Prepare the permissions slice
	var permissions []models.Permission
	if err := s.db.Where("id IN ?", dto.Permissions).Find(&permissions).Error; err != nil {
		return errors.InternalServer(err)
	}

	//Check Unique
	var count int64
	if err := s.db.Model(models.Role{}).Where("id != ?", id).Where("name = ?", dto.Name).Count(&count).Error; err != nil {
		return errors.InternalServer(err)
	}
	if count > 0 {
		return errors.Validation(map[string]string{"name": "UNIQUE"})
	}

	tx := s.db.Begin()
	//Delete Assosiation Permissions
	if err := tx.Model(&data).Association("Permissions").Clear(); err != nil {
		tx.Rollback()
		return errors.InternalServer(err)
	}

	data.Name = dto.Name
	data.Permissions = permissions

	//Update Data
	if err := tx.Save(&data).Error; err != nil {
		tx.Rollback()
		return errors.InternalServer(err)
	}
	tx.Commit()
	return nil
}

func (s *Service) RoleDelete(id string) error {
	if err := s.db.Delete(models.Role{}, id).Error; err != nil {
		return errors.InternalServer(err)
	}
	return nil
}

func (s *Service) RoleGet() (interface{}, error) {

	var data []models.Role
	if err := s.db.Model(models.Role{}).Preload("Permissions").Find(&data).Error; err != nil {
		return nil, errors.InternalServer(err)
	}
	var result []map[string]any
	for _, v := range data {
		result = append(result, map[string]any{"id": v.ID, "name": v.Name})
	}
	return result, nil
}

func (s *Service) RoleFirst(id string) (interface{}, error) {

	var data []models.Role
	if err := s.db.Model(models.Role{}).Preload("Permissions").First(&data, id).Error; err != nil {
		return nil, errors.InternalServer(err)
	}
	return data, nil
}
