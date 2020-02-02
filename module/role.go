package module

type Role struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Role   string `xorm:"not null VARCHAR(50)"`
	RoleId int    `xorm:"not null INT(11)"`
}
