package main

import (
	"fmt"
)

const (
	_operationURLParameterKey = "operation"
	_dirURLParamterKey        = "dir"

	_FTPOperationList   = "ls"
	_FTPOperationDelete = "del"
	_FTPOperationGet    = "get"

	_serverPort = ":8000"
)

var (
	helpMessage = fmt.Sprintf(`Please provide one the following option
  - %v: list down the files in a directory
  - %v: delete files in a directory
  - %v: returns a file
`, _FTPOperationList, _FTPOperationDelete, _FTPOperationGet)
)

