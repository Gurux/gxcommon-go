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
	// Baud50 defines that the baudrate is 50.
	Baud50 BaudRate = 50
	// Baud75 defines that the baudrate is 75.
	Baud75 BaudRate = 75
	// Baud110 defines that the baudrate is 110.
	Baud110 BaudRate = 110
	// Baud134 defines that the baudrate is 134.
	Baud134 BaudRate = 134
	// Baud150 defines that the baudrate is 150.
	Baud150 BaudRate = 150
	// Baud200 defines that the baudrate is 200.
	Baud200 BaudRate = 200
	// Baud300 defines that the baudrate is 300.
	Baud300 BaudRate = 300
	// Baud600 defines that the baudrate is 600.
	Baud600 BaudRate = 600
	// Baud1200 defines that the baudrate is 1200.
	Baud1200 BaudRate = 1200
	// Baud1800 defines that the baudrate is 1800.
	Baud1800 BaudRate = 1800
	// Baud2400 defines that the baudrate is 2400.
	Baud2400 BaudRate = 2400
	// Baud4800 defines that the baudrate is 4800.
	Baud4800 BaudRate = 4800
	// Baud9600 defines that the baudrate is 9600.
	Baud9600 BaudRate = 9600
	// Baud19200 defines that the baudrate is 19200.
	Baud19200 BaudRate = 19200
	// Baud38400 defines that the baudrate is 38400.
	Baud38400 BaudRate = 38400
	// Baud57600 defines that the baudrate is 57600.
	Baud57600 BaudRate = 57600
	// Baud115200 defines that the baudrate is 115200.
	Baud115200 BaudRate = 115200
	// Baud230400 defines that the baudrate is 230400.
	Baud230400 BaudRate = 230400
	// Baud460800 defines that the baudrate is 460800.
	Baud460800 BaudRate = 460800
	// Baud921600 defines that the baudrate is 921600.
	Baud921600 BaudRate = 921600
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
		ret = Baud50
	case strings.EqualFold(value, "75"):
		ret = Baud75
	case strings.EqualFold(value, "110"):
		ret = Baud110
	case strings.EqualFold(value, "134"):
		ret = Baud134
	case strings.EqualFold(value, "150"):
		ret = Baud150
	case strings.EqualFold(value, "200"):
		ret = Baud200
	case strings.EqualFold(value, "300"):
		ret = Baud300
	case strings.EqualFold(value, "600"):
		ret = Baud600
	case strings.EqualFold(value, "1200"):
		ret = Baud1200
	case strings.EqualFold(value, "1800"):
		ret = Baud1800
	case strings.EqualFold(value, "2400"):
		ret = Baud2400
	case strings.EqualFold(value, "4800"):
		ret = Baud4800
	case strings.EqualFold(value, "9600"):
		ret = Baud9600
	case strings.EqualFold(value, "19200"):
		ret = Baud19200
	case strings.EqualFold(value, "38400"):
		ret = Baud38400
	case strings.EqualFold(value, "57600"):
		ret = Baud57600
	case strings.EqualFold(value, "115200"):
		ret = Baud115200
	case strings.EqualFold(value, "230400"):
		ret = Baud230400
	case strings.EqualFold(value, "460800"):
		ret = Baud460800
	case strings.EqualFold(value, "921600"):
		ret = Baud921600
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
	case Baud50:
		ret = "50"
	case Baud75:
		ret = "75"
	case Baud110:
		ret = "110"
	case Baud134:
		ret = "134"
	case Baud150:
		ret = "150"
	case Baud200:
		ret = "200"
	case Baud300:
		ret = "300"
	case Baud600:
		ret = "600"
	case Baud1200:
		ret = "1200"
	case Baud1800:
		ret = "1800"
	case Baud2400:
		ret = "2400"
	case Baud4800:
		ret = "4800"
	case Baud9600:
		ret = "9600"
	case Baud19200:
		ret = "19200"
	case Baud38400:
		ret = "38400"
	case Baud57600:
		ret = "57600"
	case Baud115200:
		ret = "115200"
	case Baud230400:
		ret = "230400"
	case Baud460800:
		ret = "460800"
	case Baud921600:
		ret = "921600"
	}
	return ret
}

// AllBaudrate returns a slice containing all defined baudrate values.
func AllBaudrate() []BaudRate {
	return []BaudRate{
		Baud50,
		Baud75,
		Baud110,
		Baud134,
		Baud150,
		Baud200,
		Baud300,
		Baud600,
		Baud1200,
		Baud1800,
		Baud2400,
		Baud4800,
		Baud9600,
		Baud19200,
		Baud38400,
		Baud57600,
		Baud115200,
		Baud230400,
		Baud460800,
		Baud921600,
	}
}
