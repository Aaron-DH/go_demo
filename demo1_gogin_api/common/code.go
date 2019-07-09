package common

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	DBError				= &Errno{Code: 10003, Message: "Error occurred while operating the database."}
	ParamsError			= &Errno{Code: 10004, Message: "Params is not correct."}

	// user errors
	ErrUserNotFound 	= &Errno{Code: 20101, Message: "The user was not found."}
	ErrPasswdIncorrect	= &Errno{Code: 20102, Message: "The password is incorrect."}
)
