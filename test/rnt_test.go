package test
// **********************************************************************************************
// residual nitrogen time 5 Tests

test('residual nitrogen time 5/147 C', () => {
  const dive = {
    bottomTime: 5,
    depth: 147,
    sit: 250,
    nextDepth: 28,
  };
  expect(residualNitrogenTime(dive)).toBe(18);
});

test('residual nitrogen time 198/24 J', () => {
  const dive = {
    bottomTime: 198,
    depth: 24,
    sit: 250,
    nextDepth: 28,
  };
  expect(residualNitrogenTime(dive)).toBe(77);
});

test('residual nitrogen time 44/60 H', () => {
  const dive = {
    bottomTime: 44,
    depth: 60,
    sit: 22,
    nextDepth: 190,
  };
  expect(residualNitrogenTime(dive)).toBe(14);
});

test('residual nitrogen time 200/124 undefined', () => {
  const dive = {
    bottomTime: 200,
    depth: 124,
    sit: 350,
    nextDepth: 128,
  };
  expect(residualNitrogenTime(dive)).toBe(undefined);
});

test('residual nitrogen time 1500/25 undefined', () => {
  const dive = {
    bottomTime: 150,
    depth: 25,
    sit: 250,
    nextDepth: 69,
  };
  expect(residualNitrogenTime(dive)).toBe(25);
});
