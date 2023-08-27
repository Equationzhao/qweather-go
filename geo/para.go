package geo

type Type int8

func (t Type) String() string {
	switch t {
	case Scenic:
		return "scenic"
	case CSTA:
		return "CSTA"
	case TSTA:
		return "TSTA"
	default:
		panic("invalid type")
	}
}

const (
	_ Type = iota
	Scenic
	CSTA
	TSTA
)

type Para struct {
	Type     Type
	Radius   uint16
	Number   uint16
	Location string
	Adm      string
	Range    string
	Lang     string
	City     string
}
