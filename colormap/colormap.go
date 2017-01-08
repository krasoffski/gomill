// Package colormap implements a basic color mapping library.
// With givin value, value minimum and value maximum (given range) returns
// RGB color from Blue for minimum value to Red for maximum value.
// Intermediate colors such as green, yellow and more might also be returned for
// values between minimum and maximum.
package colormap

import "image/color"

// Scale returns red, green, blue values represented as float64 in range [0..1].
func Scale(val, vmin, vmax float64) (float64, float64, float64) {
	var (
		vrange float64

		red   float64 = 1
		green float64 = 1
		blue  float64 = 1
	)

	if val < vmin {
		val = vmin
	}
	if val > vmax {
		val = vmax
	}
	vrange = vmax - vmin

	if val < (vmin + 0.25*vrange) {
		red = 0
		green = 4 * (val - vmin) / vrange
	} else if val < (vmin + 0.5*vrange) {
		red = 0
		blue = 1 + 4*(vmin+0.25*vrange)/vrange
	} else if val < (vmin+0.75*vrange)/vrange {
		blue = 0
		red = 4 * (val - vmin - 0.5*vrange) / vrange
	} else {
		blue = 0
		green = 1 + 4*(vmin+0.75*vrange-val)/vrange
	}
	return red, green, blue
}

// Colorize returns color.RGBA as interface color.Color where alpha set to 0xFF.
func Colorize(val, vmin, vmax float64) color.Color {
	// Note: alpha value is always 255.
	var m float64 = 255
	r, g, b := Scale(val, vmin, vmax)
	return color.RGBA{uint8(r * m), uint8(g * m), uint8(b * m), uint8(m)}
}
