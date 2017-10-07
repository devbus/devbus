package common

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
	// User management Error Code
	ErrPasswrod         ErrorCode = "ERR_PASSWORD"
	ErrConflictEmail    ErrorCode = "ERR_CONFLICT_EMAIL"
	ErrConflictUserName ErrorCode = "ERR_CONFLICT_USER_NAME"
	// Project Error Code
	ErrProjectNotFound ErrorCode = "ERR_PROJECT_NOT_FOUND"
	ErrProjectCreate   ErrorCode = "ERR_PROJECT_CREATE"
	// Issue Error Code

)
