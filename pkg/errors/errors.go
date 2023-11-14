package errors

import (
	serrors "errors"

	"github.com/pkg/errors"
)

type MultiErr struct {
	errs []error
}

var (
	New    = serrors.New
	Join   = serrors.Join
	Unwrap = serrors.Unwrap
	Is     = serrors.Is
	As     = serrors.As
	Wrap   = errors.Wrap
	Errorf = errors.Errorf
)
