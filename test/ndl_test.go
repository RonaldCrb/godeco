package test

import (
	"testing"

	deco "github.com/RonaldCrb/diving-decompression-go"
)

func TestNoDecompressionLimit(t *testing.T) {
	dive := deco.Dive{
		BottomTime: 300,
		Depth:      147,
	}
	expectedNdl := 8

	output := dive.NoDecompressionLimit()

	if expectedNdl != output {
		t.Errorf("Failed ! got %v want %c", output, expectedNdl)
	} else {
		t.Logf("Success !")
	}
}

/*
// **************************************************************************************************
// noDecompressionLimit 5 Tests
test('noDecompressionLimit 147', () => {

  expect(noDecompressionLimit(dive)).toBe(8);
});

test('noDecompressionLimit 20', () => {
  const dive = {
    bottomTime: 300,
    depth: 20,
  };
  expect(noDecompressionLimit(dive)).toBe('unlimited');
});

test('noDecompressionLimit 33', () => {
  const dive = {
    bottomTime: 300,
    depth: 33,
  };
  expect(noDecompressionLimit(dive)).toBe(232);
});

test('noDecompressionLimit 98 undefined', () => {
  const dive = {
    bottomTime: 300,
    depth: -1,
  };
  expect(noDecompressionLimit(dive)).toBe(undefined);
});

test('noDecompressionLimit 200 undefined', () => {
  const dive = {
    bottomTime: 300,
    depth: 200,
  };
  expect(noDecompressionLimit(dive)).toBe(undefined);
});
*/
