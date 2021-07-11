package godeco_test

import (
	"testing"

	"github.com/RonaldCrb/godeco"
)

type NdlTestPair struct {
	dive     godeco.Dive
	expected uint16
}

var NdlTestCases []NdlTestPair = []NdlTestPair{
	{
		expected: 8,
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      147,
		},
	},
	{
		expected: 9999,
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      20,
		},
	},
	{
		expected: 232,
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      33,
		},
	},
	{
		expected: 9999,
		dive: godeco.Dive{
			BottomTime: 300,
			Depth:      0,
		},
	},
	{
		expected: 0,
		dive: godeco.Dive{
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
