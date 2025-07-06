package gfx

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

var (
	roundedRectangleRGBA func(*sdl.Renderer, int32, int32, int32, int32, int32, uint8, uint8, uint8, uint8) bool
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

	purego.RegisterLibFunc(&roundedRectangleRGBA, lib, "roundedRectangleRGBA")
}
