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
// The first argument is the source media (IGXMedia). The second argument carries
// the received payload and related metadata (ReceiveEventArgs). Implementations
// should return quickly; long-running work should be offloaded to another goroutine.
type ReceivedEventHandler func(IGXMedia, ReceiveEventArgs)

// MediaStateHandler is a callback invoked when the media state changes
// (for example, connecting, connected, disconnected, error). The first
// argument is the source media (IGXMedia); the second provides details
// about the state transition (MediaStateEventArgs). Handlers should
// return quickly—offload long work to another goroutine.
type MediaStateHandler func(IGXMedia, MediaStateEventArgs)

// TraceEventHandler is a callback invoked when a trace event occurs.
// The first argument is the source media (IGXMedia); the second contains
// the trace details (TraceEventArgs). Handlers should return quickly—
// offload any heavy work to a separate goroutine.
type TraceEventHandler func(IGXMedia, TraceEventArgs)

// ErrorEventHandler is a callback invoked when an error occurs in the media.
// The first argument is the source media (IGXMedia); the second is the error
// (which may be wrapped). Handlers should return quickly—offload any heavy
// work to another goroutine.
type ErrorEventHandler func(IGXMedia, error)

// IGXMedia is a common interface for all Media components.
// Using this interface GXCommunication library enables communication with
// different medias.
type IGXMedia interface {

	// Sends data asynchronously.
	// No reply from the receiver, whether or not the operation was successful, is expected.
	// data: Data to send to the device.
	// receiver : Media depend information of the receiver (optional).
	Send(data any, receiver string) error

	// Receive Waits for more reply data After SendSync if whole packet is not received yet.
	// args: Receive data arguments.</param>
	// <returns>True, if the send operation was successful.</returns>
	Receive(args *ReceiveParameters) (bool, error)

	//SetOnReceived: Media component notifies asynchronously received data through this method.
	SetOnReceived(ReceivedEventHandler)

	//SetOnError: Media component notifies the error event of a Gurux component.)
	SetOnError(ErrorEventHandler)

	//SetOnMediaStateChange: Media component notifies the state change event of a Gurux media component.
	SetOnMediaStateChange(MediaStateHandler)

	//SetOnTrace: Media component notifies the trace event of a Gurux media component.
	SetOnTrace(TraceEventHandler)

	// Copy copies the content of the media to target media.
	Copy(target IGXMedia) error

	// Media name is used to identify media connection,
	// so two different media connection can not return same media name.
	GetName() string

	// The trace level specifies which types of trace messages are emitted.
	GetTrace() TraceLevel

	// The trace level specifies which types of trace messages are emitted.
	SetTrace(TraceLevel) error

	//Open ppens the media.
	Open() error

	// IsOpen returns true, if the connection is established.
	IsOpen() bool

	//Close terminates the media connection.
	Close() error

	// MediaType returns the type of the media.
	GetMediaType() string

	// GetSettings gets the media settings.
	GetSettings() string

	// SetSettings sets the media settings.
	SetSettings(value string) error

	// GetSynchronous makes the connection synchronized and stops sending OnReceived events.
	GetSynchronous() func()

	// IsSynchronous returns true if the media is synchronous mode.
	IsSynchronous() bool

	// ResetSynchronousBuffer resets synchronous buffer.
	ResetSynchronousBuffer()

	// GetBytesSent return sent byte count.
	GetBytesSent() uint64

	// BytesReceived returns received byte count.
	GetBytesReceived() uint64

	// Resets BytesReceived and BytesSent counters.
	ResetByteCounters()

	// Validate Media settings for connection open.
	Validate() error

	// Eop is used to buffer the data is buffered until EOP is received.
	SetEop(any)

	// Eop is used to buffer the data is buffered until EOP is received.
	GetEop() any

	// localize messages for the specified language.
	// No errors is returned if language is not supported.
	localize(language language.Tag)
}
