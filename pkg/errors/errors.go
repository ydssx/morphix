package errors

import (
	serrors "errors"

	"github.com/pkg/errors"
)

type MultiErr struct {
	errs []error
}

var (
	// ErrNotFound 资源不存在
	ErrNotFound = New("resource not found")
	// ErrAlreadyExists 资源已存在
	ErrAlreadyExists = New("resource already exists")
	// ErrUnauthorized 未授权
	ErrUnauthorized = New("unauthorized")
	// ErrForbidden 禁止访问
	ErrForbidden = New("forbidden")
	// ErrBadRequest 请求错误
	ErrBadRequest = New("bad request")
	// ErrInternal 内部错误
	ErrInternal = New("internal error")
	// ErrTimeout 超时
	ErrTimeout = New("timeout")
	// ErrUnavailable 资源不可用
	ErrUnavailable = New("resource unavailable")
	// ErrLimited 资源已达上限
	ErrLimited = New("resource limited")
	// ErrConflict 冲突
	ErrConflict = New("conflict")
	// ErrCancelled 已取消
	ErrCancelled = New("cancelled")

	New    = serrors.New
	Join   = serrors.Join
	Unwrap = serrors.Unwrap
	Is     = serrors.Is
	As     = serrors.As
	Wrap   = errors.Wrap
	Errorf = errors.Errorf
)
