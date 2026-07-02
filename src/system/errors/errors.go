package errors

import (
	"fmt"
)

type ErrorCode string

type FolionError struct {
	Code    ErrorCode
	Severe  bool
	Message string
}

func NewFolionError(_code ErrorCode, _message string, _severe bool) *FolionError {
	return &FolionError{
		Code:    _code,
		Message: _message,
		Severe:  _severe,
	}
}

func (err *FolionError) Error() string {
	return fmt.Sprintf("Error [%s]", err.Code)
}

const (
	CodeServiceBootFailed ErrorCode = "SERVICE_BOOT_FAILED"
	CodeTCPStartFailed    ErrorCode = "TCP_START_FAILED"
	CodeConfigLoadFailed  ErrorCode = "CONFIG_LOAD_FAILED"
)

var (
	ErrServiceBootFailed = NewFolionError(CodeServiceBootFailed, "Failed to boot Folion service", true)
	ErrServerStartFailed = NewFolionError(CodeTCPStartFailed, "Failed to start server", true)
	ErrConfigLoadFailed  = NewFolionError(CodeConfigLoadFailed, "Failed to load configuration file", true)
)
