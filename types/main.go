// table for no decompression limit and group letter
type Group struct {
  groupLetter string
  minTime i32
  maxTime i32
}

type RowNdl struct {
  minfsw i32
  maxfsw i32
  unlimited boolean
  noStopLimit i32
  values []IGroup
}

type TableNdl struct {
  tableCode string
  tableName string
  tableData []IRowNdl
}

// table for surface interval time and repetitive letter
type RowRgl struct {
  groupLetter string
  minTime i32
  maxTime i32
  repetLetter string
}

type TableRgl struct {
  tableCode string
  tableName string
  tableData []IRowRgl
}

// table for residual nitrogen time
type Rnt struct {
  minDepth i32
  maxDepth i32
  rnt i32
}

type RowRnt struct {
  repetLetter string
  rnt []IRnt
}

type TableRnt struct {
  tableCode string
  tableName string
  tableNote9981 string
  tableData []IRowRnt
}

// table for air decompression
type DecoStops struct {
  depth i32
  time i32
}

type RowDeco struct {
  minTime i32
  maxTime i32
  airTAT string
  o2TAT string
  ttfs string
  o2cp f32
  repetLetter string
  surDo2Recom boolean
  exceptionalExposure boolean
  surDo2Req boolean
  strictlySurDo2 boolean
  airDecoStops []IDecoStops
  o2decoStops []IDecoStops
}

type DecoDepth struct {
  minfsw i32
  maxfsw i32
  rows []IRowDeco
}

type TableAirDeco struct {
  tableCode string
  tableName string
  tableData []IDecoDepth
}


// USER INPUT TYPES
type DivePlan struct {
  depth i32
  bottomTime i32
  sit i32
  nextDepth i32
}

type Dive struct {
  depth i32
  bottomTime i32
}