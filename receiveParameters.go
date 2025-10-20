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

// ReceiveParameters class is used when data is read synchronously.
type ReceiveParameters struct {
	// If true, returns the bytes from the buffer without removing.
	Peek bool

	// EOP: The end of packet (EOP) waited for.
	// The EOP can, for example be a single byte ('0xA1'),
	// a string ("OK") or an array of bytes.
	EOP any

	// Count: The number of reply data bytes to be read.
	// Count can be between 0 and n bytes.
	Count int

	// WaitTime: Maximum time, in milliseconds, to wait for reply data.
	// WaitTime -1 (Default value) indicates infinite wait time.
	WaitTime int

	// AllData: If True, all the reply data is moved to Reply.
	AllData bool

	// Reply: Received reply data.
	Reply any

	// ReplyType: The type of the reply data.
	// Supported types are: string, []byte, byte (uint8), rune (char),
	// int16, int32, int64, uint16, uint32, uint64.
	// Default value is TypeUnknown. If TypeUnknown, the type is determined
	// from the Reply field.
	ReplyType DataType
}

// NewReceiveParameters returns a new ReceiveParameters[T] initialized with
// sensible defaults for synchronous reads. By default Peek is false (data is
// consumed when read) and WaitTime is -1, which indicates an infinite wait.
//
// The caller can override fields after construction as needed.
//
// Example:
//
//	rp := NewReceiveParameters[[]byte]()
//	rp.Peek = true        // inspect without consuming
//	rp.WaitTime = 5000    // 5s timeout
func NewReceiveParameters[T any]() *ReceiveParameters {
	ret := &ReceiveParameters{
		Peek:     false,
		WaitTime: -1,
	}
	ret.ReplyType = GetType[T]()
	return ret
}
