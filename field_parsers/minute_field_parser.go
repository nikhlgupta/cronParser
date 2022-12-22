package field_parsers

import (
	"cronParser/field_parsers/common"
	"cronParser/field_parsers/time_unit"
	"cronParser/types"
)

type MinuteFieldParser struct{}

func (m *MinuteFieldParser) ParseField(fieldExpr string) (*types.FieldOutput, error) {
	return common.CommonFieldParser(fieldExpr, m.FieldTimeUnit())
}

func (m *MinuteFieldParser) FieldTimeUnit() types.TimeUnit {
	return &time_unit.Minute
}
