package expression_parser

import (
	"cronParser/field_parsers"
	"cronParser/types"
	"fmt"
	"strings"
)

const SEPARATOR = " "
const COMMAND_INDEX = 5

var inputIndexMapper = map[int]types.FieldParser{
	0: &field_parsers.MinuteFieldParser{},
	1: &field_parsers.HourFieldParser{},
	2: &field_parsers.DayOfMonthFieldParser{},
	3: &field_parsers.MonthFieldParser{},
	4: &field_parsers.WeekDayFieldParser{},
}

func ParseCronExpression(expr string) (*types.ParsedExpression, error) {
	fields := strings.Split(expr, SEPARATOR)

	//   0       1          2          3          4          5
	//<minute> <hour> <day-of-month> <month> <day-of-week> <command>
	if len(fields) != 6 {
		return nil, fmt.Errorf("invalid expression string %+v", expr)
	}

	fieldOutputs := make([]*types.FieldOutput, 0)

	for fieldIndex, fieldExpr := range fields {
		parser, found := inputIndexMapper[fieldIndex]
		if fieldIndex == COMMAND_INDEX {
			continue
		}
		if !found || parser == nil {
			return nil, fmt.Errorf("no field parser found for field expr %+v", fieldExpr)
		}

		fieldOutput, err := parser.ParseField(fieldExpr)
		if err != nil {
			return nil, err
		}
		if fieldOutput != nil {
			fieldOutputs = append(fieldOutputs, fieldOutput)
		}
	}

	return types.NewParsedExpression(fieldOutputs, fields[COMMAND_INDEX]), nil
}
