package gui

import (
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

func newGrid(p *puzzle.Puzzle, start puzzle.State, r *sdl.Renderer) (*grid, error) {
	startTexture, err := img.LoadTexture(r, "res/start.png")
	if err != nil {
		return nil, err
	}
	goalTexture, err := img.LoadTexture(r, "res/goal.png")
	if err != nil {
		return nil, err
	}
	return &grid{puzzle: p, startTexture: startTexture, goalTexture: goalTexture, state: start}, nil
}

func (g *grid) paint(renderer *sdl.Renderer) error {
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
	return nil
}

// DrawPuzzle draws a puzzle using SDL-2
func DrawPuzzle(p *puzzle.Puzzle, start puzzle.State, path []puzzle.Action) error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(int32(36*p.Width()), int32(36*p.Height()), sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer w.Destroy()

	// Improve anti-aliasing
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	g, err := newGrid(p, start, r)
	if err != nil {
		return err
	}

	r.Clear()
	if err := g.paint(r); err != nil {
		return err
	}
	r.Present()

	go func() {
		s := start
		for _, a := range path {
			time.Sleep(200 * time.Millisecond)
			s = s.Successor(p, a)
			g.state = s
			r.Clear()
			g.paint(r)
			r.Present()
		}
	}()

	runtime.LockOSThread()
	done := false
	for !done {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			done = true
		}
	}
	return nil
}
