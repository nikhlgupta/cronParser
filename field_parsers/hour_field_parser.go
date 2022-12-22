package field_parsers

import (
	"cronParser/field_parsers/common"
	"cronParser/field_parsers/time_unit"
	"cronParser/types"
)

type HourFieldParser struct{}

func (m *HourFieldParser) ParseField(fieldExpr string) (*types.FieldOutput, error) {
	return common.CommonFieldParser(fieldExpr, m.FieldTimeUnit())

}

func (m *HourFieldParser) FieldTimeUnit() types.TimeUnit {
	return &time_unit.Hour
}
