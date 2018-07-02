package base

import ()

type VavError struct {
	Err      error
	Code     int
	Describe string
}

func (e *VavError) Error() string {
	return e.Err.Error()
}

func (e *VavError) Desc() string {
	return e.Describe
}

func NewErr(err error, code int, desc string) *VavError {
	return &VavError{
		Err:      err,
		Code:     code,
		Describe: desc,
	}
}

const (
	ERR_COMMON_NOT_CAPTURE_CODE int    = 999
	ERR_COMMON_NOT_CAPTURE_DESC string = "未捕获的错误"

	ERR_NONE_CODE int    = 0
	ERR_NONE_DESC string = "成功"

	ERR_VEH_OFFLINE_CODE int    = 1
	ERR_VEH_OFFLINE_DESC string = "车机离线"
)

var (
	ERROR_NONE        *VavError = NewErr(nil, ERR_NONE_CODE, ERR_NONE_DESC)
	ERROR_NOT_CAPTURE *VavError = NewErr(nil, ERR_COMMON_NOT_CAPTURE_CODE, ERR_COMMON_NOT_CAPTURE_DESC)
	ERROR_VEH_OFFLINE *VavError = NewErr(nil, ERR_DTU_OFFLINE_CODE, ERR_DTU_OFFLINE_DESC)
)
