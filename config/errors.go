package config

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Status     		string `json:"status"`
	Description		string `json:"description"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(err error) *ApiError {
	return &ApiError{0, http.StatusInternalServerError, err.Error(), ""}
}

var StartError 			= 	&ApiError{"ERROR", "StartError"}
var StopError 			= 	&ApiError{"ERROR", "StopError"}
var FilesError			= 	&ApiError{"ERROR", "FilesError"}
var ConfigError 		= 	&ApiError{"ERROR", "ConfigError"}
var DeleteError 		= 	&ApiError{"ERROR", "DeleteError"}
var UpdatePasswordError = 	&ApiError{"ERROR", "UpdatePasswordError"}
var ReinstallError 		= 	&ApiError{"ERROR", "ReinstallError"}
var InstallError 		= 	&ApiError{"ERROR", "InstallError"}
var OK 					= 	&ApiError{"OK", ""}
