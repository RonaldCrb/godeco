package airtables

import (
	"embed"
	"encoding/json"
	"errors"
)

// table for no decompression limit and group letter

// Group ...
type Group struct {
	GroupLetter string `json:"group_letter"`
	MinTime     uint16 `json:"min_time"`
	MaxTime     uint16 `json:"max_time"`
}

// RowNdl ...
type RowNdl struct {
	MinFSW      uint16  `json:"min_fsw"`
	MaxFSW      uint16  `json:"max_fsw"`
	Unlimited   bool    `json:"unlimited"`
	NoStopLimit uint16  `json:"no_stop_limit"`
	Values      []Group `json:"values"`
}

// TableNdl ...
type TableNdl struct {
	TableCode string   `json:"table_code"`
	TableName string   `json:"table_name"`
	TableData []RowNdl `json:"table_data"`
}

// table for surface interval time and repetitive letter

// RowRgl ...
type RowRgl struct {
	GroupLetter string `json:"group_letter"`
	MinTime     uint16 `json:"min_time"`
	MaxTime     uint16 `json:"max_time"`
	RepetLetter string `json:"repet_letter"`
}

// TableRgl ...
type TableRgl struct {
	TableCode string   `json:"table_code"`
	TableName string   `json:"table_name"`
	TableData []RowRgl `json:"table_data"`
}

// table for residual nitrogen time

// Rnt ...
type Rnt struct {
	MinDepth uint16 `json:"min_depth"`
	MaxDepth uint16 `json:"max_depth"`
	RNT      uint16 `json:"rnt"`
}

// RowRnt ...
type RowRnt struct {
	RepetLetter string `json:"repet_letter"`
	RNT         []Rnt  `json:"rnt"`
}

// TableRnt ...
type TableRnt struct {
	TableCode     string   `json:"table_code"`
	TableName     string   `json:"table_name"`
	TableNote9981 string   `json:"table_note_9981"`
	TableData     []RowRnt `json:"table_data"`
}

// table for air decompression

// DecoStops ...
type DecoStops struct {
	Depth uint16 `json:"depth"`
	Time  uint16 `json:"time"`
}

// RowDeco ...
type RowDeco struct {
	MinTime             uint16      `json:"min_time"`
	MaxTime             uint16      `json:"max_time"`
	AIRTAT              string      `json:"air_tat"`
	O2TAT               string      `json:"o2_tat"`
	TTFS                string      `json:"ttfs"`
	O2CP                float32     `json:"o2cp"`
	RepetLetter         string      `json:"repetgroup_letter"`
	SurDo2Recom         bool        `json:"surdo2_recommended"`
	ExceptionalExposure bool        `json:"exceptional_exposure"`
	SurDo2Req           bool        `json:"surdo2_required"`
	StrictlySurDo2      bool        `json:"strict_surdo2"`
	AirDecoStops        []DecoStops `json:"air_deco_stops"`
	O2decoStops         []DecoStops `json:"o2_deco_stops"`
}

// DecoDepth ...
type DecoDepth struct {
	MinFSW uint16    `json:"min_fsw"`
	MaxFSW uint16    `json:"max_fsw"`
	Rows   []RowDeco `json:"rows"`
}

// TableAirDeco ...
type TableAirDeco struct {
	TableCode string      `json:"table_code"`
	TableName string      `json:"table_name"`
	TableData []DecoDepth `json:"table_data"`
}

//go:embed JSON/usnavy-air-nodeco-rev7.json
var nodecoTable embed.FS

//go:embed JSON/usnavy-air-deco-rev7.json
var decoTable embed.FS

//go:embed JSON/usnavy-air-repetgroup-rev7.json
var rglTable embed.FS

//go:embed JSON/usnavy-air-rnt-rev7.json
var rntTable embed.FS

// NoDecoTable returns a typed and serialized US Navy air
// no-decompression table from rev7 of the US Navy dive manual.
func NoDecoTable() (TableNdl, error) {
	jsonFile, err := nodecoTable.ReadFile("JSON/usnavy-air-nodeco-rev7.json")
	if err != nil {
		return TableNdl{}, errors.New(err.Error())
	}
	var sertab TableNdl

	json.Unmarshal(jsonFile, &sertab)
	return sertab, nil
}

// DecoTable returns a typed and serialized US Navy air
// decompression table from rev7 of the US Navy dive manual
func DecoTable() (TableAirDeco, error) {
	jsonFile, err := decoTable.ReadFile("JSON/usnavy-air-deco-rev7.json")
	if err != nil {
		return TableAirDeco{}, errors.New(err.Error())
	}
	var sertab TableAirDeco

	json.Unmarshal(jsonFile, &sertab)
	return sertab, nil
}

// RGLTable returns a typed and serialized US Navy repetitive group letter
// table from rev7 of the US Navy dive manual
func RGLTable() (TableRgl, error) {
	jsonFile, err := rglTable.ReadFile("JSON/usnavy-air-repetgroup-rev7.json")
	if err != nil {
		return TableRgl{}, errors.New(err.Error())
	}
	var sertab TableRgl

	json.Unmarshal(jsonFile, &sertab)
	return sertab, nil
}

// RNTTable returns a typed and serialized US Navy residual nitrogen time
// table from rev7 of the US Navy dive manual
func RNTTable() (TableRnt, error) {
	jsonFile, err := rntTable.ReadFile("JSON/usnavy-air-rnt-rev7.json")
	if err != nil {
		return TableRnt{}, errors.New(err.Error())
	}
	var sertab TableRnt

	json.Unmarshal(jsonFile, &sertab)
	return sertab, nil
}
