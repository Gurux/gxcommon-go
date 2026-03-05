// Package gxcommon provides reusable helpers, configuration types, and
// interfaces that underpin the Gurux media libraries.
//
// The package is intentionally lightweight and exposes:
//
//   * the IGXMedia interface and associated event argument types used by all
//     media implementations (serial, TCP, USB, etc.)
//   * common serial‑port enums (BaudRate, Parity, StopBits) with parsing and
//     String helpers
//   * tracing and state enums (TraceLevel, TraceTypes, MediaState) plus
//     utilities for working with them
//   * data conversion helpers (ToBytes, ToString, GetType) and a small set of
//     errors used throughout the framework
//   * simple language subscription helpers used for localized messages
//
// The package is documented with examples so that `go doc` or `pkg.go.dev` can
// show concrete usage.  See `gxcommon_examples_test.go` for sample code.
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
