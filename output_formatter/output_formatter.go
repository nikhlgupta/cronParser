package output_formatter

import (
	"cronParser/types"
	"fmt"
	"strings"
)

var outPutFieldPrefix = map[types.UnitType]string{
	types.Minute:   "minute        ",
	types.Hour:     "hour          ",
	types.WeekDay:  "day of week   ",
	types.MonthDay: "day of month  ",
	types.Month:    "month         ",
}

func FormatOutput(parsedExpression *types.ParsedExpression) [][]string {
	formattedOutput := make([][]string, 6)

	for i, fielOp := range parsedExpression.TimeFields() {
		builder := strings.Builder{}
		formattedOutput[i] = make([]string, 0)
		//append the prefix
		builder.WriteString(outPutFieldPrefix[fielOp.UnitType()])
		//append the values now spaced out
		if fielOp.TimeUnits() != nil {
			for _, val := range fielOp.TimeUnits() {
				builder.WriteString(fmt.Sprintf("%+v ", val))
			}
		}

		formattedOutput[i] = append(formattedOutput[i], builder.String())
		//print the line
		fmt.Println(builder.String())

	}
	cs :=  "command       "+parsedExpression.Command().(string)
	fmt.Println(cs)
	formattedOutput[5] = append(formattedOutput[5], cs)

	return formattedOutput
}
