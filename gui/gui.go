package gui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mzohreva/MinecraftPuzzle/puzzle"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type grid struct {
	puzzle       *puzzle.Puzzle
	startTexture *sdl.Texture
	goalTexture  *sdl.Texture
	state        puzzle.State
}

func newGrid(p *puzzle.Puzzle, s puzzle.State, r *sdl.Renderer) (*grid, error) {
	startTexture, err := img.LoadTexture(r, "res/start.png")
	if err != nil {
		return nil, err
	}
	goalTexture, err := img.LoadTexture(r, "res/goal.png")
	if err != nil {
		return nil, err
	}
	return &grid{
		puzzle:       p,
		startTexture: startTexture,
		goalTexture:  goalTexture,
		state:        s}, nil
}

func (g *grid) paint(renderer *sdl.Renderer) error {
	if err := renderer.Clear(); err != nil {
		return err
	}
	w, h, err := renderer.GetOutputSize()
	if err != nil {
		return err
	}
	pw, ph := g.puzzle.Width(), g.puzzle.Height()
	padding := 0
	sr := (int(h) - 2*padding) / ph
	sc := (int(w) - 2*padding) / pw
	for r := 0; r < ph; r++ {
		for c := 0; c < pw; c++ {
			cell := g.puzzle.Cell(r, c)
			pos := puzzle.Position{R: r, C: c}
			switch cell {
			case puzzle.Empty:
				renderer.SetDrawColor(255, 255, 255, 255)
			case puzzle.Wall:
				renderer.SetDrawColor(100, 100, 100, 255)
			case puzzle.Minable:
				if g.state.HasMined(pos) {
					renderer.SetDrawColor(150, 255, 150, 255)
				} else {
					renderer.SetDrawColor(0, 200, 0, 255)
				}
			case puzzle.Lava:
				if g.state.HasFilled(pos) {
					renderer.SetDrawColor(255, 150, 150, 255)
				} else {
					renderer.SetDrawColor(200, 0, 0, 255)
				}
			}
			rect := &sdl.Rect{X: int32(padding + c*sc), Y: int32(padding + r*sr), H: int32(sr), W: int32(sc)}
			renderer.FillRect(rect)
			cr, cc := g.state.Position()
			if g.puzzle.IsGoalPosition(r, c) {
				renderer.Copy(g.goalTexture, nil, rect)
			}
			if r == cr && c == cc {
				renderer.Copy(g.startTexture, nil, rect)
			}
			renderer.SetDrawColor(70, 70, 70, 255)
			renderer.DrawRect(rect)
		}
	}
	renderer.Present()
	return nil
}

type solution struct {
	p     *puzzle.Puzzle
	start puzzle.State
	path  []puzzle.Action
}

func (sol solution) show(events <-chan sdl.Event) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)

		width := int32(36 * sol.p.Width())
		height := int32(36 * sol.p.Height())
		w, r, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
		if err != nil {
			errc <- err
		}
		defer w.Destroy()

		g, err := newGrid(sol.p, sol.start, r)
		if err != nil {
			errc <- err
		}
		if err := g.paint(r); err != nil {
			errc <- err
		}

		tick := time.Tick(1000 * time.Millisecond)
		s, i, done := sol.start, 0, false
		for !done {
			select {

			case e := <-events:
				switch e.(type) {
				case *sdl.QuitEvent:
					fmt.Println("Done")
					done = true
				}

			case <-tick:
				if i < len(sol.path) {
					s = s.Successor(sol.p, sol.path[i])
					g.state = s
					if err := g.paint(r); err != nil {
						errc <- err
					}
					i++
				}
			}
		}
	}()

	return errc
}

// Run draws a puzzle using SDL-2
func Run(p *puzzle.Puzzle, start puzzle.State, path []puzzle.Action) error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	// Improve anti-aliasing
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	events := make(chan sdl.Event)
	sol := solution{p: p, start: start, path: path}
	errc := sol.show(events)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}
