package error_with_codes

import "strconv"

type ErrorCode int

func (e ErrorCode) Int() int {
	return int(e)
}

func (e ErrorCode) String() string {
	return strconv.Itoa(e.Int())
}

// cast error
const (
	_ ErrorCode = iota + 0
	CodeFailedToCast
)

var (
	ErrorFailedToCast = NewError("failed to cast object", CodeFailedToCast)
)

// cfg
const (
	_ ErrorCode = iota + 99
	CodeFailedToFindConfig
	CodeFailedToReadConfig
)

var (
	ErrorFailedToFindConfig = NewError("failed to find config", CodeFailedToFindConfig)
	ErrorFailedToReadConfig = NewError("failed to read config", CodeFailedToReadConfig)
)

// adapter
const (
	_ ErrorCode = iota + 199
	CodeFailedToCreateUrl
	CodeFailedToGetUrlList
	CodeFailedToGetUrlByCode
)

var (
	ErrorFailedToCreateUrl    = NewError("failed to create url", CodeFailedToCreateUrl)
	ErrorFailedToGetUrlList   = NewError("failed to get url list", CodeFailedToGetUrlList)
	ErrorFailedToGetUrlByCode = NewError("failed to get url by code", CodeFailedToGetUrlByCode)
)

// handler
const (
	_ ErrorCode = iota + 299
	CodeMethodNotAllowed
	CodeFailedToUnmarshal
	CodeFailedToValidate
)

var (
	ErrorMethodNotAllowed  = NewError("method not allowed", CodeMethodNotAllowed)
	ErrorFailedToUnmarshal = NewError("failed to unmarshal", CodeFailedToUnmarshal)
	ErrorFailedToValidate  = NewError("failed to validate", CodeFailedToValidate)
)

// unit-test
const (
	_ ErrorCode = iota + 399
	CodeInvalidUrlCode
)

var (
	ErrorInvalidUrlCode = NewError("invalid url code", CodeInvalidUrlCode)
)
