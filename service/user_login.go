package service

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
	Password string `form:"password" json:"password" validate:"required,min=8,max=40"`
}

func (userLoginService *UserLoginService) valid() *serializer.Response {
	validate := validator.New()
	if err := validate.Struct(userLoginService); err != nil {
		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(validate, trans)
		var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			tagerror := serializer.TagError{
				Tag:   err.Field(),
				Error: err.Translate(trans),
			}
			TagErrors = append(TagErrors, tagerror)
		}
		return &serializer.Response{
			Status: 40001,
			Data:   TagErrors,
			Msg:    "tag error",
		}
	}
	return nil
}
func (service *UserLoginService) Login() (module.User, *serializer.Response) {
	var user module.User
	if err := service.valid(); err != nil {
		return user, err
	}
	if !module.DB.Where("username = ?", service.UserName).GetFirst(&user).Has {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}
