package loop

import "geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"

type RenderLoop struct {
}

func NewRenderLoop() *RenderLoop {
	return &RenderLoop{}
}

func (renderLoop RenderLoop) Run() {
	ctx.Renderer.SetDrawColor(255, 255, 255, 255)
	ctx.Renderer.Clear()
	ctx.CurrentScene.Render()
	ctx.Renderer.Present()
}
