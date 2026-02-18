package gxcommon

import (
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

// BaudRate represents the communication speed of a serial port
// in bits per second (bps).
//
// It is used to configure the transmission rate between two
// serial devices. Both ends of the connection must use the same
// baud rate to communicate correctly.
//
// Common baud rates include:
//
//	9600    – Legacy devices, modems
//	19200   – Industrial devices
//	38400   – Embedded systems
//	115200  – Most common modern default
//	230400+ – High-speed serial communication
//
// Example:
//
//	cfg := GXSerial{
//	    BaudRate: Baud115200,
//	}
//
// The underlying value of BaudRate is the numeric bits-per-second value.
type BaudRate int

const (
	// BaudRate50 defines that the baudrate is 50.
	BaudRate50 BaudRate = 50
	// BaudRate75 defines that the baudrate is 75.
	BaudRate75 BaudRate = 75
	// BaudRate110 defines that the baudrate is 110.
	BaudRate110 BaudRate = 110
	// BaudRate134 defines that the baudrate is 134.
	BaudRate134 BaudRate = 134
	// BaudRate150 defines that the baudrate is 150.
	BaudRate150 BaudRate = 150
	// BaudRate200 defines that the baudrate is 200.
	BaudRate200 BaudRate = 200
	// BaudRate300 defines that the baudrate is 300.
	BaudRate300 BaudRate = 300
	// BaudRate600 defines that the baudrate is 600.
	BaudRate600 BaudRate = 600
	// BaudRate1200 defines that the baudrate is 1200.
	BaudRate1200 BaudRate = 1200
	// BaudRate1800 defines that the baudrate is 1800.
	BaudRate1800 BaudRate = 1800
	// BaudRate2400 defines that the baudrate is 2400.
	BaudRate2400 BaudRate = 2400
	// BaudRate4800 defines that the baudrate is 4800.
	BaudRate4800 BaudRate = 4800
	// BaudRate9600 defines that the baudrate is 9600.
	BaudRate9600 BaudRate = 9600
	// BaudRate19200 defines that the baudrate is 19200.
	BaudRate19200 BaudRate = 19200
	// BaudRate38400 defines that the baudrate is 38400.
	BaudRate38400 BaudRate = 38400
	// BaudRate57600 defines that the baudrate is 57600.
	BaudRate57600 BaudRate = 57600
	// BaudRate115200 defines that the baudrate is 115200.
	BaudRate115200 BaudRate = 115200
	// BaudRate230400 defines that the baudrate is 230400.
	BaudRate230400 BaudRate = 230400
	// BaudRate460800 defines that the baudrate is 460800.
	BaudRate460800 BaudRate = 460800
	// BaudRate921600 defines that the baudrate is 921600.
	BaudRate921600 BaudRate = 921600
)

// BaudRateParse converts the given string into a Baudrate value.
//
// It returns the corresponding BaudRate constant if the string matches
// a known level name, or an error if the input is invalid.
func BaudRateParse(value string) (BaudRate, error) {
	var ret BaudRate
	var err error
	switch {
	case strings.EqualFold(value, "50"):
		ret = BaudRate50
	case strings.EqualFold(value, "75"):
		ret = BaudRate75
	case strings.EqualFold(value, "110"):
		ret = BaudRate110
	case strings.EqualFold(value, "134"):
		ret = BaudRate134
	case strings.EqualFold(value, "150"):
		ret = BaudRate150
	case strings.EqualFold(value, "200"):
		ret = BaudRate200
	case strings.EqualFold(value, "300"):
		ret = BaudRate300
	case strings.EqualFold(value, "600"):
		ret = BaudRate600
	case strings.EqualFold(value, "1200"):
		ret = BaudRate1200
	case strings.EqualFold(value, "1800"):
		ret = BaudRate1800
	case strings.EqualFold(value, "2400"):
		ret = BaudRate2400
	case strings.EqualFold(value, "4800"):
		ret = BaudRate4800
	case strings.EqualFold(value, "9600"):
		ret = BaudRate9600
	case strings.EqualFold(value, "19200"):
		ret = BaudRate19200
	case strings.EqualFold(value, "38400"):
		ret = BaudRate38400
	case strings.EqualFold(value, "57600"):
		ret = BaudRate57600
	case strings.EqualFold(value, "115200"):
		ret = BaudRate115200
	case strings.EqualFold(value, "230400"):
		ret = BaudRate230400
	case strings.EqualFold(value, "460800"):
		ret = BaudRate460800
	case strings.EqualFold(value, "921600"):
		ret = BaudRate921600
	default:
		err = fmt.Errorf("%w: %q", ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the baudrate.
// It satisfies fmt.Stringer.
func (g BaudRate) String() string {
	var ret string
	switch g {
	case BaudRate50:
		ret = "50"
	case BaudRate75:
		ret = "75"
	case BaudRate110:
		ret = "110"
	case BaudRate134:
		ret = "134"
	case BaudRate150:
		ret = "150"
	case BaudRate200:
		ret = "200"
	case BaudRate300:
		ret = "300"
	case BaudRate600:
		ret = "600"
	case BaudRate1200:
		ret = "1200"
	case BaudRate1800:
		ret = "1800"
	case BaudRate2400:
		ret = "2400"
	case BaudRate4800:
		ret = "4800"
	case BaudRate9600:
		ret = "9600"
	case BaudRate19200:
		ret = "19200"
	case BaudRate38400:
		ret = "38400"
	case BaudRate57600:
		ret = "57600"
	case BaudRate115200:
		ret = "115200"
	case BaudRate230400:
		ret = "230400"
	case BaudRate460800:
		ret = "460800"
	case BaudRate921600:
		ret = "921600"
	}
	return ret
}

// AllBaudrate returns a slice containing all defined baudrate values.
func AllBaudrate() []BaudRate {
	return []BaudRate{
		BaudRate50,
		BaudRate75,
		BaudRate110,
		BaudRate134,
		BaudRate150,
		BaudRate200,
		BaudRate300,
		BaudRate600,
		BaudRate1200,
		BaudRate1800,
		BaudRate2400,
		BaudRate4800,
		BaudRate9600,
		BaudRate19200,
		BaudRate38400,
		BaudRate57600,
		BaudRate115200,
		BaudRate230400,
		BaudRate460800,
		BaudRate921600,
	}
}
