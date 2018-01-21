package gui

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellHeight = 40
	cellWidth  = 40
)

type runner interface {
	run(<-chan sdl.Event) <-chan error
}

func runGUI(r runner) error {
	runtime.LockOSThread()
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	// Improve anti-aliasing
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	events := make(chan sdl.Event)
	errc := r.run(events)

	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}
