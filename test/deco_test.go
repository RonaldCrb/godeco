package test
// ****************************************************************************
// decoDive 5 tests

test('decoDive 99/135 decoObject', () => {
  const dive = {
    bottomTime: 135,
    depth: 99,
    sit: 250,
    nextDepth: 69,
  };
  expect(decoDive(dive)).toMatchObject({
    minTime: 121,
    maxTime: 150,
    ttfs: '1:40',
    airTAT: '538:20',
    o2TAT: '183:40',
    o2cp: 5,
    repetLetter: 'N/A',
    surDo2Recom: true,
    surDo2Req: true,
    strictlySurDo2: true,
    exceptionalExposure: true,
    airDecoStops: [{ depth: 20, time: 461 }, { depth: 30, time: 46 }, { depth: 40, time: 26 }, { depth: 50, time: 3 }],
    o2decoStops: [{ depth: 20, time: 109 }, { depth: 30, time: 23 }, { depth: 40, time: 26 }, { depth: 50, time: 3 }],
  });
});

test('decoDive 1000/135 => No Table Matched', () => {
  const dive = {
    bottomTime: 135,
    depth: 1000,
    sit: 250,
    nextDepth: 69,
  };
  expect(decoDive(dive)).toBe('No Table Found');
});

test('decoDive 100/1305 => no decoObject Matched', () => {
  const dive = {
    bottomTime: 1305,
    depth: 100,
    sit: 250,
    nextDepth: 69,
  };
  expect(decoDive(dive)).toBe('no deco profile Found');
});

test('decoDive 69/130 => decoObject', () => {
  const dive = {
    bottomTime: 130,
    depth: 69,
    sit: 250,
    nextDepth: 69,
  };
  expect(decoDive(dive)).toMatchObject({
    minTime: 121,
    maxTime: 130,
    ttfs: '1:20',
    airTAT: '169:20',
    o2TAT: '58:20',
    o2cp: 2,
    repetLetter: 'Z',
    surDo2Recom: true,
    exceptionalExposure: true,
    airDecoStops: [{ depth: 20, time: 167 }],
    o2decoStops: [{ depth: 20, time: 51 }],
  });
});

test('decoDive 61/125 => decoObject', () => {
  const dive = {
    bottomTime: 125,
    depth: 61,
    sit: 250,
    nextDepth: 69,
  };
  expect(decoDive(dive)).toMatchObject({
    minTime: 121,
    maxTime: 130,
    ttfs: '1:20',
    airTAT: '169:20',
    o2TAT: '58:20',
    o2cp: 2,
    repetLetter: 'Z',
    surDo2Recom: true,
    exceptionalExposure: true,
    airDecoStops: [{ depth: 20, time: 167 }],
    o2decoStops: [{ depth: 20, time: 51 }],
  });
});