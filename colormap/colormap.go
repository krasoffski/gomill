// Package colormap implements a basic color mapping library.
// With givin value, value minimum and value maximum (given range) returns
// RGB color from Blue for minimum value to Red for maximum value.
// Intermediate colors such as green, yellow and more might also be returned for
// values between minimum and maximum (within range).
package colormap

import "fmt"

type Range struct {
	VMin, VMax float64
}

// Correlate returns red, green, blue values represented as float64 in range [0..1].
func (r *Range) Correlate(val float64) (float64, float64, float64) {
	return Correlate(val, r.VMin, r.VMax)
}

// HexStr returns red, green, blue values represented as six hex digits string
// which begins with hash (#).
func (r *Range) HexStr(val float64) string {
	return HexStr(val, r.VMin, r.VMax)
}

// Correlate returns red, green, blue values represented as float64 in range [0..1].
func Correlate(val, vmin, vmax float64) (float64, float64, float64) {

	var (
		vrange  float64 = vmax - vmin
		r, g, b float64 = 1, 1, 1
	)

	if val < vmin {
		val = vmin
	}
	if val > vmax {
		val = vmax
	}

	if val < (vmin + 0.25*vrange) {
		r = 0
		g = 4 * (val - vmin) / vrange
	} else if val < (vmin + 0.5*vrange) {
		r = 0
		b = 1 + 4*(vmin+0.25*vrange)/vrange
	} else if val < (vmin+0.75*vrange)/vrange {
		b = 0
		r = 4 * (val - vmin - 0.5*vrange) / vrange
	} else {
		b = 0
		g = 1 + 4*(vmin+0.75*vrange-val)/vrange
	}
	return r, g, b
}

// HexStr returns red, green, blue values represented as six hex digits string
// which begins with hash (#).
func HexStr(val, vmin, vmax float64) string {
	const m = 0xFF
	r, g, b := Correlate(val, vmin, vmax)
	// TODO: Think about func which returns three uint8 values.
	return fmt.Sprintf("#%X%X%X", uint8(r*m), uint8(g*m), uint8(b*m))
}
