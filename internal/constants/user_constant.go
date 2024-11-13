package constants

import "time"

var (
	TIME_NOW            = time.Now().Format(time.RFC3339)
	SUCCESS_CREATE_USER = "Success create user"
	FIELD_IS_MISSING    = "Username or Password or Email is empty"
	DATE_TIME_FORMATTER = "2006-01-02 15:04:05"
)
