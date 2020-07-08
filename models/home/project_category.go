package home_model

type ProjectCategory struct {
	Id            int64  `json:"id"`
	CategoryTitle string `xorm:"varchar(24)" json:"categoryTitle"`
}
