package config

import "errors"

const (
	HTTPInvalidBodyMessage          = "Invalid body. Please provide a valid JSON."
	HTTPInvalidPathParameterMessage = "Invalid path parameter. Please provide a valid ID."
	HTTPNotFoundMessage             = "Not found. Please provide a valid ID."
)

var (
	QuoteDoesNotExistError = errors.New("quote does not exist")
)
