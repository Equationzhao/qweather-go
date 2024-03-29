package qweather

import "testing"

func TestUnitType_String(t *testing.T) {
	units := []UnitType{METRIC, IMPERIAL, 0}
	for _, unit := range units {
		t.Log(unit.String())
	}
}
