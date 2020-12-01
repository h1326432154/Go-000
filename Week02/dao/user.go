package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	// ErrFooUserNotFound 用户未找到
	ErrFooUserNotFound = errors.New("user not found")
)

// User 用户表结构
type User struct {
	ID   int    `gorm:"column:u_id" json:"id"`
	Name string `gorm:"column:u_nic" json:"name"`
}

// MockDataNotFound .
func (u *User) MockDataNotFound() error {
	return ErrFooUserNotFound
}

// GetUserByID 获取用户信息
func (u *User) GetUserByID() error {
	query := "SELECT u_id,u_nic FROM tb_user where u_id = ? limit 1"
	if err := DB.Raw(query, u.ID).Scan(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.Wrap(err, u.MockDataNotFound().Error())
		}
		return err
	}
	return nil
}
