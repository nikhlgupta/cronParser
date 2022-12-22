package common

import (
	"cronParser/field_parsers/special_char_handlers"
	"cronParser/types"
	"fmt"
	"strconv"
	"strings"
)

var SpecialCharHandler = map[string]types.SpecialCharHandler{
	"*": &special_char_handlers.AllCharHandler{},
	"-": &special_char_handlers.RangeCharHandler{},
	"?": &special_char_handlers.AnyCharHandler{},
	"/": &special_char_handlers.IncrementCharHandler{},
	",": &special_char_handlers.MultiValCharHandler{},
}

func ParseSpecialChar(fieldExpr string) string {
	specialChars := []string{"/", ",", "-", "*", "?"}

	for _, specialChar := range specialChars {
		if found := strings.Contains(fieldExpr, specialChar); found {
			return specialChar
		}
	}

	return ""
}

func CommonFieldParser(fieldExpr string, timeUnit types.TimeUnit) (*types.FieldOutput, error) {
	specialChar := ParseSpecialChar(fieldExpr)
	//if special char found, parse and return
	if specialChar != "" {
		if handler := SpecialCharHandler[specialChar]; handler != nil {
			return handler.HandleSpecialChar(fieldExpr, timeUnit)
		} else {
			return nil, fmt.Errorf("no handler found for special char %+v for field type %+v", specialChar, timeUnit.GetUnitType())
		}
	}

	//no special char hence field is a single hour value
	if len(fieldExpr) != 1 {
		return nil, fmt.Errorf("invalid value %+v for field type %+v", fieldExpr, timeUnit.GetUnitType())
	}

	hourVal, err := strconv.Atoi(fieldExpr)
	if err != nil {
		return nil, fmt.Errorf("invalid value %+v for field type %+v. Err details %+v", fieldExpr, timeUnit.GetUnitType(), err.Error())
	}

	if hourVal < timeUnit.GetMinVal() || hourVal > timeUnit.GetMaxVal() {
		return nil, fmt.Errorf("invalid value %+v for field type %+v", fieldExpr, timeUnit.GetUnitType())
	}

	return types.NewFieldOutput([]int{hourVal}, timeUnit.GetUnitType()), nil
}
