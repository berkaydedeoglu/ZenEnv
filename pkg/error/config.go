package errors

import "errors"

var (
	ErrFileNotFoundInDefaults = errors.New("we cannot find any config files in any default directories")
	ErrSpecifiedFileNotFound  = errors.New("there is no configuration file you have specified")
)
