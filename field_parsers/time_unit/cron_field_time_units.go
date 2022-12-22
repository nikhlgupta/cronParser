package time_unit

import "cronParser/types"

var Hour = hour{}
var Minute = minute{}
var Month = month{}
var DayOfMonth = dayOfMonth{}
var DayOfWeek = dayOfWeek{}

//hour
type hour struct{}

func (h *hour) GetMaxVal() int {
	return 23
}
func (h *hour) GetMinVal() int {
	return 0
}
func (h *hour) GetUnitType() types.UnitType {
	return types.Hour
}

//minute
type minute struct{}

func (h *minute) GetMaxVal() int {
	return 59
}
func (h *minute) GetMinVal() int {
	return 0
}
func (h *minute) GetUnitType() types.UnitType {
	return types.Minute
}

//month
type month struct{}

func (h *month) GetMaxVal() int {
	return 12
}
func (h *month) GetMinVal() int {
	return 1
}
func (h *month) GetUnitType() types.UnitType {
	return types.Month
}

//day of month
type dayOfMonth struct{}

func (d *dayOfMonth) GetMaxVal() int {
	return 31
}
func (d *dayOfMonth) GetMinVal() int {
	return 1
}
func (d *dayOfMonth) GetUnitType() types.UnitType {
	return types.MonthDay
}

//day of the week
type dayOfWeek struct{}

func (d *dayOfWeek) GetMaxVal() int {
	return 7
}
func (d *dayOfWeek) GetMinVal() int {
	return 1
}
func (d *dayOfWeek) GetUnitType() types.UnitType {
	return types.WeekDay
}




