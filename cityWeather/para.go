package cityWeather

type Para struct {
	Location string
	Lang     string
	Unit     UnitType
}

type UnitType int8

func (u UnitType) String() string {
	switch u {
	case METRIC:
		return "m"
	case IMPERIAL:
		return "i"
	default:
		panic("invalid unit")
	}
}

const (
	_        UnitType = iota
	METRIC            // 公制单位
	IMPERIAL          // 英制单位
)
