package loop

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"github.com/tevino/abool/v2"
)

type Loop struct {
	isRunning *abool.AtomicBool
}

func NewLoop() *Loop {
	return &Loop{abool.New()}
}

func (loop Loop) Start() {
	loop.isRunning.Set()
	for loop.isRunning.IsSet() {
		ctx.EventLoop.Run()
		ctx.PhysicsLoop.Run()
		ctx.RenderLoop.Run()
	}
}

func (loop Loop) Stop() {
	loop.isRunning.UnSet()
}
