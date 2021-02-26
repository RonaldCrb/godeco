package test

import (
	"testing"

	deco "github.com/RonaldCrb/diving-decompression-go"
)

type RlTestPair struct {
	dive     deco.DivePlan
	expected string
}

var RlTestCases []RlTestPair = []RlTestPair{
	{
		expected: "A",
		dive: deco.DivePlan{
			BottomTime: 5,
			Depth:      147,
			SIT:        250,
			NextDepth:  28,
		},
	},
	{
		expected: "F",
		dive: deco.DivePlan{
			BottomTime: 198,
			Depth:      24,
			SIT:        250,
			NextDepth:  28,
		},
	},
	{
		expected: "",
		dive: deco.DivePlan{
			BottomTime: 44,
			Depth:      60,
			SIT:        2500,
			NextDepth:  28,
		},
	},
	{
		expected: "",
		dive: deco.DivePlan{
			BottomTime: 200,
			Depth:      124,
			SIT:        350,
			NextDepth:  128,
		},
	},
	{
		expected: "",
		dive: deco.DivePlan{
			BottomTime: 1500,
			Depth:      25,
			SIT:        250,
			NextDepth:  28,
		},
	},
}

func TestRepetLetter(t *testing.T) {
	for _, d := range RlTestCases {
		di, err := d.dive.RepetLetter()
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
