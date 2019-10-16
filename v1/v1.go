package v1
import (
	"io"
)

type v1 struct{}
func New() operator.Operator {
	return v1{}
}

func (v v1) List(dirName string) (io.Reader, error) {
}

func (v v1) Delete(dirName string) error {

}

func (v v1) Get(fileName string) (io.Reader, error) {
	return nil, nil
}
