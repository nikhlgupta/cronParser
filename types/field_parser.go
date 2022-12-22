package types

type FieldParser interface {
	ParseField(fieldExpr string) (*FieldOutput, error)
	FieldTimeUnit() TimeUnit
}
