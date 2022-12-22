package special_char_handlers

import (
	"cronParser/types"
	"fmt"
)

type AnyCharHandler struct {}

func (a *AnyCharHandler)HandleSpecialChar(fieldExpr string, timeUnit types.TimeUnit) (*types.FieldOutput, error) {
	if len(fieldExpr) != 1 {
		return nil, fmt.Errorf("invalid expression for unit type %+v", timeUnit.GetUnitType())
	}
	return types.NewFieldOutput(make([]int, 0), timeUnit.GetUnitType()), nil
}
