package gxcommon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

// DataType is used to specify the type of data.
type DataType int

const (
	//DataTypeUnknown is unknown.
	DataTypeUnknown DataType = iota
	//DataTypeString is string.
	DataTypeString
	//DataTypeBytes is byte array.
	DataTypeBytes
	//DataTypeByte is byte.
	DataTypeByte
	//DataTypeRune is rune.
	DataTypeRune
	//DataTypeInt16 is int16.
	DataTypeInt16
	//DataTypeInt32 is int32.
	DataTypeInt32
	//DataTypeInt64 is int64.
	DataTypeInt64
	//DataTypeUint16 is uint16.
	DataTypeUint16
	//DataTypeUint32 is uint32.
	DataTypeUint32
	//DataTypeUint64 is uint64.
	DataTypeUint64
)

// GetType returns DataType based on generic type T.
func GetType[T any]() DataType {
	var zero T
	switch any(zero).(type) {
	case string:
		return DataTypeString
	case []byte:
		return DataTypeBytes
	case uint8:
		return DataTypeByte
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

// ToString method is used to convert different types to string.
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

// BytesToAny2 converts a byte slice to a value of specified DataType, interpreting the bytes according to the specified byte order.
func BytesToAny2(b []byte, t DataType, order binary.ByteOrder) (any, error) {
	switch t {
	case DataTypeString:
		return BytesToAny[string](b, order)
	case DataTypeBytes:
		return BytesToAny[[]byte](b, order)
	case DataTypeByte:
		return BytesToAny[uint8](b, order)
	case DataTypeRune:
		return BytesToAny[string](b, order)
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

// BytesToAny converts a byte slice to a value of type T, interpreting the bytes according to the specified byte order.
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

// readFixed lukee kohteeseen täsmälleen sen koon verran.
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

// ToBytes method is used to convert different types to byte slice.
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
