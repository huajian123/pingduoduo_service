package utils

const (
	RECODE_OK = "0"
	/****************************** 通用系统错误 **************************************/
	RECODE_GLOBAL               = "1000"
	RECODE_CREATE_FAILED        = "1001"
	RECODE_UPDATE_FAILED        = "1002"
	RECODE_DELETE_FAILED        = "1003"
	RECODE_SEARCH_FAILED        = "1004"
	RECODE_COUNT_FAILED         = "1005"
	RECODE_DUPLICATED_DATA      = "1006"
	RECODE_RECORD_NOT_FOUND     = "1007"
	RECODE_RECORD_FILE_DOWNLOAD = "1011"
	RECODE_RECORD_FILE_UPLOAD   = "1012"
	/****************************** http请求错误 **************************************/
	RECODE_HTTP_METHOD_NOT_ALLOWED     = "2001"
	RECODE_HTTP_UNSUPPORTED_MEDIA_TYPE = "2002"

	/****************************** 请求校验错误 **************************************/
	RECODE_VALIDATION_MISSING_PARAMS         = "3001"
	RECODE_VALIDATION_PARAMS_ERROR           = "3002"
	RECODE_VALIDATION_PARAMS_TYPE_ERROR      = "3003"
	RECODE_VALIDATION_PARAMS_JSON_TYPE_ERROR = "3004"
	RECODE_VALIDATION_BODY_JSON_TYPE_ERROR   = "3005"

	RECODE_DBERR      = "4001"
	RECODE_NODATA     = "4002"
	RECODE_DATAEXIST  = "4003"
	RECODE_DATAERR    = "4004"
	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERERR    = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_REQERR     = "4201"
	RECODE_IPERR      = "4202"
	RECODE_THIRDERR   = "4301"
	RECODE_IOERR      = "4302"
	RECODE_SERVERERR  = "4500"
	RECODE_UNKNOWERR  = "4501"
)

var recodeText = map[string]string{
	RECODE_OK: "成功",

	/****************************** 通用系统错误 **************************************/
	RECODE_GLOBAL:               "系统处理异常，请稍后重试",
	RECODE_CREATE_FAILED:        "新增数据失败",
	RECODE_UPDATE_FAILED:        "修改数据失败",
	RECODE_DELETE_FAILED:        "删除数据失败",
	RECODE_SEARCH_FAILED:        "查询数据失败",
	RECODE_COUNT_FAILED:         "查询数据总数失败",
	RECODE_DUPLICATED_DATA:      "数据已存在",
	RECODE_RECORD_NOT_FOUND:     "数据不存在",
	RECODE_RECORD_FILE_DOWNLOAD: "文件下载失败",
	RECODE_RECORD_FILE_UPLOAD:   "文件上传失败",
	/****************************** http请求错误 **************************************/
	RECODE_HTTP_METHOD_NOT_ALLOWED:     "不支持当前请求方法",
	RECODE_HTTP_UNSUPPORTED_MEDIA_TYPE: "不支持的媒体类型",

	/****************************** 请求校验错误 **************************************/
	RECODE_VALIDATION_MISSING_PARAMS:         "请求参数校验失败",
	RECODE_VALIDATION_PARAMS_ERROR:           "请求参数校验失败",
	RECODE_VALIDATION_PARAMS_TYPE_ERROR:      "请求参数校验失败",
	RECODE_VALIDATION_PARAMS_JSON_TYPE_ERROR: "请求参数校验失败",
	RECODE_VALIDATION_BODY_JSON_TYPE_ERROR:   "请求参数校验失败",

	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
