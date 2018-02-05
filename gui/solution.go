package gui

import (
	"runtime"
	"time"

	"github.com/mzohreva/MinecraftPuzzle/puzzle"
	"github.com/veandco/go-sdl2/sdl"
)

type solution puzzle.Solution

func (sol solution) run(events <-chan sdl.Event) <-chan error {
	errc := make(chan error)
	reportError := func(e error) {
		if e != nil {
			errc <- e
			runtime.Goexit()
		}
	}

	go func() {
		defer close(errc)

		runtime.LockOSThread()

		p := sol.Problem.GetPuzzle()
		width := int32(cellWidth * p.Width())
		height := int32(cellHeight * p.Height())
		w, r, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
		reportError(err)
		defer w.Destroy()

		w.SetTitle(puzzle.Solution(sol).String())

		reportError(r.Clear())
		reportError(r.SetDrawColor(100, 100, 255, 255))
		reportError(r.FillRect(&sdl.Rect{X: 0, Y: 0, W: width, H: height}))
		r.Present()
		time.Sleep(time.Second)

		start := puzzle.Solution(sol).Start()
		g, err := newGrid(p, &start, r)
		reportError(err)
		if len(sol.Path) == 0 {
			g.sad = true
		}
		reportError(g.paint(r))

		tick := time.Tick(250 * time.Millisecond)
		s, i, done := start, 0, false
		for !done {
			select {

			case e := <-events:
				switch e.(type) {
				case *sdl.QuitEvent:
					done = true
				}

			case <-tick:
				if i < len(sol.Path) {
					s = s.Successor(p, sol.Path[i])
					g.state = &s
					reportError(g.paint(r))
					i++
				}
			}
		}
	}()

	return errc
}

// ShowSolution shows a puzzle and A* solution using SDL-2
func ShowSolution(sol puzzle.Solution) error {
	s := solution(sol)
	return runGUI(&s)
}
