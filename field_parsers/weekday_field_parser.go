package field_parsers

import (
	"cronParser/field_parsers/common"
	"cronParser/field_parsers/time_unit"
	"cronParser/types"
	"strings"
)

var weekDayLiteralToInt = map[string]string {
	"MON": "1",
	"TUE": "2",
	"WED": "3",
	"THU": "4",
	"FRI": "5",
	"SAT": "6",
	"SUN": "7",

}
type WeekDayFieldParser struct{}

func (m *WeekDayFieldParser) ParseField(fieldExpr string) (*types.FieldOutput, error) {
	//first sanitise the expr for any literals for weekday and replace them with integers values
	for weekLiteral, monthIntVal := range weekDayLiteralToInt {
		fieldExpr = strings.ReplaceAll(fieldExpr, weekLiteral, monthIntVal)
	}
	return common.CommonFieldParser(fieldExpr, m.FieldTimeUnit())
}

func (m *WeekDayFieldParser) FieldTimeUnit() types.TimeUnit {
	return &time_unit.DayOfWeek
}
