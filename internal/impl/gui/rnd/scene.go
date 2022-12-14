package rnd

import (
	"container/list"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/api"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	. "geniot.com/geniot/go-sdl2-cp-examples/internal/glb"
	"github.com/jakecoffman/cp"
)

type Scene struct {
	renderables *list.List
	Space       *cp.Space
}

func (scene *Scene) GetSpace() *cp.Space {
	return scene.Space
}

func NewScene() *Scene {
	space := cp.NewSpace()
	space.Iterations = 30
	space.SetGravity(cp.Vector{0, 0})
	space.SleepTimeThreshold = 0.5

	_, _, sW, sH := ctx.DeviceIns.GetWindowPosAndSize()
	screenWidth := float64(sW)
	screenHeight := float64(sH)
	screenWidthHalf := float64(screenWidth / 2)
	screenHeightHalf := float64(screenHeight / 2)

	walls := []cp.Vector{
		{-screenWidthHalf, -screenHeightHalf}, {-screenWidthHalf, screenHeightHalf},
		{screenWidthHalf, -screenHeightHalf}, {screenWidthHalf, screenHeightHalf},
		{-screenWidthHalf, -screenHeightHalf}, {screenWidthHalf, -screenHeightHalf},
		{-screenWidthHalf, screenHeightHalf}, {screenWidthHalf, screenHeightHalf},
	}
	for i := 0; i < len(walls)-1; i += 2 {
		shape := space.AddShape(cp.NewSegment(space.StaticBody, walls[i], walls[i+1], 0))
		shape.SetElasticity(1)
		shape.SetFriction(1)
		shape.SetFilter(cp.ShapeFilter{cp.NO_GROUP, ^GRABBABLE_MASK_BIT, ^GRABBABLE_MASK_BIT})
	}

	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(NewBall(space))
	}
	l.PushBack(NewFpsCounter())

	return &Scene{l, space}
}

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}
