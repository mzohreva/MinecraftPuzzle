package gui

import (
	"github.com/mzohreva/MinecraftPuzzle/puzzle"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type grid struct {
	puzzle         *puzzle.Puzzle
	playerTextures []*sdl.Texture
	goalTexture    *sdl.Texture
	state          *puzzle.State
	sad            bool
}

func newGrid(p *puzzle.Puzzle, s *puzzle.State, r *sdl.Renderer) (*grid, error) {
	neutral, err := img.LoadTexture(r, "res/player-neutral.png")
	if err != nil {
		return nil, err
	}
	happy, err := img.LoadTexture(r, "res/player-happy.png")
	if err != nil {
		return nil, err
	}
	sad, err := img.LoadTexture(r, "res/player-sad.png")
	if err != nil {
		return nil, err
	}
	goalTexture, err := img.LoadTexture(r, "res/goal.png")
	if err != nil {
		return nil, err
	}
	playerTextures := make([]*sdl.Texture, 3)
	playerTextures[0] = neutral
	playerTextures[1] = happy
	playerTextures[2] = sad
	return &grid{
		puzzle:         p,
		playerTextures: playerTextures,
		goalTexture:    goalTexture,
		state:          s}, nil
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
				if g.state != nil && g.state.HasMined(pos) {
					renderer.SetDrawColor(150, 255, 150, 255)
				} else {
					renderer.SetDrawColor(0, 200, 0, 255)
				}
			case puzzle.Lava:
				if g.state != nil && g.state.HasFilled(pos) {
					renderer.SetDrawColor(255, 150, 150, 255)
				} else {
					renderer.SetDrawColor(200, 0, 0, 255)
				}
			}
			rect := &sdl.Rect{
				X: int32(padding + c*sc),
				Y: int32(padding + r*sr),
				H: int32(sr),
				W: int32(sc)}
			imgRect := &sdl.Rect{
				X: rect.X + 2,
				Y: rect.Y + 2,
				H: rect.H - 4,
				W: rect.W - 4}
			renderer.FillRect(rect)
			if g.puzzle.IsGoalPosition(r, c) {
				renderer.Copy(g.goalTexture, nil, imgRect)
			}
			if g.puzzle.IsStartPosition(r, c) {
				renderer.SetDrawColor(50, 50, 100, 255)
				renderer.DrawRect(&sdl.Rect{
					X: int32(padding + c*sc + sc/5),
					Y: int32(padding + r*sr + sr/5),
					H: int32(sr * 3 / 5),
					W: int32(sc * 3 / 5)})
			}
			if g.state != nil {
				cr, cc := g.state.Position()
				if r == cr && c == cc {
					if g.sad {
						renderer.Copy(g.playerTextures[2], nil, imgRect)
					} else if g.puzzle.IsGoalPosition(r, c) {
						renderer.Copy(g.playerTextures[1], nil, imgRect)
					} else {
						renderer.Copy(g.playerTextures[0], nil, imgRect)
					}
				}
			}
			renderer.SetDrawColor(70, 70, 70, 255)
			renderer.DrawRect(rect)
		}
	}
	renderer.Present()
	return nil
}
