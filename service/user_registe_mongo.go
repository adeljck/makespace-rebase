package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/serializer"
)

type Student struct {
	UserName        string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
	Password        string `form:"password" json:"password" validate:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" validate:"required,min=8,max=40"`
	Phone           string `form:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email           string `form:"email" json:"email" validate:"required,email,min=8,max=50"`
	Name            string `form:"name" json:"name" validate:"required,min=2,max=6"`
	Academy         string `form:"academy" json:"academy" validate:"required,min=3,max=20"`
	StudentId       string `form:"teacherid" json:"teacherid" validate:"required,min=8,max=15"`
	Major           string `form:"major" json:"major" validate:"required,min=4,max=15"`
	Class           string `form:"class" json:"class" validate:"required,min=4,max=10"`
}
type Teacher struct {
	UserName        string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
	Password        string `form:"password" json:"password" validate:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" validate:"required,min=8,max=40"`
	Phone           string `form:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email           string `form:"email" json:"email" validate:"required,email,min=8,max=50"`
	Name            string `form:"name" json:"name" validate:"required,min=2,max=6"`
	Academy         string `form:"academy" json:"academy" validate:"required,min=3,max=20"`
	TeacherId       string `form:"teacherid" json:"teacherid" validate:"required,min=8,max=15"`
	Knowledge       string `form:"knowledge" json:"knowledge" validate:"required,min=4,max=15"`
}
type Bussiness struct {
	UserName        string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
	Password        string `form:"password" json:"password" validate:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" validate:"required,min=8,max=40"`
	Phone           string `form:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email           string `form:"email" json:"email" validate:"required,email,min=8,max=50"`
	Legal           string `form:"legal" json:"legal" validate:"required,min=2,max=6"`
	Info            string `form:"info" json:"info" validate:"required,max=150"`
	Website         string `form:"pwebsite" json:"website" validate:"required,min=8,max=40"`
	Company         string `form:"company" json:"company" validate:"required,min=2,max=25"`
	RegisterId      string `form:"rehisterid" json:"registerid" validate:"required,min=18,max=18"`
}

func (service *Student) valid() *serializer.Response {
	var errors []serializer.PureErrorResponse = make([]serializer.PureErrorResponse, 0)
	if service.PasswordConfirm != service.Password {
		errors = append(errors, serializer.PureErrorResponse{
			Status: 0,
			Msg:    "两次输入的密码不相同",
		})
	}
	if len(errors) != 0 {
		return &serializer.Response{
			Status: 40001,
			Data:   errors,
			Msg:    "something wrong",
		}
	}
	return nil
}

func (service *Teacher) valid() *serializer.Response {
	var errors []serializer.PureErrorResponse = make([]serializer.PureErrorResponse, 0)
	if service.PasswordConfirm != service.Password {
		errors = append(errors, serializer.PureErrorResponse{
			Status: 0,
			Msg:    "两次输入的密码不相同",
		})
	}
	if len(errors) != 0 {
		return &serializer.Response{
			Status: 40001,
			Data:   errors,
			Msg:    "something wrong",
		}
	}
	return nil
}

// Valid 验证表单
func (service *Bussiness) valid() *serializer.Response {
	var errors []serializer.PureErrorResponse = make([]serializer.PureErrorResponse, 0)
	if service.PasswordConfirm != service.Password {
		errors = append(errors, serializer.PureErrorResponse{
			Status: 0,
			Msg:    "两次输入的密码不相同",
		})
	}
	if len(errors) != 0 {
		return &serializer.Response{
			Status: 40001,
			Data:   errors,
			Msg:    "something wrong",
		}
	}
	return nil
}

func (user *Student) registe() *serializer.Response {
	if err := TagValid(user); err != nil {
		return err
	}
	if err := user.valid(); err != nil {
		return err
	}
	data, err := bson.Marshal(user)
}
