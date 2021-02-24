package airtables

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// table for no decompression limit and group letter

// Group ...
type Group struct {
	GroupLetter string `json:"groupLetter"`
	MinTime     uint16 `json:"minTime"`
	MaxTime     uint16 `json:"maxTime"`
}

// RowNdl ...
type RowNdl struct {
	MinFSW      uint16  `json:"minfsw"`
	MaxFSW      uint16  `json:"maxfsw"`
	Unlimited   bool    `json:"unlimited"`
	NoStopLimit uint16  `json:"noStopLimit"`
	Values      []Group `json:"values"`
}

// TableNdl ...
type TableNdl struct {
	TableCode string   `json:"tableCode"`
	TableName string   `json:"tableName"`
	TableData []RowNdl `json:"tableData"`
}

// table for surface interval time and repetitive letter

// RowRgl ...
type RowRgl struct {
	GroupLetter string `json:"groupLetter"`
	MinTime     uint16 `json:"minTime"`
	MaxTime     uint16 `json:"maxTime"`
	RepetLetter string `json:"repetLetter"`
}

// TableRgl ...
type TableRgl struct {
	TableCode string   `json:"tableCode"`
	TableName string   `json:"tableName"`
	TableData []RowRgl `json:"tableData"`
}

// table for residual nitrogen time

// Rnt ...
type Rnt struct {
	MinDepth uint16 `json:"minDepth"`
	MaxDepth uint16 `json:"maxDepth"`
	RNT      uint16 `json:"rnt"`
}

// RowRnt ...
type RowRnt struct {
	RepetLetter string `json:"repetLetter"`
	RNT         []Rnt  `json:"rnt"`
}

// TableRnt ...
type TableRnt struct {
	TableCode     string   `json:"tableCode"`
	TableName     string   `json:"tableName"`
	TableNote9981 string   `json:"tableNote9981"`
	TableData     []RowRnt `json:"tableData"`
}

// table for air decompression

// DecoStops ...
type DecoStops struct {
	Depth uint16 `json:"depth"`
	Time  uint16 `json:"time"`
}

// RowDeco ...
type RowDeco struct {
	MinTime             uint16      `json:"minTime"`
	MaxTime             uint16      `json:"maxTime"`
	AIRTAT              string      `json:"airTAT"`
	O2TAT               string      `json:"o2TAT"`
	TTFS                string      `json:"ttfs"`
	O2CP                float32     `json:"o2cp"`
	RepetLetter         string      `json:"repetLetter"`
	SurDo2Recom         bool        `json:"surDo2Recom"`
	ExceptionalExposure bool        `json:"exceptionalExposure"`
	SurDo2Req           bool        `json:"surDo2Req"`
	StrictlySurDo2      bool        `json:"strictlySurDo2"`
	AirDecoStops        []DecoStops `json:"airDecoStops"`
	O2decoStops         []DecoStops `json:"o2decoStops"`
}

// DecoDepth ...
type DecoDepth struct {
	MinFSW uint16    `json:"minfsw"`
	MaxFSW uint16    `json:"maxfsw"`
	Rows   []RowDeco `json:"rows"`
}

// TableAirDeco ...
type TableAirDeco struct {
	TableCode string      `json:"tableCode"`
	TableName string      `json:"tableName"`
	TableData []DecoDepth `json:"tableData"`
}

// NoDecoTable returns a typed and serialized US Navy air
// no-decompression table from rev7 of the US Navy dive manual.
func NoDecoTable() (TableNdl, error) {
	jsonFile, err := os.Open("JSON/usnavy-air-nodeco-rev7.json")
	if err != nil {
		return TableNdl{}, errors.New(err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return TableNdl{}, errors.New(err.Error())
	}
	var sertab TableNdl

	json.Unmarshal(byteValue, &sertab)
	return sertab, nil
}

// DecoTable returns a typed and serialized US Navy air
// decompression table from rev7 of the US Navy dive manual
func DecoTable() (TableAirDeco, error) {
	jsonFile, err := os.Open("JSON/usnavy-air-deco-rev7.json")
	if err != nil {
		return TableAirDeco{}, errors.New(err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var sertab TableAirDeco

	json.Unmarshal(byteValue, &sertab)
	return sertab, nil
}

// RGLTable returns a typed and serialized US Navy repetitive group letter
// table from rev7 of the US Navy dive manual
func RGLTable() (TableRgl, error) {
	jsonFile, err := os.Open("JSON/usnavy-air-repetgroup-rev7.json")
	if err != nil {
		return TableRgl{}, errors.New(err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return TableRgl{}, errors.New(err.Error())
	}
	var sertab TableRgl

	json.Unmarshal(byteValue, &sertab)
	return sertab, nil
}

// RNTTable returns a typed and serialized US Navy residual nitrogen time
// table from rev7 of the US Navy dive manual
func RNTTable() (TableRnt, error) {
	jsonFile, err := os.Open("JSON/usnavy-air-rnt-rev7.json")
	if err != nil {
		return TableRnt{}, errors.New(err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return TableRnt{}, errors.New(err.Error())
	}
	var sertab TableRnt

	json.Unmarshal(byteValue, &sertab)
	return sertab, nil
}
