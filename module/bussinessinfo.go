package module

type Bussinessinfo struct {
	Id         int64  `xorm:"pk autoincr BIGINT(20)"`
	Username   string `xorm:"not null default ''0'' VARCHAR(50)"`
	Legal      string `xorm:"not null default ''0'' VARCHAR(50)"`
	Info       string `xorm:"not null default ''0'' VARCHAR(2000)"`
	Website    string `xorm:"not null default ''0'' VARCHAR(80)"`
	Company    string `xorm:"not null default ''0'' VARCHAR(80)"`
	RegisterId string `xorm:"not null default ''0'' VARCHAR(80)"`
}
