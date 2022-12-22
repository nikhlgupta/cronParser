package types

type ParsedExpression struct {
	timeFields []*FieldOutput
	command    interface{}
}

func NewParsedExpression(timeFields []*FieldOutput, command interface{}) *ParsedExpression {
	return &ParsedExpression{timeFields: timeFields, command: command}
}

func (p *ParsedExpression) Command() interface{} {
	return p.command
}

func (p *ParsedExpression) TimeFields() []*FieldOutput {
	return p.timeFields
}

type FieldOutput struct {
	timeUnits []int
	unitType  UnitType
}

func NewFieldOutput(timeUnits []int, unitType UnitType) *FieldOutput {
	return &FieldOutput{timeUnits: timeUnits, unitType: unitType}
}

func (f *FieldOutput) TimeUnits() []int {
	return f.timeUnits
}

func (f *FieldOutput) UnitType() UnitType {
	return f.unitType
}

type UnitType string

const (
	Minute   UnitType = "minute"
	Hour              = "hour"
	WeekDay           = "weekday"
	MonthDay          = "monthday"
	Month             = "month"
)
