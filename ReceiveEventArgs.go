package gxcommon

import (
	"encoding/hex"
	"fmt"
	"strings"
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

// ReceiveEventArgs encapsulates data received from a media source.
// It carries the raw payload and optional sender-specific metadata.
type ReceiveEventArgs struct {
	// data is the raw payload received from the device.
	data []byte

	// senderInfo holds media-dependent metadata about the sender
	// (e.g., the remote address for TCP).
	senderInfo string
}

// NewReceiveEventArgs creates a ReceiveEventArgs with the given payload and
// sender information. The fields are stored in unexported form; access them
// via the corresponding getters (e.g., Data(), SenderInfo()) if exposed.
func NewReceiveEventArgs(data []byte, senderInfo string) *ReceiveEventArgs {
	return &ReceiveEventArgs{data: data, senderInfo: senderInfo}
}

// Data returns the received payload.
func (e *ReceiveEventArgs) Data() []byte {
	return e.data
}

// SenderInfo returns metadata about the sender/receiver endpoint.
func (e *ReceiveEventArgs) SenderInfo() string {
	return e.senderInfo
}

// String returns the content of the trace event as a string.
func (e *ReceiveEventArgs) String() string {
	str, _ := ToString(e.data)
	return fmt.Sprintf("%s\t%s", e.senderInfo, str)
}

// ToHex method is used to convert byte slice to hex string.
func ToHex(value []byte) string {
	if len(value) == 0 {
		return ""
	}
	out := make([]byte, len(value)*3-1)
	hex.Encode(out[0:2], value[0:1])
	for i := 1; i < len(value); i++ {
		out[i*3-1] = ' '
		hex.Encode(out[i*3:i*3+2], value[i:i+1])
	}
	return strings.ToUpper(string(out))
}
