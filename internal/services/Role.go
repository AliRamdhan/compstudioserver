package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (rs *RoleService) CreateRole(role *model.Role) error {
	role.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(role).Error
}

func (rs *RoleService) GetAllRole() ([]model.Role, error) {
	var roles []model.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (rs *RoleService) UpdateRole(roleId uint, updatedRole *model.Role) error {
	// Find the Ticket with the given ID
	var existingRole model.Role
	if err := config.DB.First(&existingRole, "roleId = ?", roleId).Error; err != nil {
		return err // Ticket not found
	}

	// Update fields of existing Ticket with the new values
	existingRole.Name = updatedRole.Name
	existingRole.Description = updatedRole.Description
	existingRole.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated Ticket
	return config.DB.Save(&existingRole).Error
}

func (rs *RoleService) DeleteRole(roleId uint) error {
	// Find the Ticket with the given ID
	var role model.Role
	if err := config.DB.First(&role, "roleId = ?", roleId).Error; err != nil {
		return err // role not found
	}
	// Delete the role
	return config.DB.Delete(&role).Error
}
