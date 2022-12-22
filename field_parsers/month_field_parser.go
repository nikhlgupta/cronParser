package field_parsers

import (
	"cronParser/field_parsers/common"
	"cronParser/field_parsers/time_unit"
	"cronParser/types"
	"strings"
)

var monthLiteralToInt = map[string]string {
	"JAN": "1",
	"FEB": "2",
	"MAR": "3",
	"APR": "4",
	"MAY": "5",
	"JUN": "6",
	"JUL": "7",
	"AUG": "8",
	"SEP": "9",
	"OCT": "10",
	"NOV": "11",
	"DEC": "12",

}
type MonthFieldParser struct{}

func (m *MonthFieldParser) ParseField(fieldExpr string) (*types.FieldOutput, error) {
	//first sanitise the expr for any literals for month and replace them with integers values
	for monthLiteral, monthIntVal := range monthLiteralToInt {
		fieldExpr = strings.ReplaceAll(fieldExpr, monthLiteral, monthIntVal)
	}
	//now common parser can be called which deals with in values only
	return common.CommonFieldParser(fieldExpr, m.FieldTimeUnit())
}

func (m *MonthFieldParser) FieldTimeUnit() types.TimeUnit {
	return &time_unit.Month
}
