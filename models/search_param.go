package models

type SearchParam struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Filters  interface{} `json:"filters"`
}

func NewSearchParam(searchParam SearchParam) *SearchParam {
	return &SearchParam{PageSize: searchParam.PageSize, PageNum: searchParam.PageNum, Filters: searchParam.Filters}
}
