package special_char_handlers

import (
	"cronParser/types"
	"fmt"
	"strconv"
	"strings"
)

type MultiValCharHandler struct {}

func (a *MultiValCharHandler)HandleSpecialChar(fieldExpr string, timeUnit types.TimeUnit) (*types.FieldOutput, error) {
	timeUnits := make([]int, 0)
	vals := strings.Split(fieldExpr, ",")
	for _, val := range vals {
		valInt, err := strconv.Atoi(val)
		if err != nil || valInt < timeUnit.GetMinVal() || valInt > timeUnit.GetMaxVal() {
			return nil, fmt.Errorf("invalid value %+v for field type %+v", val, timeUnit.GetUnitType())
		}
		timeUnits = append(timeUnits, valInt)
	}
	return types.NewFieldOutput(timeUnits, timeUnit.GetUnitType()), nil
}
