package user_model

type User struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"VARCHAR(20)"`
	Password string `xorm:"VARCHAR(20)"`
	IsManage int    `xorm:"comment('1管理员0非管理员') TINYINT(1)"`
}
