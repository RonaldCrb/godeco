package deco

import (
	"fmt"

	"github.com/RonaldCrb/diving-decompression-go/airtables"
)

// DivePlan ...
/// single dive plan object
type DivePlan struct {
	/// the depth of the first planned dive expressed in feet of sea water
	Depth uint16 `json:"depth"`
	/// the bottom time of the first planned dive expressed in minutes
	BottomTime uint16 `json:"bottomTime"`
	/// the planned surface interval time expressed in minutes
	SIT uint16 `json:"sit"`
	/// the depth of the next planned dive expressed in feet of sea water
	NextDepth uint16 `json:"nextDepth"`
}

// PlanFromDive Instantiates a new Dive Plan object from an existing Dive Object
// the next_dive_depth is expressed in feet of sea water
// the surface_interval_time is expressed in minutes
// the returned object is of type DivePlan
func PlanFromDive(dive Dive, sit uint16, ndd uint16) DivePlan {
	return DivePlan{
		Depth:      dive.Depth,
		BottomTime: dive.BottomTime,
		SIT:        sit,
		NextDepth:  ndd,
	}
}

// NoDecompressionLimit returns the no decompression limit for the first dive
// of a DivePlan object No decompression limit is returned in minutes as u16 integer
func (d DivePlan) NoDecompressionLimit() (uint16, error) {
	var ndl uint16
	if d.Depth > 190 {
		return ndl, nil
	}
	NDtable, err := airtables.NoDecoTable()
	if err != nil {
		return 0, err
	}
	for _, row := range NDtable.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			ndl = row.NoStopLimit
		}
	}
	return ndl, nil
}

// GroupLetter returns the group letter of the first dive of a DivePlan object.
// the depth is expressed in feet of sea water
// the bottom_time is expressed in minutes
// the group letter is returned as a String
func (d DivePlan) GroupLetter() (string, error) {
	var gl string
	NDtable, err := airtables.NoDecoTable()
	if err != nil {
		return gl, err
	}
	for _, row := range NDtable.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			for _, value := range row.Values {
				if value.MinTime <= d.BottomTime && d.BottomTime <= value.MaxTime {
					gl = value.GroupLetter
				}
			}
		}
	}
	if gl == "" && d.Depth > 0 && d.Depth <= 10 && d.BottomTime > 462 {
		gl = "F"
	} else if gl == "" && d.Depth > 10 && d.Depth <= 15 && d.BottomTime > 449 {
		gl = "I"
	} else if gl == "" && d.Depth > 15 && d.Depth <= 20 && d.BottomTime > 461 {
		gl = "L"
	}
	if gl == "" && d.Depth <= 190 {
		gl = fmt.Sprint("this dive is out of the time range for no-decompression air dives")
	} else if gl == "" && d.Depth > 190 {
		gl = fmt.Sprint("this dive is out of the depth range for no-decompression air dives")
	}
	return gl, nil
}

// RepetLetter Returns the repetitive group letter of the DivePlan object.
// the depth and next_dive_depth are expressed in feet of sea water
// the bottom_time and surface_interval_time are expressed in minutes
// the repetitive group letter is returned as a String
func (d DivePlan) RepetLetter() (string, error) {
	NDtable, err := airtables.NoDecoTable()
	if err != nil {
		return "", err
	}

	RGLtable, err := airtables.RGLTable()
	if err != nil {
		return "", err
	}

	var rl string

	for _, row := range NDtable.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			for _, group := range row.Values {
				if group.MinTime <= d.BottomTime && d.BottomTime <= group.MaxTime {
					for _, rglRow := range RGLtable.TableData {
						if rglRow.GroupLetter == group.GroupLetter && rglRow.MinTime <= d.SIT && d.SIT <= rglRow.MaxTime {
							rl = rglRow.RepetLetter
						}
					}
				}
			}
		}
	}
	return rl, nil
}

// ResidualNitrogenTime Returns the residual nitrogen time of the DivePlan object.
// the depth and next_dive_depth are expressed in feet of sea water
// the bottom_time and surface_interval_time are expressed in minutes
// the residual nitrogen time is returned as a u16 integer
func (d DivePlan) ResidualNitrogenTime() (uint16, error) {
	NDtable, err := airtables.NoDecoTable()
	if err != nil {
		return 0, err
	}

	RGLtable, err := airtables.RGLTable()
	if err != nil {
		return 0, err
	}

	RNTtable, err := airtables.RNTTable()
	if err != nil {
		return 0, err
	}

	var rnt uint16

	for _, row := range NDtable.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			for _, group := range row.Values {
				if group.MinTime <= d.BottomTime && d.BottomTime <= group.MaxTime {
					for _, rglRow := range RGLtable.TableData {
						if rglRow.GroupLetter == group.GroupLetter && rglRow.MinTime <= d.SIT && d.SIT <= rglRow.MaxTime {
							for _, rntColumn := range RNTtable.TableData {
								if rntColumn.RepetLetter == rglRow.RepetLetter {
									for _, element := range rntColumn.RNT {
										if element.MinDepth <= d.NextDepth && d.NextDepth <= element.MaxDepth {
											rnt = element.RNT
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return rnt, nil
}
