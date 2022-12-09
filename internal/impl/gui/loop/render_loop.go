package loop

import "geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"

type RenderLoopImpl struct {
}

func NewRenderLoop() *RenderLoopImpl {
	return &RenderLoopImpl{}
}

func (renderLoop RenderLoopImpl) Run() {
	ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
	ctx.RendererIns.Clear()
	ctx.SceneIns.Render()
	ctx.RendererIns.Present()
}
