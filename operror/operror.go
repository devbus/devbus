package operror

type OpError struct {
	Code ErrorCode
}

func (e OpError) Error() string {
	return string(e.Code)
}

type ErrorCode string

const (
	// General Error Code
	ErrUnknown ErrorCode = "ERR_UNKNOWN"
	// Project Error Code
	ErrProjectNotFound ErrorCode = "ERR_PROJECT_NOT_FOUND"
	ErrProjectCreate   ErrorCode = "ERR_PROJECT_CREATE"
	// Issue Error Code

)
