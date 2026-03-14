package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model        // auto add fields ID, CreatedAt, UpdatedAt, DeletedAt
	ID         int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; index:idx_uuid; unique;"`
	RoleName   string `gorm:"column:role_name;"`
	InActive   bool   `gorm:"column:is_active; type:boolean;"`
	RoleNote   string `gorm:"column:role_note; type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
