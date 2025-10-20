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

// TraceLevel enumerates trace verbosity levels.
// Higher levels typically include all lower-severity messages.
type TraceLevel int

const (
	// TraceLevelOff disables tracing.
	TraceLevelOff TraceLevel = iota
	// TraceLevelError reports errors only.
	TraceLevelError
	// TraceLevelWarning reports warnings and errors.
	TraceLevelWarning
	// TraceLevelInfo reports informational messages, warnings, and errors.
	TraceLevelInfo
	// TraceLevelVerbose enables verbose/debug output (includes info, warnings, errors).
	TraceLevelVerbose
)

// TraceLevelParse converts the given string into a TraceLevel value.
//
// It returns the corresponding TraceLevel constant if the string matches
// a known level name, or an error if the input is invalid.
func TraceLevelParse(value string) (TraceLevel, error) {
	var ret TraceLevel
	var err error
	switch strings.ToUpper(value) {
	case "OFF":
		ret = TraceLevelOff
	case "ERROR":
		ret = TraceLevelError
	case "WARNING":
		ret = TraceLevelWarning
	case "INFO":
		ret = TraceLevelInfo
	case "VERBOSE":
		ret = TraceLevelVerbose
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the trace level.
// It satisfies fmt.Stringer.
func (value TraceLevel) String() string {
	var ret string
	switch value {
	case TraceLevelOff:
		ret = "OFF"
	case TraceLevelError:
		ret = "ERROR"
	case TraceLevelWarning:
		ret = "WARNING"
	case TraceLevelInfo:
		ret = "INFO"
	case TraceLevelVerbose:
		ret = "VERBOSE"
	default:
	}
	return ret
}
