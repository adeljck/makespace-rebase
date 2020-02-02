package module

type Studentinfo struct {
	Id        int64  `xorm:"pk autoincr BIGINT(20)"`
	Username  string `xorm:"not null VARCHAR(50)"`
	Academy   string `xorm:"not null VARCHAR(50)"`
	Major     string `xorm:"not null VARCHAR(50)"`
	Class     string `xorm:"not null VARCHAR(50)"`
	StudentId string `xorm:"not null VARCHAR(50)"`
}
