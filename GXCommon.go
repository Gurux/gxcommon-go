package gxcommon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"golang.org/x/text/language"
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

// DataType specifies a supported value type for conversion helpers.
// It is used by the generic conversion utilities to indicate the kind of
// data being handled.
type DataType int

const (
	// DataTypeUnknown is an unknown type.
	DataTypeUnknown DataType = iota
	// DataTypeString represents string values.
	DataTypeString
	// DataTypeBytes represents a byte slice ([]byte).
	DataTypeBytes
	// DataTypeUint8 represents uint8 (byte).
	DataTypeUint8
	// DataTypeInt8 represents int8.
	DataTypeInt8
	// DataTypeInt16 represents int16.
	DataTypeInt16
	// DataTypeInt32 represents int32.
	DataTypeInt32
	// DataTypeInt64 represents int64.
	DataTypeInt64
	// DataTypeUint16 represents uint16.
	DataTypeUint16
	// DataTypeUint32 represents uint32.
	DataTypeUint32
	// DataTypeUint64 represents uint64.
	DataTypeUint64
)

var currentLang atomic.Value
var (
	mu   sync.RWMutex
	subs = map[chan language.Tag]struct{}{}
)

// String provides a human-readable name for the DataType.
// It satisfies fmt.Stringer and is primarily used in logging/examples.
func (dt DataType) String() string {
	switch dt {
	case DataTypeUnknown:
		return "Unknown"
	case DataTypeString:
		return "String"
	case DataTypeBytes:
		return "Bytes"
	case DataTypeUint8:
		return "Uint8"
	case DataTypeInt8:
		return "Int8"
	case DataTypeInt16:
		return "Int16"
	case DataTypeInt32:
		return "Int32"
	case DataTypeInt64:
		return "Int64"
	case DataTypeUint16:
		return "Uint16"
	case DataTypeUint32:
		return "Uint32"
	case DataTypeUint64:
		return "Uint64"
	default:
		return "Unknown"
	}
}

// GetType returns the DataType that corresponds to T.
// AllDataTypes returns the list of all defined DataType values.
// It is mainly useful for validation and documentation purposes.
func AllDataTypes() []DataType {
	return []DataType{
		DataTypeUnknown,
		DataTypeString,
		DataTypeBytes,
		DataTypeUint8,
		DataTypeInt8,
		DataTypeInt16,
		DataTypeInt32,
		DataTypeInt64,
		DataTypeUint16,
		DataTypeUint32,
		DataTypeUint64,
	}
}

// GetType returns the DataType that corresponds to T.
func GetType[T any]() DataType {
	var zero T
	switch any(zero).(type) {
	case string:
		return DataTypeString
	case []byte:
		return DataTypeBytes
	case uint8:
		return DataTypeUint8
	case int16:
		return DataTypeInt16
	case int32:
		return DataTypeInt32
	case int64:
		return DataTypeInt64
	case uint16:
		return DataTypeUint16
	case uint32:
		return DataTypeUint32
	case uint64:
		return DataTypeUint64
	default:
		return DataTypeUnknown
	}
}

// ToString converts a supported value to a string.
//
// Byte slices are returned as space-separated uppercase hex.
func ToString(value any) (string, error) {
	var str string
	switch x := value.(type) {
	case []byte:
		str = ToHex(x)
	default:
		str = fmt.Sprint(value)
	}
	return str, nil
}

// BytesToAny2 converts b to a value of the given DataType.
func BytesToAny2(b []byte, t DataType, order binary.ByteOrder) (any, error) {
	switch t {
	case DataTypeString:
		return BytesToAny[string](b, order)
	case DataTypeBytes:
		return BytesToAny[[]byte](b, order)
	case DataTypeUint8:
		return BytesToAny[uint8](b, order)
	case DataTypeInt8:
		return BytesToAny[int8](b, order)
	case DataTypeInt16:
		return BytesToAny[int16](b, order)
	case DataTypeInt32:
		return BytesToAny[int32](b, order)
	case DataTypeInt64:
		return BytesToAny[int64](b, order)
	case DataTypeUint16:
		return BytesToAny[uint16](b, order)
	case DataTypeUint32:
		return BytesToAny[uint32](b, order)
	case DataTypeUint64:
		return BytesToAny[uint64](b, order)
	default:
		return nil, ErrInvalidArgument
	}
}

