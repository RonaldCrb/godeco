package test
// **********************************************************************************************
// repetLetter 5 Tests

test('repetLetter 5/147 C', () => {
  const dive = {
    bottomTime: 5,
    depth: 147,
    sit: 250,
    nextDepth: 28,
  };
  expect(repetLetter(dive)).toBe('A');
});

test('repetLetter 198/24 J', () => {
  const dive = {
    bottomTime: 198,
    depth: 24,
    sit: 250,
    nextDepth: 28,
  };
  expect(repetLetter(dive)).toBe('F');
});

test('repetLetter 44/60 H', () => {
  const dive = {
    bottomTime: 44,
    depth: 60,
    sit: 2500,
    nextDepth: 28,
  };
  expect(repetLetter(dive)).toBe(undefined);
});

test('repetLetter 200/124 undefined', () => {
  const dive = {
    bottomTime: 200,
    depth: 124,
    sit: 250,
    nextDepth: 28,
  };
  expect(repetLetter(dive)).toBe(undefined);
});

test('repetLetter 1500/25 undefined', () => {
  const dive = {
    bottomTime: 1500,
    depth: 25,
    sit: 250,
    nextDepth: 28,
  };
  expect(repetLetter(dive)).toBe(undefined);
});
