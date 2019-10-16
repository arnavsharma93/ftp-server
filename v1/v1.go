package v1
type v1 struct{}
func New() operator.Operator {
	return v1{}
}
