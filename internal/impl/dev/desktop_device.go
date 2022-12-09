package dev

import (
	"geniot.com/geniot/go-sdl2-cp-examples/internal/ctx"
	"geniot.com/geniot/go-sdl2-cp-examples/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type DesktopDeviceImpl struct {
}

func (desktopDevice DesktopDeviceImpl) GetJoystickAxis(axis int) int16 {
	return 0 //no joystick on desktop
}

func (desktopDevice DesktopDeviceImpl) Stop() {
	closeCommon()
}

func (desktopDevice DesktopDeviceImpl) ProcessKeyActions() {
	if ctx.PressedKeysCodes.Contains(sdl.K_q) {
		ctx.Loop.Stop()
	}
}

func (desktopDevice DesktopDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return int32(ctx.Config.Get(glb.WINDOW_XPOS_KEY)),
		int32(ctx.Config.Get(glb.WINDOW_YPOS_KEY)),
		glb.SCREEN_WIDTH, glb.SCREEN_HEIGHT
	//int32(ctx.Config.Get(glb.WINDOW_WIDTH_KEY)),
	//int32(ctx.Config.Get(glb.WINDOW_HEIGHT_KEY))
}

func (desktopDevice DesktopDeviceImpl) GetWindowState() uint32 {
	//return ctx.Config.Get(glb.WINDOW_STATE_KEY)
	return sdl.WINDOW_SHOWN
}

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (desktopDevice DesktopDeviceImpl) init() {
	initCommon()
}