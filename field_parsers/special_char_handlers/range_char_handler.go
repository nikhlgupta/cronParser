package special_char_handlers

import (
	"cronParser/types"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type RangeCharHandler struct {}

func (a *RangeCharHandler)HandleSpecialChar(fieldExpr string, timeUnit types.TimeUnit) (*types.FieldOutput, error) {
	rangeFields := strings.Split(fieldExpr, "-")
	timeUnits := make([]int, 0)
	if len(rangeFields) != 2 {
		return nil, errors.New("invalid minute range expression")
	}
	rangeStart, err := strconv.Atoi(rangeFields[0])
	if err != nil {
		return nil, fmt.Errorf("invalid range start variable. not able to map %+v to int", rangeFields[1])
	}

	rangeEnd, err := strconv.Atoi(rangeFields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid range start variable. not able to map %+v to int", rangeFields[1])
	}

	if rangeStart < timeUnit.GetMinVal() || rangeStart > timeUnit.GetMaxVal() {
		return nil, fmt.Errorf("invalid range value %+v for minute", rangeStart)
	}

	if rangeEnd < timeUnit.GetMinVal() || rangeEnd > timeUnit.GetMaxVal() {
		return nil, fmt.Errorf("invalid range value %+v for minute", rangeEnd)
	}

	if rangeStart >= rangeEnd {
		return nil, fmt.Errorf("invalid range start %+v and end value %+v for unit type %+v. strat value cant be bigger than end value", rangeStart, rangeEnd, timeUnit.GetUnitType())
	}
	for i := rangeStart; i <= rangeEnd; i++ {
		timeUnits = append(timeUnits, i)
	}
	return types.NewFieldOutput(timeUnits, timeUnit.GetUnitType()), nil
}
