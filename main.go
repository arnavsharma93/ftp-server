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

	operation, ok := req.URL.Query()[_operationURLParameterKey]
	if !ok {
		w.Write([]byte(helpMessage))
		return
	}

	dir, ok := req.URL.Query()[_dirURLParamterKey]
	if !ok {
		w.Write([]byte("Missing required parameter"))
		return
	}

	//fmt.Println(operation, dir)

	response := make([]byte, 1024*1024)

	switch operation[0] {
	case _FTPOperationList:
		reader, err := h.operator.List(dir[0])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		n, _ := reader.Read(response)
		w.Write(response[:n])

	case _FTPOperationDelete:
		err := h.operator.Delete(dir[0])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("RemovedDirectory"))

	case _FTPOperationGet:
		resp, err := h.operator.Get(dir[0])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		n, _ := resp.Read(response)
		w.Write(response[:n])
	}
}

// Select
}
