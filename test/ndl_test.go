package test

import (
	"testing"

	deco "github.com/RonaldCrb/diving-decompression-go"
)

type NdlTestPair struct {
	dive     deco.Dive
	expected uint16
}

var NdlTestCases []NdlTestPair = []NdlTestPair{
	{
		expected: 8,
		dive: deco.Dive{
			BottomTime: 300,
			Depth:      147,
		},
	},
	{
		expected: 9999,
		dive: deco.Dive{
			BottomTime: 300,
			Depth:      20,
		},
	},
	{
		expected: 232,
		dive: deco.Dive{
			BottomTime: 300,
			Depth:      33,
		},
	},
	{
		expected: 9999,
		dive: deco.Dive{
			BottomTime: 300,
			Depth:      0,
		},
	},
	{
		expected: 0,
		dive: deco.Dive{
			BottomTime: 300,
			Depth:      300,
		},
	},
}

func TestNoDecompressionLimit(t *testing.T) {
	for _, d := range NdlTestCases {
		di, err := d.dive.NoDecompressionLimit()
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
