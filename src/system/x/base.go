package x

import (
	"fmt"
	"folion-server/src/infrastructure/logging"
	"log/slog"
	"os"
)

type ErrorCode string

type FolionError struct {
	Code    ErrorCode
	Severe  bool
	Desc    string
	message string
	Err     error
}

func NewFolionError(_code ErrorCode, _desc string, _severe bool) *FolionError {
	return &FolionError{
		Code:   _code,
		Desc:   _desc,
		Severe: _severe,
	}
}

func (err *FolionError) Error() string {
	return fmt.Sprintf("Error [%s] \n\t Desc:\t%s \n\t %s", err.Code, err.Desc, err.message)
}

func (err *FolionError) Wrap(_error error) *FolionError {
	err.Err = _error
	return err
}

func (err *FolionError) Message(_message string) *FolionError {
	err.message = _message
	return err
}

func (err *FolionError) Unwrap() error {
	return err.Err
}

func Cautious(callable func() error) {
	err := callable()
	if err != nil {
		if logging.GetInstance() != nil {
			slog.Error(err.Error())
		} else {
			os.Stdout.WriteString(err.Error())
		}

		os.Exit(1)
	}
}
