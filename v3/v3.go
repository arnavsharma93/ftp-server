package v3

import (
	"github.com/anshal21/ftp-server/operator"
	"github.com/anshal21/ftp-server/v2"
)
type v3 struct {
	operator.Operator

	fileCache map[string][]byte
	cacheHit  int
	cacheMiss int
}

func New() operator.Operator {
	return v3{
		Operator:  v2.New(),
		fileCache: make(map[string][]byte),
	}
}

