package service

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"reflect"
	"time"
)

var blackusername = []string{"admin", "user", "supperuser", "test"}

func isExistItem(value interface{}, slice interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

//注册
type UserRegiste struct {
	UserName        string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
	Password        string `form:"password" json:"password" validate:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" validate:"required,min=8,max=40"`
	Phone           string `form:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email           string `form:"email" json:"email" validate:"required,email,min=8,max=50"`
}

//字段验证
func (service *UserRegiste) tagValid() *serializer.Response {
	validate := validator.New()
	if err := validate.Struct(service); err != nil {
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
			Status: 40002,
			Data:   TagErrors,
			Msg:    "tag error",
		}
	}
	return nil
}

//表单验证
func (service *UserRegiste) valid() *serializer.Response {
	user := new(module.User)
	var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
	if service.PasswordConfirm != service.Password {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "password",
			Error: "两次输入的密码不相同",
		})
	}
	if ok := isExistItem(service.UserName, blackusername); ok {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "username",
			Error: "不允许使用的用户名",
		})
	}
	if count, _ := module.DB.Where("username=?", service.UserName).Count(user); count != 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "username",
			Error: "用户名已被注册",
		})
	}
	if count, _ := module.DB.Where("email=?", service.Email).Count(user); count != 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "email",
			Error: "邮箱已被注册",
		})
	}
	if count, _ := module.DB.Where("phone=?", service.Phone).Count(user); count != 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "phone",
			Error: "手机号码已被注册",
		})
	}
	if len(TagErrors) == 0 {
		return nil
	}
	return &serializer.Response{
		Status: 4001,
		Data:   TagErrors,
		Msg:    "Tag Error",
	}
}

//用户注册
func (service *UserRegiste) Registe() (module.User, *serializer.Response) {
	if err := service.tagValid(); err != nil {
		return module.User{}, err
	}
	if err := service.valid(); err != nil {
		return module.User{}, err
	}
	user := module.User{
		Username:   service.UserName,
		Email:      service.Email,
		Phone:      service.Phone,
		RoleId:     module.Unauthorized,
		CreateTime: time.Now(),
	}
	if err := user.SetPassword(service.PasswordConfirm); err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "password encode failed",
		}
	}
	if affected, err := module.DB.InsertOne(user); err != nil || affected == 0 {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}
	return user, nil
}
