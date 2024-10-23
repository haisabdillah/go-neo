package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"` // Relasi ke Role
	Level    int    `json:"level"`                         // Level pengguna (misalnya: 1, 2, 3)
}
