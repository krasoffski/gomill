// Package htcmap implements a basic hot to color ramp functions.
// With givin value, value minimum and value maximum (given range) returns
// RGB color from Blue for minimum value to Red for maximum value.
// Intermediate colors such as green, yellow and more might also be returned for
// values between minimum and maximum (within range).
package htcmap

import "fmt"

type Range struct {
	VMin, VMax float64
}

func (r *Range) AsFloat(val float64) (float64, float64, float64) {
	return AsFloat(val, r.VMin, r.VMax)
}

func (r *Range) AsUInt8(val float64) (uint8, uint8, uint8) {
	return AsUInt8(val, r.VMin, r.VMax)
}

func (r *Range) AsStr(val float64) string {
	return AsStr(val, r.VMin, r.VMax)
}

// AsFloat returns red, green, blue values each as float64 in range [0..1].
func AsFloat(val, vmin, vmax float64) (float64, float64, float64) {
	return htc(val, vmin, vmax)
}

// AsUInt8 returns red, green, blue values each as uint8 in range [0..255].
func AsUInt8(val, vmin, vmax float64) (uint8, uint8, uint8) {
	const m = 0xFF
	r, g, b := AsFloat(val, vmin, vmax)
	return uint8(r * m), uint8(g * m), uint8(b * m)
}

// AsStr returns red, green, blue values represented as six hex digits string
// which begins with hash (#). It looks like a HTML RGB color #00FF00.
func AsStr(val, vmin, vmax float64) string {
	r, g, b := AsUInt8(val, vmin, vmax)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Stop points can be modified/adjusted for better correlation.
func htc(val, vmin, vmax float64) (float64, float64, float64) {

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
