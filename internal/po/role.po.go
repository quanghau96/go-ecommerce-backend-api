package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type: int; not null; primaryKey; autoIncrement; comment: 'Primary key is ID'"`
	RoleName string `gorm:"column:role_name"`
	IsActive bool   `gorm:"column:is_active; type: boolean;"`
	RoleRote string `gorm:"column:role_rote; type: text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
