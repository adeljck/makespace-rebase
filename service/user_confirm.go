package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"makespace-remaster/middleware"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"strings"
)

type TeacherConfirm struct {
	Name      string `form:"name" json:"name" validate:"required,min=2,max=6"`
	Academy   string `form:"academy" json:"academy" validate:"required,min=3,max=20"`
	TeacherId string `form:"teacherid" json:"teacherid" validate:"required,min=8,max=15"`
	Knowledge string `form:"knowledge" json:"knowledge" validate:"required,min=4,max=15"`
}
type StudentConfirm struct {
	Name      string `form:"name" json:"name" validate:"required,min=2,max=6"`
	Academy   string `form:"academy" json:"academy" validate:"required,min=3,max=20"`
	StudentId string `form:"teacherid" json:"teacherid" validate:"required,min=8,max=15"`
	Major     string `form:"major" json:"major" validate:"required,min=4,max=15"`
	Class     string `form:"class" json:"class" validate:"required,min=4,max=10"`
}
type BussinessConfirm struct {
	Legal      string `form:"legal" json:"legal" validate:"required,min=2,max=6"`
	Info       string `form:"info" json:"info" validate:"required,max=150"`
	Website    string `form:"pwebsite" json:"website" validate:"required,min=8,max=40"`
	Company    string `form:"company" json:"company" validate:"required,min=2,max=25"`
	RegisterId string `form:"rehisterid" json:"registerid" validate:"required,min=18,max=18"`
}

//表单验证
func (service *BussinessConfirm) valid() *serializer.Response {
	user := new(module.Bussinessinfo)
	var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
	if count, _ := module.DB.Where("name=?", service.Company).Count(user); count > 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "name",
			Error: "此公司已入驻",
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

//表单验证
func (service *TeacherConfirm) valid() *serializer.Response {
	user := new(module.Bussinessinfo)
	var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
	if count, _ := module.DB.Where("name=?", service.Name).Count(user); count > 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "name",
			Error: "此身份已被注册",
		})
	}
	if count, _ := module.DB.Where("teacher_id=?", service.TeacherId).Count(user); count > 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "name",
			Error: "此身份已被注册",
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

//表单验证
func (service *StudentConfirm) valid() *serializer.Response {
	user := new(module.Bussinessinfo)
	var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
	if count, _ := module.DB.Where("name=?", service.Name).Count(user); count > 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "name",
			Error: "此身份已被注册",
		})
	}
	if count, _ := module.DB.Where("teacher_id=?", service.StudentId).Count(user); count > 0 {
		TagErrors = append(TagErrors, serializer.TagError{
			Tag:   "name",
			Error: "此身份已被注册",
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
func (service *BussinessConfirm) confirm() *serializer.Response {
	if err := TagValid(service); err != nil {
		return err
	}
	if err := service.valid(); err != nil {
		return err
	}
	return nil
}
func (service *TeacherConfirm) confirm() *serializer.Response {
	if err := TagValid(service); err != nil {
		return err
	}
	if err := service.valid(); err != nil {
		return err
	}
	return nil
}
func (service *StudentConfirm) confirm() *serializer.Response {
	if err := TagValid(service); err != nil {
		return err
	}
	if err := service.valid(); err != nil {
		return err
	}
	return nil
}
func ConfirmService(c *gin.Context) (*serializer.Response, error) {
	j := middleware.NewJWT()
	claims, err := j.ParseToken(c.Request.Header.Get("token"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(claims.Id)
	var result *serializer.Response
	switch strings.ToLower(c.Param("type")) {
	case "teacher":
		var service TeacherConfirm
		if err := c.ShouldBindJSON(&service); err == nil {
			if err := service.confirm(); err != nil {
				return err, nil
			}
		} else {
			return nil, err
		}
	case "student":
		var service StudentConfirm
		if err := c.ShouldBindJSON(&service); err == nil {
			if err := service.confirm(); err != nil {
				return err, nil
			}
		} else {
			return nil, err
		}
	case "bussiness":
		var service BussinessConfirm
		if err := c.ShouldBindJSON(&service); err == nil {
			if err := service.confirm(); err != nil {
				return err, nil
			}
		} else {
			return nil, err
		}
	default:
		return &serializer.Response{
			Status: 404,
			Msg:    "wrong type",
		}, nil
	}
	return result, nil
}
