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

// MediaState enumerates the lifecycle states of a media/connection.
type MediaState int

const (
	// MediaStateClosed indicates the media is closed.
	MediaStateClosed MediaState = 1
	// MediaStateOpen indicates the media is open.
	MediaStateOpen MediaState = 2
	// MediaStateOpening indicates the media is in the process of opening.
	MediaStateOpening MediaState = 3
	// MediaStateClosing indicates the media is in the process of closing.
	MediaStateClosing MediaState = 4
	// MediaStateChanged indicates the media type or configuration has changed.
	MediaStateChanged MediaState = 5
)

// MediaStateParse converts the given string into a MediaState value.
//
// It returns the corresponding MediaState constant if the string matches
// a known level name, or an error if the input is invalid.
func MediaStateParse(value string) (MediaState, error) {
	var ret MediaState
	var err error
	switch strings.ToUpper(value) {
	case "CLOSED":
		ret = MediaStateClosed
	case "OPEN":
		ret = MediaStateOpen
	case "OPENING":
		ret = MediaStateOpening
	case "CLOSING":
		ret = MediaStateClosing
	case "CHANGED":
		ret = MediaStateChanged
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Media state.
// It satisfies fmt.Stringer.
func (g MediaState) String() string {
	var ret string
	switch g {
	case MediaStateClosed:
		ret = "CLOSED"
	case MediaStateOpen:
		ret = "OPEN"
	case MediaStateOpening:
		ret = "OPENING"
	case MediaStateClosing:
		ret = "CLOSING"
	case MediaStateChanged:
		ret = "CHANGED"
	}
	return ret
}
