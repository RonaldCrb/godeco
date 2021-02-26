package test

import (
	"testing"

	deco "github.com/RonaldCrb/diving-decompression-go"
)

type RntTestPair struct {
	dive     deco.DivePlan
	expected uint16
}

var RntTestCases []RntTestPair = []RntTestPair{
	{
		expected: 18,
		dive: deco.DivePlan{
			BottomTime: 5,
			Depth:      147,
			SIT:        250,
			NextDepth:  28,
		},
	},
	{
		expected: 77,
		dive: deco.DivePlan{
			BottomTime: 198,
			Depth:      24,
			SIT:        250,
			NextDepth:  28,
		},
	},
	{
		expected: 14,
		dive: deco.DivePlan{
			BottomTime: 44,
			Depth:      60,
			SIT:        22,
			NextDepth:  190,
		},
	},
	{
		expected: 0,
		dive: deco.DivePlan{
			BottomTime: 200,
			Depth:      124,
			SIT:        350,
			NextDepth:  128,
		},
	},
	{
		expected: 25,
		dive: deco.DivePlan{
			BottomTime: 150,
			Depth:      25,
			SIT:        250,
			NextDepth:  69,
		},
	},
}

func TestResidualNitrogenTime(t *testing.T) {
	for _, d := range RntTestCases {
		di, err := d.dive.ResidualNitrogenTime()
		if err != nil {
			t.Errorf("%s", err)
		}
		if di != d.expected {
			t.Errorf("Failed ! got %v want %c", di, d.expected)
		} else {
			t.Logf("Success !")
		}
	}
}
