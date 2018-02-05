package gui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mzohreva/MinecraftPuzzle/puzzle"
	"github.com/veandco/go-sdl2/sdl"
)

type designer struct {
	width, height int
	puzzle        *puzzle.Puzzle
}

func (d *designer) run(events <-chan sdl.Event) <-chan error {
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

		width := int32(cellWidth * d.width)
		height := int32(cellHeight * d.height)
		w, r, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
		reportError(err)
		defer w.Destroy()

		w.SetTitle(fmt.Sprintf("Designer %v x %v", d.height, d.width))

		reportError(r.Clear())
		reportError(r.SetDrawColor(100, 100, 255, 255))
		reportError(r.FillRect(&sdl.Rect{X: 0, Y: 0, W: width, H: height}))
		r.Present()
		time.Sleep(time.Second)

		g, err := newGrid(d.puzzle, nil, r)
		reportError(err)
		reportError(g.paint(r))

		done := false
		for !done {
			select {

			case e := <-events:
				switch event := e.(type) {
				case *sdl.QuitEvent:
					done = true
				case *sdl.MouseButtonEvent:
					if event.State == sdl.RELEASED {
						col, row := int(event.X/cellWidth), int(event.Y/cellHeight)
						t := d.puzzle.Cell(row, col)
						d.puzzle.SetCell(row, col, nextCellType(t))
						reportError(g.paint(r))
					}
				case *sdl.KeyboardEvent:
					if event.State == sdl.RELEASED {
						switch event.Keysym.Sym {
						case sdl.K_UP:
							d.puzzle.MoveGoal(-1, 0)
						case sdl.K_DOWN:
							d.puzzle.MoveGoal(+1, 0)
						case sdl.K_LEFT:
							d.puzzle.MoveGoal(0, -1)
						case sdl.K_RIGHT:
							d.puzzle.MoveGoal(0, +1)
						case sdl.K_w:
							d.puzzle.MoveStart(-1, 0)
						case sdl.K_s:
							d.puzzle.MoveStart(+1, 0)
						case sdl.K_a:
							d.puzzle.MoveStart(0, -1)
						case sdl.K_d:
							d.puzzle.MoveStart(0, +1)
						}
						reportError(g.paint(r))
					}
				}
			}
		}
	}()

	return errc
}

func nextCellType(t puzzle.CellType) puzzle.CellType {
	switch t {
	case puzzle.Empty:
		return puzzle.Wall
	case puzzle.Wall:
		return puzzle.Minable
	case puzzle.Minable:
		return puzzle.Lava
	case puzzle.Lava:
		return puzzle.Empty
	}
	return puzzle.Empty
}

// DesignPuzzle starts GUI for designing puzzles
func DesignPuzzle(height, width int, p *puzzle.Puzzle) (*puzzle.Puzzle, error) {
	if p == nil {
		p = puzzle.NewEmptyPuzzle(height, width)
	} else {
		height, width = p.Height(), p.Width()
	}
	d := designer{width: width, height: height, puzzle: p}
	err := runGUI(&d)
	if err != nil {
		return nil, err
	}
	return d.puzzle, nil // TODO
}
