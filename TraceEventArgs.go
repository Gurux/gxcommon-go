package gxcommon

import (
	"fmt"
	"time"
)

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

// TraceEventArgs carries information about a trace event.
// It includes the event timestamp, trace type (direction/category),
// optional payload, and receiver metadata.
type TraceEventArgs struct {
	// timestamp is the time when the event occurred.
	timestamp time.Time

	// traceType tells whether data was sent or received and the trace category.
	traceType TraceTypes

	// data is the received/sent payload (optional).
	data any

	// Receiver is media-dependent metadata about the receiver (optional),
	// e.g., a remote address for TCP.
	Receiver string
}

// Timestamp returns the time when the event occurred.
func (e *TraceEventArgs) Timestamp() time.Time { return e.timestamp }

// Type returns the trace category/direction.
func (e *TraceEventArgs) Type() TraceTypes { return e.traceType }

// Data returns the optional payload associated with the event.
func (e *TraceEventArgs) Data() any { return e.data }

// NewTraceEventArgs creates a TraceEventArgs with the given trace type,
// optional payload, and receiver metadata. The timestamp is set to time.Now()
// at the moment of construction.
func NewTraceEventArgs(traceType TraceTypes, data any, receiver string) *TraceEventArgs {
	return &TraceEventArgs{
		timestamp: time.Now(),
		traceType: traceType,
		data:      data,
		Receiver:  receiver,
	}
}

// String returns the content of the trace event as a string.
func (e *TraceEventArgs) String() string {
	str, _ := ToString(e.data)
	return fmt.Sprintf("%s\t%s\t%s", e.timestamp.Format("15:04:05.000"), e.traceType.String(), str)
}
