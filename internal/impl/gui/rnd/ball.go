package rnd

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/glb"
	"geniot.com/geniot/go-sdl2-cp-examples/resources"
	"github.com/jakecoffman/cp"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	space   *cp.Space
	width   int32
	height  int32
	texture *sdl.Texture
	body    *cp.Body
}

func NewBall(space *cp.Space) Ball {
	surface, _ := img.LoadRW(resources.GetResource("ball.png"), true)
	defer surface.Free()
	txt, err := ctx.RendererIns.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}

	radius := 25.0
	body := space.AddBody(cp.NewBody(10, cp.MomentForCircle(10, 0, radius, cp.Vector{})))
	body.SetPosition(cp.Vector{0, 0})
	body.SetVelocity(0, 300)

	shape := space.AddShape(cp.NewCircle(body, radius, cp.Vector{}))
	shape.SetElasticity(1)
	shape.SetFriction(0.0)

	return Ball{space, surface.W, surface.H, txt, body}
}

func (ball Ball) Render() {
	dstRect := sdl.Rect{
		int32(ball.body.Position().X) + glb.SCREEN_WIDTH_HALF,
		int32(ball.body.Position().Y) + glb.SCREEN_HEIGHT_HALF,
		ball.width, ball.height}

	ctx.RendererIns.Copy(ball.texture, nil, &dstRect)
}
