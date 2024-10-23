package models

type Permission struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique;not null" json:"name"` // Contoh: "create_user", "delete_report"
	Level int    `json:"level"`                       // Level izin (misalnya: 1 untuk dasar, 2 untuk menengah, 3 untuk lanjutan)
}

// Role merepresentasikan peran dalam sistem.
type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique;not null" json:"name"`                                                // Contoh: "finance", "admin"
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:onDelete:cascade;" json:"permissions"` // Relasi many-to-many
}
