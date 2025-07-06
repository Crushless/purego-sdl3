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

// RoundedBoxColor draws a rounded box with the specified color.
func RoundedBoxColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, color *sdl.Color) bool {
	return roundedBoxRGBA(renderer, x1, y1, x2, y2, rad, color.R, color.G, color.B, color.A)
}

// RoundedBoxRGBA draws a rounded box with the specified RGBA color.
func RoundedBoxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	return roundedBoxRGBA(renderer, x1, y1, x2, y2, rad, r, g, b, a)
}
