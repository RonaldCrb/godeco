package test

import (
	"testing"

	deco "github.com/RonaldCrb/diving-decompression-go"
)

type glTestPair struct {
	dive     deco.Dive
	expected string
}

var glTestCases []glTestPair = []glTestPair{
	{
		expected: "C",
		dive: deco.Dive{
			BottomTime: 5,
			Depth:      147,
		},
	},
	{
		expected: "this dive is out of the depth range for no-decompression air dives",
		dive: deco.Dive{
			BottomTime: 24,
			Depth:      198,
		},
	},
	{
		expected: "H",
		dive: deco.Dive{
			BottomTime: 44,
			Depth:      60,
		},
	},
	{
		expected: "this dive is out of the time range for no-decompression air dives",
		dive: deco.Dive{
			BottomTime: 200,
			Depth:      124,
		},
	},
	{
		expected: "this dive is out of the time range for no-decompression air dives",
		dive: deco.Dive{
			BottomTime: 1500,
			Depth:      25,
		},
	},
}

func TestGroupLetter(t *testing.T) {
	for _, d := range glTestCases {
		di, err := d.dive.GroupLetter()
		if err != nil {
			t.Errorf("%s", err)
		}
		if di != d.expected {
			t.Errorf("Failed ! got %v want %s", di, d.expected)
		} else {
			t.Logf("Success !")
		}
	}
}
