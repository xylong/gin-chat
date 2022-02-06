package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);uniqueIndex;comment:名称" json:"username" `
	Password string `gorm:"type:varchar(100);comment:密码" json:"password" binding:"required"`
}
