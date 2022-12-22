package special_char_handlers

import (
	"cronParser/types"
	"fmt"
)

type AllCharHandler struct {}

func (a *AllCharHandler)HandleSpecialChar(fieldExpr string, unitType types.TimeUnit) (*types.FieldOutput, error) {
	if len(fieldExpr) != 1 {
		return nil, fmt.Errorf("invalid expression %+v for unit Type %+v", fieldExpr, unitType.GetUnitType())
	}

	allTimeUnits := make([]int, 0)
	for i:=unitType.GetMinVal(); i <= unitType.GetMaxVal(); i++ {
		allTimeUnits = append(allTimeUnits, i)
	}

	return types.NewFieldOutput(allTimeUnits, unitType.GetUnitType()), nil
}
