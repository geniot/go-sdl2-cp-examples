package main

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/impl/gui"
)

func main() {
	ctx.ApplicationIns = gui.NewApplication()
	ctx.ApplicationIns.Start()
}
