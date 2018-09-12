package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"

	pkgerrors "github.com/pkg/errors"
)

// 错误码
const (
	ErrorCodeNil = 0

	ErrorCodeNotRegister        = 10101 // 未注册
	ErrorCodeNotLogin           = 10102 // 未登录
	ErrorCodeResourcesNotFound  = 10103 // 访问资源不存在
	ErrorCodeResourcesNotPermit = 10104 // 访问资源不被允许
	ErrorCodeDataAbnormal       = 10105 // 访问的数据异常

	ErrorCodeRequestParamsDecode = 10401 // 参数解码错误,包含格式化失败,字段类型错误
	ErrorCodeRequestParamsValue  = 10402 // 参数值错误
	ErrorCodeInsertRepeat        = 10403 // 重复插入

	ErrorCodeServer   = 10500 // 服务端错误
	ErrorCodeDBAccess = 10501 // 数据库访问异常
)

// BaseError base struct
type BaseError struct {
	ErrorCode int
	err       error
}

func (e *BaseError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.err.Error()
}

// NotRegisterError 未注册
type NotRegisterError struct {
	BaseError
}

// NewNotRegisterError make a not register error
func NewNotRegisterError(msg ...interface{}) *NotRegisterError {
	err := &NotRegisterError{}
	err.ErrorCode = ErrorCodeNotRegister
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// NotLoginError  未登录
type NotLoginError struct {
	BaseError
}

// NewNotLoginError make a not register error
func NewNotLoginError(msg ...interface{}) *NotLoginError {
	err := &NotLoginError{}
	err.ErrorCode = ErrorCodeNotLogin
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// ResourcesNotFoundError 资源不存在
type ResourcesNotFoundError struct {
	BaseError
}

// NewResourcesNotFoundError make a resource not found error
func NewResourcesNotFoundError(msg ...interface{}) *ResourcesNotFoundError {
	err := &ResourcesNotFoundError{}
	err.ErrorCode = ErrorCodeResourcesNotFound
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// ResourcesNotPermitError 资源无权限
type ResourcesNotPermitError struct {
	BaseError
}

// NewResourcesNotPermitError make a not permit error
func NewResourcesNotPermitError(msg ...interface{}) *ResourcesNotPermitError {
	err := &ResourcesNotPermitError{}
	err.ErrorCode = ErrorCodeResourcesNotPermit
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// DataAbnormalError 数据异常
type DataAbnormalError struct {
	BaseError
}

// NewDataAbnormalError make a DataAbnormalError error
func NewDataAbnormalError(msg ...interface{}) *DataAbnormalError {
	err := &DataAbnormalError{}
	err.ErrorCode = ErrorCodeDataAbnormal
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// RequestParamsDecodeError 请求数据结构或参数类型错误
type RequestParamsDecodeError struct {
	BaseError
}

// NewRequestParamsDecodeError make a request decode error
func NewRequestParamsDecodeError(msg ...interface{}) *RequestParamsDecodeError {
	err := &RequestParamsDecodeError{}
	err.ErrorCode = ErrorCodeRequestParamsDecode
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// RequestParamsValueError 请求数据值不合法
type RequestParamsValueError struct {
	BaseError
}

// NewRequestParamsValueError make a request paras value error
func NewRequestParamsValueError(msg ...interface{}) *RequestParamsValueError {
	err := &RequestParamsValueError{}
	err.ErrorCode = ErrorCodeRequestParamsValue
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// InsertRepeatError 重复的数据
type InsertRepeatError struct {
	BaseError
}

//NewInsertRepeatError make a insert repeat error
func NewInsertRepeatError(msg ...interface{}) *InsertRepeatError {
	err := &InsertRepeatError{}
	err.ErrorCode = ErrorCodeInsertRepeat
	err.err = fmt.Errorf(fmt.Sprint(msg...))
	return err
}

// DBAccessError 数据库访问异常
type DBAccessError struct {
	BaseError
}

// Stack return stack info
func (e *DBAccessError) Stack() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", e.err)
}

// NewDBAccessError make a db access error
func NewDBAccessError(msg ...interface{}) error {
	if len(msg) > 0 {
		err := msg[0]
		if err == gorm.ErrRecordNotFound {
			return NewResourcesNotFoundError()
		}
	}
	err := &DBAccessError{}
	err.ErrorCode = ErrorCodeDBAccess
	err.err = pkgerrors.Errorf(fmt.Sprint(msg...))
	return err
}

// GetErrorInfo 获取error信息
func GetErrorInfo(err error) (code int, msg string) {
	if err == nil {
		return ErrorCodeNil, ""
	}
	switch t := err.(type) {
	case *NotLoginError:
		return t.ErrorCode, t.Error()
	case *ResourcesNotPermitError:
		return t.ErrorCode, t.Error()
	case *RequestParamsValueError:
		return t.ErrorCode, t.Error()
	case *validator.InvalidValidationError:
		return ErrorCodeRequestParamsValue, t.Error()
	case *DBAccessError:
		return t.ErrorCode, t.Error()
	case *InsertRepeatError:
		return t.ErrorCode, t.Error()
	case *NotRegisterError:
		return t.ErrorCode, t.Error()
	case *ResourcesNotFoundError:
		return t.ErrorCode, t.Error()
	case *RequestParamsDecodeError:
		return t.ErrorCode, t.Error()
	case *DataAbnormalError:
		return t.ErrorCode, t.Error()
	default:
		return -1, t.Error()
	}
}
