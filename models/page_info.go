package models

type PageInfo struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	List     interface{} `json:"list"`
}

func NewPageInfo(pageInfo PageInfo) *PageInfo {
	return &PageInfo{PageNum: pageInfo.PageNum, PageSize: pageInfo.PageSize, Total: pageInfo.Total, List: pageInfo.List}
}
