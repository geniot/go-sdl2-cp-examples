package rnd

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type FpsCounter struct {
	startTicks    uint64
	frameCount    uint64
	currentSecond uint64
	currentFPS    uint64
}

func NewFpsCounter() *FpsCounter {
	return &FpsCounter{sdl.GetTicks64(), 0, sdl.GetTicks64() / 1000, 0}
}

func (fpsCounter *FpsCounter) Render() {
	fpsCounter.frameCount += 1
	currentTicks := sdl.GetTicks64()
	fps := 1000 / ((currentTicks - fpsCounter.startTicks) / fpsCounter.frameCount)

	sec := currentTicks / 1000
	if sec > fpsCounter.currentSecond {
		fpsCounter.currentFPS = fps
		fpsCounter.frameCount = 0
		fpsCounter.startTicks = currentTicks
		fpsCounter.currentSecond = sec
	}
	drawText("FPS: "+strconv.FormatInt(int64(fpsCounter.currentFPS), 10), 10, 10, glb.COLOR_BLACK)
}
