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

// TraceTypes is a enumerated value that selects which traces are emitted.
type TraceTypes int

const (
	// TraceTypesSent indicates data was sent.
	TraceTypesSent TraceTypes = 0x1

	// TraceTypesReceived indicates data was received.
	TraceTypesReceived TraceTypes = 0x2

	// TraceTypesError indicates an error occurred.
	TraceTypesError TraceTypes = 0x4

	// TraceTypesWarning indicates a warning was issued.
	TraceTypesWarning TraceTypes = 0x8

	// TraceTypesInfo indicates an informational message (e.g., media state notifications).
	TraceTypesInfo TraceTypes = 0x10
)

// TraceTypesParse converts the given string into a TraceTypes value.
//
// It returns the corresponding TraceTypes constant if the string matches
// a known level name, or an error if the input is invalid.
func (TraceTypes*) TraceTypesParse(value string) (TraceTypes, error) {
	var ret TraceTypes
	var err error
	switch strings.ToUpper(value) {
	case "SENT":
		ret = TraceTypesSent
	case "RECEIVED":
		ret = TraceTypesReceived
	case "ERROR":
		ret = TraceTypesError
	case "WARNING":
		ret = TraceTypesWarning
	case "INFO":
		ret = TraceTypesInfo
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the trace type.
// It satisfies fmt.Stringer.
func (g TraceTypes) String() string {
	var ret string
	switch g {
	case TraceTypesSent:
		ret = "SENT"
	case TraceTypesReceived:
		ret = "RECEIVED"
	case TraceTypesError:
		ret = "ERROR"
	case TraceTypesWarning:
		ret = "WARNING"
	case TraceTypesInfo:
		ret = "INFO"
	}
	return ret
}
