package gxcommon

import "golang.org/x/text/language"

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

// ReceivedEventHandler is a callback invoked when data is received from a media.
// The first argument is the source media and the second carries the payload.
type ReceivedEventHandler func(IGXMedia, ReceiveEventArgs)

// MediaStateHandler is a callback invoked when media state changes.
// The first argument is the source media and the second contains state details.
type MediaStateHandler func(IGXMedia, MediaStateEventArgs)

// TraceEventHandler is a callback invoked when a trace event occurs.
// The first argument is the source media and the second contains trace details.
type TraceEventHandler func(IGXMedia, TraceEventArgs)

// ErrorEventHandler is a callback invoked when an error occurs in the media.
// The first argument is the source media and the second is the error.
type ErrorEventHandler func(IGXMedia, error)

// IGXMedia defines a common interface for media components.
type IGXMedia interface {

	// Send transmits data asynchronously.
	// receiver contains media-specific receiver information and may be empty.
	Send(data any, receiver string) error

	// Receive waits for reply data according to args.
	// The returned bool reports whether data was received.
	Receive(args *ReceiveParameters) (bool, error)

	// SetOnReceived sets a callback for asynchronously received data.
	SetOnReceived(ReceivedEventHandler)

	// SetOnError sets a callback for media errors.
	SetOnError(ErrorEventHandler)

	// SetOnMediaStateChange sets a callback for media state changes.
	SetOnMediaStateChange(MediaStateHandler)

	// SetOnTrace sets a callback for trace events.
	SetOnTrace(TraceEventHandler)

	// Copy copies the media configuration to target.
	Copy(target IGXMedia) error

	// GetName returns a unique media connection name.
	GetName() string

	// GetTrace returns the current trace level.
	GetTrace() TraceLevel

	// SetTrace sets the trace level.
	SetTrace(TraceLevel) error

	// Open opens the media connection.
	Open() error

	// IsOpen reports whether the media connection is open.
	IsOpen() bool

	// Close closes the media connection.
	Close() error

	// GetMediaType returns the media type.
	GetMediaType() string

	// GetSettings returns media settings in serialized form.
	GetSettings() string

	// SetSettings applies serialized media settings.
	SetSettings(value string) error

	// GetSynchronous enters synchronous mode and returns a restore function.
	GetSynchronous() func()

	// IsSynchronous reports whether synchronous mode is enabled.
	IsSynchronous() bool

	// ResetSynchronousBuffer clears the synchronous receive buffer.
	ResetSynchronousBuffer()

	// GetBytesSent returns the sent byte count.
	GetBytesSent() uint64

	// GetBytesReceived returns the received byte count.
	GetBytesReceived() uint64

	// ResetByteCounters resets sent and received byte counters.
	ResetByteCounters()

	// Validate validates media settings before opening a connection.
	Validate() error

	// SetEop sets the end-of-packet marker for receive buffering.
	SetEop(any)

	// GetEop returns the configured end-of-packet marker.
	GetEop() any

	// Localize localizes messages for the specified language.
	Localize(language language.Tag)
}
