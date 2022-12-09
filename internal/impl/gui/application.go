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
	ctx.Device = dev.NewDevice()
	ctx.Config = NewConfig()
	ctx.Window = NewWindow()

	ctx.Loop = loop.NewLoop()
	ctx.EventLoop = loop.NewEventLoop()
	ctx.PhysicsLoop = loop.NewPhysicsLoop()
	ctx.RenderLoop = loop.NewRenderLoop()

	ctx.CurrentScene = rnd.NewScene()

	ctx.Font, _ = ttf.OpenFontRW(resources.GetResource(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)

	ctx.Loop.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) Stop() {
	ctx.Font.Close()
	ctx.Device.Stop()
}
