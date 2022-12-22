package types

type SpecialCharHandler interface {
	HandleSpecialChar(fieldExpr string, timeUnit TimeUnit) (*FieldOutput, error)
}
