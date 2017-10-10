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
	ErrUnknown  ErrorCode = "ERR_UNKNOWN"
	ErrArgument ErrorCode = "ERR_ARGUMENT"

	// User management Error Code
	ErrUserPasswrod      ErrorCode = "ERR_USER_PASSWORD"
	ErrUserConflictEmail ErrorCode = "ERR_USER_CONFLICT_EMAIL"
	ErrUserConflictName  ErrorCode = "ERR_USER_CONFLICT_NAME"
	ErrUserNotFound      ErrorCode = "ERR_USER_NOT_FOUND"
	ErrUserSamePassword  ErrorCode = "ERR_USER_SAME_PASSWORD"

	// Project Error Code
	ErrProjectNotFound ErrorCode = "ERR_PROJECT_NOT_FOUND"
	ErrProjectCreate   ErrorCode = "ERR_PROJECT_CREATE"
	// Issue Error Code

)