// BytesToAny converts b to type T using the given byte order.
func BytesToAny[T any](b []byte, order binary.ByteOrder) (T, error) {
	var zero T

	switch any(zero).(type) {
	case string:
		// koko puskuri UTF-8:na
		return any(string(b)).(T), nil

	case []byte:
		cp := make([]byte, len(b))
		copy(cp, b)
		return any(cp).(T), nil

	case uint8:
		if len(b) < 1 {
			return zero, fmt.Errorf("buffer too short for uint8")
		}
		return any(b[0]).(T), nil
	case int16:
		var v int16
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil
	case int32:
		var v int32
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil
	case int64:
		var v int64
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil

	case uint16:
		var v uint16
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil
	case uint32:
		var v uint32
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil
	case uint64:
		var v uint64
		if err := readFixed(b, order, &v); err != nil {
			return zero, err
		}
		return any(v).(T), nil
	}

	return zero, fmt.Errorf("unsupported target type %T", zero)
}

// readFixed is used to read value to the output.
func readFixed(b []byte, order binary.ByteOrder, out any) error {
	var need int
	switch out.(type) {
	case *int16, *uint16:
		need = 2
	case *int32, *uint32:
		need = 4
	case *int64, *uint64:
		need = 8
	default:
		return fmt.Errorf("unsupported fixed-size type %T", out)
	}
	if len(b) < need {
		return io.ErrUnexpectedEOF
	}
	return binary.Read(bytes.NewReader(b[:need]), order, out)
}

// ToBytes converts a supported value to bytes using the given byte order.
func ToBytes(v any, order binary.ByteOrder) ([]byte, error) {
	if v == nil {
		return []byte{}, nil
	}
	switch x := v.(type) {
	case []byte:
		return x, nil
	case string:
		return []byte(x), nil
	case uint8:
		return []byte{any(x).(byte)}, nil
	case int16:
		return writeFixed(order, x)
	case int32:
		return writeFixed(order, x)
	case int64:
		return writeFixed(order, x)
	case int:
		return writeFixed(order, int64(x))
	case uint16:
		return writeFixed(order, x)
	case uint32:
		return writeFixed(order, x)
	case uint64:
		return writeFixed(order, x)
	case uint:
		return writeFixed(order, uint64(x))
	}
	return nil, fmt.Errorf("ToBytes: unsupported type %T", v)
}

func writeFixed[T any](order binary.ByteOrder, x T) ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, order, x); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// SetLanguage sets the package-wide language used by internal status and error
// messages.
//
// The value is stored atomically and is safe to update concurrently. If a
// translation for the selected language is not available, message lookups use
// the default catalog fallback behavior.
func SetLanguage(lang language.Tag) {
	prev := Language()
	if prev == lang {
		return
	}

	currentLang.Store(lang)
	notify(lang)
}

// Language returns the currently configured package-wide language.
//
// If no language has been configured with SetLanguage, Language returns
// language.AmericanEnglish.
func Language() language.Tag {
	v := currentLang.Load()
	if v == nil {
		return language.AmericanEnglish
	}
	return v.(language.Tag)
}

// Subscribe registers a listener for package-wide language changes.
//
// The returned buffered channel immediately receives the current language and
// then receives future values when SetLanguage changes the language.
func Subscribe() chan language.Tag {
	ch := make(chan language.Tag, 1)
	mu.Lock()
	subs[ch] = struct{}{}
	mu.Unlock()
	ch <- Language()
	return ch
}

// Unsubscribe removes a language listener created by Subscribe and closes it.
//
// Calling Unsubscribe for a channel that is not subscribed is a no-op.
func Unsubscribe(ch chan language.Tag) {
	mu.Lock()
	if _, ok := subs[ch]; ok {
		delete(subs, ch)
		close(ch)
	}
	mu.Unlock()
}

// notify notifies subscribes when language is changed.
func notify(lang language.Tag) {
	mu.RLock()
	list := make([]chan language.Tag, 0, len(subs))
	for ch := range subs {
		list = append(list, ch)
	}
	mu.RUnlock()
	for _, ch := range list {
		select {
		case ch <- lang:
		default:
		}
	}
}
