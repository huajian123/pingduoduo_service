package user_model

type User struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"  json:"id"`
	Name     string `xorm:"VARCHAR(20)"  json:"name"`
	Password string `xorm:"VARCHAR(20)"  json:"password"`
	IsManage int    `xorm:"comment('1管理员0非管理员') TINYINT(1)"  json:"isManage"`
}
