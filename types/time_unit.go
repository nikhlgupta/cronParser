package types

type TimeUnit interface {
	GetMaxVal() int
	GetMinVal() int
	GetUnitType() UnitType
}
