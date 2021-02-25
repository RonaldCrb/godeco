// diving-decompression crate
// Written in 2021 by
// Ronald Alonzo <alonzo.ronald@gmail.com>

// To the extent possible under law, the author(s) have dedicated all
// copyright and related and neighboring rights to this software to
// the public domain worldwide. This software is distributed without
// any warranty.

//! # Go diving-decompression library
//!
//! # UNDER CONSTRUCTION

//! ## DO NOT USE THIS PACKAGE UNTIL STABLE VERSION HAS BEEN RELEASED!!

//! ## IMPORTANT NOTE FROM THE AUTHOR
//! this package is under construction, it is **__NOT__** suitable for
//! usage in real dive operations neither commercial nor recreational,
//! as we need to make extensive test and audit the package reliability.
//! it is not only a matter of applying unit testing as this calculations
//! are crucial for divers safety; also regardless of the extensive tests
//! and trials in humans performed by the US Navy along the years with
//! regards of decompression sickness, it has been stated many times by
//! relevant stakeholders that these trials do not necessarily entail 100%
//! accuracy on the results of undertaking dive operations within the
//! constraints of these dive tables. there are many factors that are not
//! taken into consideration (e.g: water temperature, diver physiological
//! fitness, unadverted PFOs... to name a few).

//! This is a library created with the purpose of assisting diving
//! professionals in planning decompression procedures for air diving
//! operations as per the US Navy diving manual rev7.
//!
//! It was initially written in TypeScript and then ported to Rust to
//! harness the benefits of a much stronger type system. These safety
//! guarantees are of crucial importance when dealing with operational
//! and procedural safety in the commercial diving industry.

//! In 2021 this library was ported to Golang with the interest of creating
//! a series of bots that can distribute the access to these calculations
//! and hopefully educate the end users on the dangers of ther underwater
//! activities. it will also ease the economic calculation overhead by
//! allowing a digital actor that calculates decompression protocols with
//! no error.
//!
//! This project is and will always be 100% free and open source.
//! it is open for public review and we welcome PRs as long as they
//! adhere to international guidelines and acknowledged best practices
//! in the industry, specially those contained within the US Navy dive
//! manual which is __THE ONLY__ scientifically derived set of guidelines.
//!
//! Pull Requests based on anecdotical or empirical evidence or those
//! that could contain private parties agendas will always be dismissed
//! by the authors of this project. we do not tolerate private tables
//! and protocols that aim to distort the good practices in order to
//! increase allowed diving depth and time limits and shortened decompression
//! procedures with economical purposes.

package deco

import (
	"errors"
	"fmt"

	"github.com/RonaldCrb/diving-decompression-go/airtables"
)

// Dive object
type Dive struct {
	/// depth of the dive expressed in feet of sea water
	Depth uint16 `json:"depth"`
	/// bottom time of the dive expressed in minutes
	BottomTime uint16 `json:"bottomTime"`
}

// NoDecompressionLimit returns the no decompression limit of the Dive
// Object up to a depth of 190 feet of sea water no decompression limit
// is returned in minutes as u16 integer
func (d Dive) NoDecompressionLimit() (uint16, error) {
	var ndl uint16 = 0

	if d.Depth > 190 {
		return ndl, nil
	}

	table, err := airtables.NoDecoTable()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	for _, row := range table.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			ndl = row.NoStopLimit
		}
	}

	return ndl, nil
}

// GroupLetter returns the group letter of the Dive object.
// the depth is expressed in feet of sea water
// the bottom_time is expressed in minutes
// the group letter is returned as a String
func (d Dive) GroupLetter() (string, error) {
	var gl string

	nodecoTable, err := airtables.NoDecoTable()
	if err != nil {
		return "", errors.New(err.Error())
	}

	for _, row := range nodecoTable.TableData {
		if row.MinFSW <= d.Depth && d.Depth <= row.MaxFSW {
			for _, value := range row.Values {
				if value.MinTime <= d.BottomTime && d.BottomTime <= value.MaxTime {
					gl = fmt.Sprint(value.GroupLetter)
				}
			}
		}
	}

	if gl == "" && d.Depth > 0 && d.Depth <= 10 && d.BottomTime > 462 {
		gl = fmt.Sprint("F")
	} else if gl == "" && d.Depth > 10 && d.Depth <= 15 && d.BottomTime > 449 {
		gl = fmt.Sprint("I")
	} else if gl == "" && d.Depth > 15 && d.Depth <= 20 && d.BottomTime > 461 {
		gl = fmt.Sprint("L")
	}

	if gl == "" && d.Depth <= 190 {
		gl = fmt.Sprint("this dive is out of the time range for no-decompression air dives")
	} else if gl == "" && d.Depth > 190 {
		gl = fmt.Sprint("this dive is out of the depth range for no-decompression air dives")
	}

	return gl, nil
}

// DecoDive returns the decompression profile of the Dive object.
// the depth is expressed in feet of sea water
// the bottom_time is expressed in minutes
// the decompression profile is returned as a RowDeco struct
func (d Dive) DecoDive() (airtables.RowDeco, error) {
	table, err := airtables.DecoTable()
	if err != nil {
		return airtables.RowDeco{}, errors.New("Error deserializing no decompression table")
	}

	var decoProfile airtables.RowDeco = airtables.RowDeco{}

	for _, rd := range table.TableData {
		if rd.MinFSW <= d.Depth && d.Depth <= rd.MaxFSW {
			for _, profile := range rd.Rows {
				if profile.MinTime <= d.BottomTime && profile.MaxTime <= d.BottomTime {
					decoProfile = profile
				}
			}
		}
	}
	return decoProfile, nil
}
