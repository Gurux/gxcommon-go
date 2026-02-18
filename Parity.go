package gxcommon

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software; you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
// ---------------------------------------------------------------------------

import (
	"fmt"
	"strings"
)

// Parity represents the parity bit configuration for serial communication.
type Parity int

const (
	// ParityNone indicates that no parity bit is used.
	ParityNone Parity = iota
	// ParityOdd indicates that the parity bit is set to 1 if the number of 1 bits in the data is odd, and 0 otherwise.
	ParityOdd
	// ParityEven indicates that the parity bit is set to 1 if the number of 1 bits in the data is even, and 0 otherwise.
	ParityEven
	// ParityMark indicates that the parity bit is always set to 1.
	ParityMark
	// ParitySpace indicates that the parity bit is always set to 0.
	ParitySpace
)

// ParityParse converts the given string into a Parity value.
//
// It returns the corresponding Parity constant if the string matches
// a known level name, or an error if the input is invalid.
func ParityParse(value string) (Parity, error) {
	var ret Parity
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = ParityNone
	case strings.EqualFold(value, "Odd"):
		ret = ParityOdd
	case strings.EqualFold(value, "Even"):
		ret = ParityEven
	case strings.EqualFold(value, "Mark"):
		ret = ParityMark
	case strings.EqualFold(value, "Space"):
		ret = ParitySpace
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Parity.
// It satisfies fmt.Stringer.
func (g Parity) String() string {
	var ret string
	switch g {
	case ParityNone:
		ret = "None"
	case ParityOdd:
		ret = "Odd"
	case ParityEven:
		ret = "Even"
	case ParityMark:
		ret = "Mark"
	case ParitySpace:
		ret = "Space"
	}
	return ret
}

// AllParity returns a slice containing all defined Parity values.
func AllParity() []Parity {
	return []Parity{
		ParityNone,
		ParityOdd,
		ParityEven,
		ParityMark,
		ParitySpace,
	}
}
