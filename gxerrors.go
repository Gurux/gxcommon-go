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

import (
	"errors"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ErrUnknownEnum Unknown enum value.
var ErrUnknownEnum = errors.New("unknown enum value")

// ErrConnectionClosed means that the connection is closed.
var ErrConnectionClosed = errors.New("connection closed")

// ErrInvalidArgument means that the argument is invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrArgumentOutOfRange means that the argument is out of range.
var ErrArgumentOutOfRange = errors.New("argument out of range")

// ErrBufferTooSmall means that the there is not enought data in the buffer.
var ErrBufferTooSmall = errors.New("buffer too small")

// Initialize error messages.
func init() {
	// --- English (en-US) ---
	message.SetString(language.AmericanEnglish, "error.unknown_enum", "Unknown enum value.")
	message.SetString(language.AmericanEnglish, "error.connection_closed", "Connection closed.")
	message.SetString(language.AmericanEnglish, "error.invalid_argument", "Invalid argument.")
	message.SetString(language.AmericanEnglish, "error.argument_out_of_range", "Argument out of range.")
	message.SetString(language.AmericanEnglish, "error.buffer_too_small", "Buffer too small.")

	// --- German (de) ---
	message.SetString(language.German, "error.unknown_enum", "Unbekannter Enum-Wert.")
	message.SetString(language.German, "error.connection_closed", "Verbindung geschlossen.")
	message.SetString(language.German, "error.invalid_argument", "Ungültiges Argument.")
	message.SetString(language.German, "error.argument_out_of_range", "Argument außerhalb des gültigen Bereichs.")
	message.SetString(language.German, "error.buffer_too_small", "Puffer zu klein.")

	// --- Finnish (fi) ---
	message.SetString(language.Finnish, "error.unknown_enum", "Tuntematon enum-arvo.")
	message.SetString(language.Finnish, "error.connection_closed", "Yhteys on suljettu.")
	message.SetString(language.Finnish, "error.invalid_argument", "Virheellinen argumentti.")
	message.SetString(language.Finnish, "error.argument_out_of_range", "Argumentti on sallitun alueen ulkopuolella.")
	message.SetString(language.Finnish, "error.buffer_too_small", "Puskuri on liian pieni.")

	// --- Swedish (sv) ---
	message.SetString(language.Swedish, "error.unknown_enum", "Okänt enum-värde.")
	message.SetString(language.Swedish, "error.connection_closed", "Anslutningen är stängd.")
	message.SetString(language.Swedish, "error.invalid_argument", "Ogiltigt argument.")
	message.SetString(language.Swedish, "error.argument_out_of_range", "Argumentet är utanför giltigt intervall.")
	message.SetString(language.Swedish, "error.buffer_too_small", "Bufferten är för liten.")

	// --- Spanish (es) ---
	message.SetString(language.Spanish, "error.unknown_enum", "Valor de enumeración desconocido.")
	message.SetString(language.Spanish, "error.connection_closed", "Conexión cerrada.")
	message.SetString(language.Spanish, "error.invalid_argument", "Argumento no válido.")
	message.SetString(language.Spanish, "error.argument_out_of_range", "Argumento fuera de rango.")
	message.SetString(language.Spanish, "error.buffer_too_small", "El búfer es demasiado pequeño.")

	// --- Estonian (et) ---
	message.SetString(language.Estonian, "error.unknown_enum", "Tundmatu enum-väärtus.")
	message.SetString(language.Estonian, "error.connection_closed", "Ühendus suletud.")
	message.SetString(language.Estonian, "error.invalid_argument", "Vigane argument.")
	message.SetString(language.Estonian, "error.argument_out_of_range", "Argument väljaspool lubatud vahemikku.")
	message.SetString(language.Estonian, "error.buffer_too_small", "Puhver on liiga väike.")
}
