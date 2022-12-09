package ctx

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/api"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	ApplicationIns api.IApplication
	WindowIns      api.IWindow
	DeviceIns      api.IDevice
	ConfigIns      api.IConfig

	LoopIns        api.IStartable
	EventLoopIns   api.IRunnable
	PhysicsLoopIns api.IRunnable
	RenderLoopIns  api.IRunnable

	SceneIns api.IScene

	RendererIns            *sdl.Renderer
	FontIns                *ttf.Font
	PressedKeysCodesSetIns mapset.Set[sdl.Keycode] = mapset.NewSet[sdl.Keycode]()
)
