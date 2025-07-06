package gfx

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

// RoundedRectangleColor draws a rounded rectangle with the specified color.
func RoundedRectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, color *sdl.Color) bool {
	return roundedRectangleRGBA(renderer, x1, y1, x2, y2, rad, color.R, color.G, color.B, color.A)
}

// RoundedRectangleRGBA draws a rounded rectangle with the specified RGBA color.
func RoundedRectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	return roundedRectangleRGBA(renderer, x1, y1, x2, y2, rad, r, g, b, a)
}
