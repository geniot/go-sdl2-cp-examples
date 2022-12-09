package rnd

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"geniot.com/geniot/go-sdl2-cp-examples/resources"
	"github.com/jakecoffman/cp"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
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

	_, _, sW, sH := ctx.DeviceIns.GetWindowPosAndSize()

	radius := 25.0
	body := space.AddBody(cp.NewBody(10, cp.MomentForCircle(10, 0, radius, cp.Vector{})))
	body.SetPosition(cp.Vector{
		float64(rand.Intn(int(sW)) - int(sW)/2),
		float64(rand.Intn(int(sH)) - int(sH)/2)})
	body.SetVelocity(float64(sW/2), float64(sH/2))

	shape := space.AddShape(cp.NewCircle(body, radius, cp.Vector{}))
	shape.SetElasticity(1)
	shape.SetFriction(0.5)

	return Ball{space, surface.W, surface.H, txt, body}
}

func (ball Ball) Render() {
	_, _, sW, sH := ctx.DeviceIns.GetWindowPosAndSize()
	dstRect := sdl.Rect{
		int32(ball.body.Position().X) - ball.width/2 + sW/2,
		int32(ball.body.Position().Y) - ball.height/2 + sH/2,
		ball.width, ball.height}

	degrees := math.Abs(2 * math.Pi * ball.body.Angle())

	ctx.RendererIns.CopyEx(ball.texture, nil,
		&dstRect, math.Mod(degrees, 360),
		&sdl.Point{ball.width / 2, ball.height / 2}, sdl.FLIP_NONE)
}
