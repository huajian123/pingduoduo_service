package list_models

type SearchList struct {
	Id int64 `json:"id"`
	Name string `xorm:"varchar(24)" json:"categoryTitle"`
}
