package gui

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/glb"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/impl/dev"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/impl/gui/loop"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/impl/gui/rnd"
	"geniot.com/geniot/go-sdl2-cp-examples/resources"
	"github.com/veandco/go-sdl2/ttf"
)

type ApplicationImpl struct {
}

func NewApplication() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (app *ApplicationImpl) Start() {
	ctx.DeviceIns = dev.NewDevice()
	ctx.ConfigIns = NewConfig()
	ctx.WindowIns = NewWindow()

	ctx.LoopIns = loop.NewLoop()
	ctx.EventLoopIns = loop.NewEventLoop()
	ctx.PhysicsLoopIns = loop.NewPhysicsLoop()
	ctx.RenderLoopIns = loop.NewRenderLoop()

	ctx.SceneIns = rnd.NewScene()

	ctx.FontIns, _ = ttf.OpenFontRW(resources.GetResource(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)

	ctx.LoopIns.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) Stop() {
	ctx.FontIns.Close()
	ctx.DeviceIns.Stop()
}
