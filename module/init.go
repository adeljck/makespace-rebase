package module

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var DB *xorm.Engine

const (
	// PassWordCost 密码加密难度
	PassWordCost     = 12
	Admin        int = 0
	Teacher      int = 1
	Student      int = 2
	Busssiness   int = 3
	Unauthorized int = 4
	UBussiness   int = 5
	Ban          int = 6
)

func Database(connString string) {
	db, err := xorm.NewEngine("mysql", connString)
	if err != nil {
		panic(err)
	}
	if gin.Mode() != "release" {
		db.ShowSQL(true)
		db.Logger().SetLevel(core.LOG_DEBUG)
	}
	//设置连接池
	//空闲
	db.SetMaxIdleConns(20)
	//打开
	db.SetMaxOpenConns(100)
	//超时
	db.SetConnMaxLifetime(time.Second * 30)

	DB = db

	//migration()
}
func migration() {
	// 自动迁移模式
	DB.Charset("utf8mb4")
	err := DB.Sync2(new(User), new(Role), new(Studentinfo), new(Bussinessinfo), new(Teacherinfo))
	if err != nil {
		log.Fatal(err)
	}
}
func CheckAdmin() {
	var admin User
	count, err := DB.Where("role_id=?", 0).Count(&admin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	if count == 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), PassWordCost)
		if err != nil {
			panic(err)
		}
		_, err = DB.Insert(User{
			Username:       os.Getenv("ADMIN_NAME"),
			PasswordDigest: string(password),
			RoleId:         Admin,
			CreateTime:     time.Now(),
		})
		if err != nil {
			panic(err)
		}
	}
}
