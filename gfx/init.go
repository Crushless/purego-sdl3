package gfx

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
)

var (
	roundedRectangleColor uintptr
	roundedRectangleRGBA  uintptr
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3_gfx.so.1"
	case "windows":
		filename = "SDL3_gfx.dll"
	case "darwin":
		filename = "libSDL3_gfx.dylib"
	}

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&roundedRectangleColor, lib, "roundedRectangleColor")
	purego.RegisterLibFunc(&roundedRectangleRGBA, lib, "roundedRectangleRGBA")
}
