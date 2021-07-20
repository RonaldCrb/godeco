package godeco_test

import (
	"testing"

	"github.com/RonaldCrb/godeco"
)

type TestDivePlan struct {
	dive        godeco.Dive
	expectedNDL uint16
	expectedGL  string
	expectedRL  string
	expectedRNT uint16
}

var TestDivePlans []TestDivePlan

func TestDivePlanNoDecompressionLimit(t *testing.T) {
	for _, d := range TestDivePlans {
		di, err := d.dive.NoDecompressionLimit()
		if err != nil {
			t.Errorf("%s", err)
		}
		if di != d.expectedNDL {
			t.Errorf("Failed [%d/%d] got %d expected %d", d.dive.Depth, d.dive.BottomTime, di, d.expectedNDL)
		} else {
			t.Logf("Success !")
		}
	}
}
func TestGroupLetter(t *testing.T) {

}
func TestRepetLetter(t *testing.T) {

}
func TestResidualNitrogenTime(t *testing.T) {

}
