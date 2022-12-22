package field_parsers

import (
	"cronParser/field_parsers/common"
	"cronParser/field_parsers/time_unit"
	"cronParser/types"
)

type DayOfMonthFieldParser struct{}

func (m *DayOfMonthFieldParser) ParseField(fieldExpr string) (*types.FieldOutput, error) {
	return common.CommonFieldParser(fieldExpr, m.FieldTimeUnit())
}

func (m *DayOfMonthFieldParser) FieldTimeUnit() types.TimeUnit {
	return &time_unit.DayOfMonth
}

