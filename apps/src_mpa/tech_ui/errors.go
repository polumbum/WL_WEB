package ui

import "errors"

var (
	ErrNilPointer = errors.New("nil pointer")
	ErrNoTCamps   = errors.New("no available training camps")
	ErrNoComps    = errors.New("no available competitions")
	ErrNoSm       = errors.New("no available sportsmen")
	ErrNoCoaches  = errors.New("no available coaches")
	ErrRole       = errors.New("wrong role")
)
