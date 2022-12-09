package rnd

import (
	"container/list"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/api"
)

type Scene struct {
	renderables *list.List
}

func NewScene() *Scene {
	l := list.New()
	l.PushBack(NewFpsCounter())
	return &Scene{l}
}

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}
