// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"

	"math"
	"math/rand"
)

// Define epsilon constant for floating point checks
const epsilon = 1E-3

func abs(x float32) float32 {
	if x < 0.0 {
		return -x
	}
	return x
}

func get_float_quantization_error(value, resolution float32, mantis int) float32 {
	// Calculate the exponent for the value given the mantis and value resolution
	value_m := value / (resolution * float32(mantis))
	value_exp := math.Ceil(math.Log2(float64(value_m)))

	// Calculate the worst quantization error by assuming the
	// mantis to be off by one
	q := math.Pow(2, value_exp)

	// Scale back the quantization error with the given resolution
	return float32(q) / resolution
}

func genUint8Slice(start, end, step uint8) []uint8 {
	if step <= 0 || end < start {
		return make([]uint8, 0)
	}
	s := make([]uint8, 0, (end-start)/step)
	for start <= end {
		s = append(s, uint8(start))
		start += step
	}
	return s
}

// Test DPT 1.001 (Switch) with values within range
func TestDPT_1001(t *testing.T) {
	var buf []byte
	var src, dst DPT_1001

	for _, value := range []bool{true, false} {
		src = DPT_1001(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 1.002 (Bool) with values within range
func TestDPT_1002(t *testing.T) {
	var buf []byte
	var src, dst DPT_1002

	for _, value := range []bool{true, false} {
		src = DPT_1002(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 1.003 (Enable) with values within range
func TestDPT_1003(t *testing.T) {
	var buf []byte
	var src, dst DPT_1003

	for _, value := range []bool{true, false} {
		src = DPT_1003(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 1.008 (OpenClose) with values within range
func TestDPT_1008(t *testing.T) {
	var buf []byte
	var src, dst DPT_1008

	for _, value := range []bool{true, false} {
		src = DPT_1008(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 1.009 (OpenClose) with values within range
func TestDPT_1009(t *testing.T) {
	var buf []byte
	var src, dst DPT_1009

	for _, value := range []bool{true, false} {
		src = DPT_1009(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 1.010 (Start) with values within range
func TestDPT_1010(t *testing.T) {
	var buf []byte
	var src, dst DPT_1010

	for _, value := range []bool{true, false} {
		src = DPT_1010(value)
		if bool(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if bool(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 3.007 (Increase/Decrease by value) with values within range
func TestDPT_3007(t *testing.T) {
	var buf []byte
	var src, dst DPT_3007

	vs := genUint8Slice(0, 7, 1)
	for _, c := range []bool{true, false} {
		for _, v := range vs {
			fmt.Printf("C: %+v; V: %+v\n", c, v)
			src = DPT_3007{Increase: c, Value: v}

			if src.Increase != c {
				t.Errorf("Assignment of control value \"%v\" failed! Has value \"%t\".", c, src.Increase)
			}
			if src.Value != v {
				t.Errorf("Assignment of value \"%v\" failed! Has value \"%d\".", v, src.Value)
			}
			buf = src.Pack()
			fmt.Printf("Byte: %+v\n", buf)
			dst.Unpack(buf)
			if dst.Increase != c {
				t.Errorf("Wrong control value \"%t\" after pack/unpack! Original control value was \"%v\".", dst.Increase, c)
			}
			if dst.Value != v {
				t.Errorf("Wrong value \"%v\" after pack/unpack! Original value was \"%d\".", dst.Value, v)
			}
		}
	}
}

// Test DPT 5.001 (Scaling) with values within range
func TestDPT_5001(t *testing.T) {
	var buf []byte
	var src, dst DPT_5001

	// Calculate the quantization error we expect
	const Q = float32(100) / 255

	for i := 1; i <= 10; i++ {
		value := rand.Float32()

		// Scale the random number to the given range
		value *= 100

		// Pack and unpack to test value
		src = DPT_5001(value)
		if abs(float32(src)-value) > epsilon {
			t.Errorf("Assignment of value \"%v\" failed for source of type DPT_5001! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if abs(float32(dst)-value) > (Q + epsilon) {
			t.Errorf("Value \"%s\" after pack/unpack above quantization noise! Original value was \"%v\", noise is \"%f\"", dst, value, Q)
		}
	}
}

// Test DPT 5.003 (Angle) with values within range
func TestDPT_5003(t *testing.T) {
	var buf []byte
	var src, dst DPT_5003

	// Calculate the quantization error we expect
	const Q = float32(360) / 255

	for i := 1; i <= 10; i++ {
		value := rand.Float32()

		// Scale the random number to the given range
		value *= 360

		// Pack and unpack to test value
		src = DPT_5003(value)
		if abs(float32(src)-value) > epsilon {
			t.Errorf("Assignment of value \"%v\" failed for source of type DPT_5003! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if abs(float32(dst)-value) > (Q + epsilon) {
			t.Errorf("Value \"%s\" after pack/unpack above quantization noise! Original value was \"%v\", noise is \"%f\"", dst, value, Q)
		}
	}
}

// Test DPT 9.001 (Temperature) with values within range
func TestDPT_9001(t *testing.T) {
	var buf []byte
	var src, dst DPT_9001

	for i := 1; i <= 10; i++ {
		value := rand.Float32()

		// Scale the random number to the given range
		value *= 670760 - -273
		value += -273

		// Calculate the quantization error we expect
		Q := get_float_quantization_error(value, 0.01, 2047)

		// Pack and unpack to test value
		src = DPT_9001(value)
		if abs(float32(src)-value) > epsilon {
			t.Errorf("Assignment of value \"%v\" failed for source of type DPT_9001! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if abs(float32(dst)-value) > (Q + epsilon) {
			t.Errorf("Value \"%s\" after pack/unpack above quantization noise! Original value was \"%v\", noise is \"%f\"", dst, value, Q)
		}
	}
}

// Test DPT 9.004 (Illumination) with values within range
func TestDPT_9004(t *testing.T) {
	var buf []byte
	var src, dst DPT_9004

	for i := 1; i <= 10; i++ {
		value := rand.Float32()

		// Scale the random number to the given range
		value *= 670760

		// Calculate the quantization error we expect
		Q := get_float_quantization_error(value, 0.01, 2047)

		// Pack and unpack to test value
		src = DPT_9004(value)
		if abs(float32(src)-value) > epsilon {
			t.Errorf("Assignment of value \"%v\" failed for source of type DPT_9004! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if abs(float32(dst)-value) > (Q + epsilon) {
			t.Errorf("Value \"%s\" after pack/unpack above quantization noise! Original value was \"%v\", noise is \"%f\"", dst, value, Q)
		}
	}
}

// Test DPT 12.001 (Unsigned counter) with values within range
func TestDPT_12001(t *testing.T) {
	var buf []byte
	var src, dst DPT_12001

	for i := 1; i <= 10; i++ {
		value := rand.Uint32()

		// Pack and unpack to test value
		src = DPT_12001(value)
		if uint32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed for source of type DPT_12001! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if uint32(dst) != value {
			t.Errorf("Value \"%s\" after pack/unpack different from Original value. Was \"%v\"", dst, value)
		}
	}
}

// Test DPT 13.001 (counter pulses)
func TestDPT_13001(t *testing.T) {
	var buf []byte
	var src, dst DPT_13001

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13001(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13001(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13001(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.002 (flow rate)
func TestDPT_13002(t *testing.T) {
	var buf []byte
	var src, dst DPT_13002

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13002(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13002(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13002(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.010 (active energy)
func TestDPT_13010(t *testing.T) {
	var buf []byte
	var src, dst DPT_13010

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13010(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13010(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13010(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.011 (apparant energy)
func TestDPT_13011(t *testing.T) {
	var buf []byte
	var src, dst DPT_13011

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13011(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13011(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13011(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.012 (reactive energy)
func TestDPT_13012(t *testing.T) {
	var buf []byte
	var src, dst DPT_13012

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13012(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13012(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13012(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.013 (active energy (kWh))
func TestDPT_13013(t *testing.T) {
	var buf []byte
	var src, dst DPT_13013

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13013(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13013(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13013(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.014 (apparant energy (kVAh))
func TestDPT_13014(t *testing.T) {
	var buf []byte
	var src, dst DPT_13014

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13014(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13014(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13014(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}

// Test DPT 13.015 (reactive energy (kVARh))
func TestDPT_13015(t *testing.T) {
	var buf []byte
	var src, dst DPT_13015

	// Corner cases
	for _, value := range []int32{math.MinInt32, math.MaxInt32} {
		src = DPT_13015(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Positive
	for i := 1; i <= 10; i++ {
		value := rand.Int31()

		src = DPT_13015(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}

	// Negative
	for i := 1; i <= 10; i++ {
		value := -rand.Int31()

		src = DPT_13015(value)
		if int32(src) != value {
			t.Errorf("Assignment of value \"%v\" failed! Has value \"%s\".", value, src)
		}
		buf = src.Pack()
		dst.Unpack(buf)
		if int32(dst) != value {
			t.Errorf("Wrong value \"%s\" after pack/unpack! Original value was \"%v\".", dst, value)
		}
	}
}
