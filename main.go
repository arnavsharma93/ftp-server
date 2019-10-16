package main

import (
	"fmt"
	"net/http"

	"github.com/anshal21/ftp-server/operator"
	ftp "github.com/anshal21/ftp-server/v1"
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

type handler struct {
	operator operator.Operator
}

func NewHandler() handler {
	return handler{
		operator: ftp.New(),
	}
}

func (h handler) ftpHandler(w http.ResponseWriter, req *http.Request) {
}
