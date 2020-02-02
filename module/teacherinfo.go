package module

type Teacherinfo struct {
	Id        int64  `xorm:"pk autoincr BIGINT(20)"`
	Username  string `xorm:"not null VARCHAR(50)"`
	Name      string `xorm:"not null VARCHAR(50)"`
	Academy   string `xorm:"not null VARCHAR(50)"`
	TeacherId string `xorm:"not null VARCHAR(50)"`
	Knowledge string `xorm:"not null VARCHAR(50)"`
}
