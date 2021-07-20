package godeco_test

import (
	"testing"

	"github.com/RonaldCrb/godeco"
)

type TestDive struct {
	dive        godeco.Dive
	expectedNDL uint16
	expectedGL  string
}

var TestDives []TestDive = []TestDive{
	{
		expectedNDL: 8,
		expectedGL:  "this dive is out of the time range for no-decompression air dives",
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      147,
		},
	},
	{
		expectedNDL: 9999,
		expectedGL:  "J",
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      20,
		},
	},
	{
		expectedNDL: 232,
		expectedGL:  "this dive is out of the time range for no-decompression air dives",
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      33,
		},
	},
	{
		expectedNDL: 9999,
		expectedGL:  "E",
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      0,
		},
	},
	{
		expectedNDL: 0,
		expectedGL:  "this dive is out of the depth range for no-decompression air dives",
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      300,
		},
	},
	{
		expectedGL:  "C",
		expectedNDL: 8,
		dive: godeco.Dive{
			BottomTime: 5,
			Depth:      147,
		},
	},
	{
		expectedGL: "this dive is out of the depth range for no-decompression air dives",
		dive: godeco.Dive{
			BottomTime: 24,
			Depth:      198,
		},
	},
	{
		expectedGL:  "H",
		expectedNDL: 63,
		dive: godeco.Dive{
			BottomTime: 44,
			Depth:      60,
		},
	},
	{
		expectedGL:  "this dive is out of the time range for no-decompression air dives",
		expectedNDL: 12,
		dive: godeco.Dive{
			BottomTime: 200,
			Depth:      124,
		},
	},
	{
		expectedGL:  "this dive is out of the time range for no-decompression air dives",
		expectedNDL: 1102,
		dive: godeco.Dive{
			BottomTime: 1500,
			Depth:      25,
		},
	},
}

func TestDiveNoDecompressionLimit(t *testing.T) {
	for _, d := range TestDives {
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

func TestDiveGroupLetter(t *testing.T) {
	for _, d := range TestDives {
		di, err := d.dive.GroupLetter()
		if err != nil {
			t.Errorf("%s", err)
		}
		if di != d.expectedGL {
			t.Errorf("Failed [%d/%d] got %s expected %s", d.dive.Depth, d.dive.BottomTime, di, d.expectedGL)
		} else {
			t.Logf("Success !")
		}
	}
}
