package gfx

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

// RoundedRectangleColor draws a rounded rectangle with the specified color.
func RoundedRectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, color *sdl.Color) bool {
	ret, _, _ := purego.SyscallN(roundedRectangleColor, uintptr(unsafe.Pointer(renderer)), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(rad), uintptr(unsafe.Pointer(color)))
	return byte(ret) != 0
}

// RoundedRectangleRGBA draws a rounded rectangle with the specified RGBA color.
func RoundedRectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	ret, _, _ := purego.SyscallN(roundedRectangleRGBA, uintptr(unsafe.Pointer(renderer)), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(rad), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	return byte(ret) != 0
}
