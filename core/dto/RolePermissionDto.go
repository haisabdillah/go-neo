package dto

type PermissionDto struct {
	Name  string `json:"name" validate:"required"`
	Level int    `json:"level" validate:"required,number"`
}

type RoleDto struct {
	Name        string ` json:"name" validate:"required"`        // Contoh: "finance", "admin"
	Permissions []uint ` json:"permissions" validate:"required"` // Contoh: "finance", "admin"
}
