package errors

import "errors"

var ErrDbIsNotOpened = errors.New("you've tried to close db but it's not opened")
