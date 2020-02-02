package module

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Username       string    `xorm:"not null VARCHAR(50)"`
	PasswordDigest string    `xorm:"not null VARCHAR(150)"`
	Email          string    `xorm:"not null VARCHAR(50)"`
	Phone          string    `xorm:"not null VARCHAR(50)"`
	RoleId         int       `xorm:"not null INT(11)"`
	CreateTime     time.Time `xorm:"default 'NULL' DATETIME"`
	Avatar         string    `xorm:"default 'NULL' VARCHAR(150)"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
func GetUser(user_name interface{}) (User, error) {
	var user User
	_, err := DB.Where("username = ?", user_name).Get(&user)
	return user, err
}
