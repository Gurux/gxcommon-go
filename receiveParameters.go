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

// ReceiveParameters defines options for synchronous receive operations.
type ReceiveParameters struct {
	// Peek returns bytes from the buffer without consuming them.
	Peek bool

	// EOP is the end-of-packet marker to wait for.
	// It can be, for example, a single byte, a string, or a byte slice.
	EOP any

	// Count is the number of bytes to read.
	Count int

	// WaitTime is the maximum wait time in milliseconds.
	// A value of -1 means infinite wait.
	WaitTime int

	// AllData moves all available reply data to Reply when true.
	AllData bool

	// Reply contains received reply data.
	Reply any

	// ReplyType is the expected reply data type.
	// Supported types are string, []byte, uint8, rune, int16, int32, int64,
	// uint16, uint32, and uint64. If ReplyType is DataTypeUnknown, the type
	// is inferred from Reply.
	ReplyType DataType
}

// NewReceiveParameters returns a new ReceiveParameters initialized with
// defaults for synchronous reads.
//
// Default values are Peek=false and WaitTime=-1 (infinite wait).
func NewReceiveParameters[T any]() *ReceiveParameters {
	ret := &ReceiveParameters{
		Peek:     false,
		WaitTime: -1,
	}
	ret.ReplyType = GetType[T]()
	return ret
}
