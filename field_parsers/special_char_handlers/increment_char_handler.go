package special_char_handlers

import (
	"cronParser/types"
	"fmt"
	"strconv"
	"strings"
)

type IncrementCharHandler struct{}

func (a *IncrementCharHandler) HandleSpecialChar(fieldExpr string, timeUnit types.TimeUnit) (*types.FieldOutput, error) {

	fieldExpr = strings.ReplaceAll(fieldExpr, "*", fmt.Sprint(timeUnit.GetMinVal()))
	intervalRange := strings.Split(fieldExpr, "/")
	timeUnits := make([]int, 0)

	if len(intervalRange) != 2 {
		return nil, fmt.Errorf("invalid expression for interval. %+v", fieldExpr)
	}

	intervalStart, err := strconv.Atoi(intervalRange[0])
	if err != nil || intervalStart < timeUnit.GetMinVal() || intervalStart > timeUnit.GetMaxVal() {
		return nil, fmt.Errorf("invalid increment start variable %+v ffor unt type %+v", intervalRange[0], timeUnit.GetUnitType())
	}
	intervalIncr, err := strconv.Atoi(intervalRange[1])
	if err != nil {
		return nil, fmt.Errorf("invalid range ncrement variable. not able to map %+v to int", intervalRange[1])
	}

	if intervalIncr <= 0 || intervalIncr < timeUnit.GetMinVal() || intervalIncr > timeUnit.GetMaxVal() {
		return nil, fmt.Errorf("invalid range ncrement variable %+v for unt type %+v", intervalIncr, timeUnit.GetUnitType())
	}

	for i := intervalStart; i <= timeUnit.GetMaxVal(); i += intervalIncr {
		timeUnits = append(timeUnits, i)
	}
	return types.NewFieldOutput(timeUnits, timeUnit.GetUnitType()), nil
}
