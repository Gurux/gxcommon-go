// Package gxcommon_test holds examples for the gxcommon package.
package gxcommon_test

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/Gurux/gxcommon-go"
)

// ExampleGetType demonstrates the generic GetType helper.
func ExampleGetType() {
	fmt.Println(gxcommon.GetType[string]())
	fmt.Println(gxcommon.GetType[[]byte]())
	fmt.Println(gxcommon.GetType[int16]())
	// Output:
	// String
	// Bytes
	// Int16
}

// ExampleToString shows how arbitrary values are converted.
func ExampleToString() {
	// byte slices are formatted as space-separated hex
	b := []byte{0x01, 0xAB, 0xFF}
	str, _ := gxcommon.ToString(b)
	fmt.Println(str)

	// any other value uses fmt.Sprint
	fmt.Println(gxcommon.ToString(123))

	// Output:
	// 01 AB FF
	// 123 <nil>
}

// Example demonstrates parsing and stringifying of common enums (and
// the DataType utility).
func Example() {
	rate, _ := gxcommon.BaudRateParse("9600")
	fmt.Println(rate)
	fmt.Println(gxcommon.ParityEven)
	fmt.Println(gxcommon.TraceLevelVerbose)
	fmt.Println(gxcommon.AllDataTypes()[1]) // skip Unknown
	// Output:
	// 9600
	// Even
	// VERBOSE
	// String
}

// ExampleToHex demonstrates the public hex formatter that lives next to
// ReceiveEventArgs (it is not specific to that type).
func ExampleToHex() {
	fmt.Println(gxcommon.ToHex([]byte{0, 0x1F, 0xA0}))
	// Output:
	// 00 1F A0
}

// ExampleLanguage shows the language hooks provided by the package.
func ExampleLanguage() {
	fmt.Println(gxcommon.Language())
	gxcommon.SetLanguage(language.German)
	fmt.Println(gxcommon.Language())
	// Output:
	// en-US
	// de
}
