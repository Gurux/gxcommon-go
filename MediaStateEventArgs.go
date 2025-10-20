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

// MediaStateEventArgs is an argument class for media state changed event.
type MediaStateEventArgs struct {
	// Status information from media state.
	state MediaState

	// True is returned if media state change succeeded.
	accept bool
}

// NewMediaStateEventArgs creates a MediaStateEventArgs with the given state.
// The event is marked as accepted by default (accept = true).
func NewMediaStateEventArgs(state MediaState) *MediaStateEventArgs {
	return &MediaStateEventArgs{accept: true, state: state}
}

// State returns the media state associated with this event.
func (e *MediaStateEventArgs) State() MediaState {
	return e.state
}

// Accepted reports whether the event has been accepted/handled.
func (e *MediaStateEventArgs) Accepted() bool {
	return e.accept
}

// SetAccepted marks the event as accepted or not.
func (e *MediaStateEventArgs) SetAccepted(v bool) {
	e.accept = v
}
