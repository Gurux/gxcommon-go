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

// StopBits defines the stop-bit configuration used in serial communication.
type StopBits int

const (
	// StopBitsNone indicates that no stop bits are used.
	StopBitsNone StopBits = iota
	// StopBitsOne indicates that one stop bit is used.
	StopBitsOne
	// StopBitsTwo indicates that two stop bits are used.
	StopBitsTwo
	// StopBitsOnePointFive indicates that one and a half stop bits are used.
	StopBitsOnePointFive
)

// StopBitsParse parses a string value into a StopBits.
//
// It returns ErrUnknownEnum if value does not match a supported stop-bit mode.
func StopBitsParse(value string) (StopBits, error) {
	var ret StopBits
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = StopBitsNone
	case strings.EqualFold(value, "One"):
		ret = StopBitsOne
	case strings.EqualFold(value, "Two"):
		ret = StopBitsTwo
	case strings.EqualFold(value, "OnePointFive"):
		ret = StopBitsOnePointFive
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical stop-bit name.
// It satisfies fmt.Stringer.
func (g StopBits) String() string {
	var ret string
	switch g {
	case StopBitsNone:
		ret = "None"
	case StopBitsOne:
		ret = "One"
	case StopBitsTwo:
		ret = "Two"
	case StopBitsOnePointFive:
		ret = "OnePointFive"
	}
	return ret
}

// AllStopBits returns all defined StopBits values.
func AllStopBits() []StopBits {
	return []StopBits{
		StopBitsNone,
		StopBitsOne,
		StopBitsTwo,
		StopBitsOnePointFive,
	}
}
