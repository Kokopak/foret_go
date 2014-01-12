package main

import (
	"github.com/banthar/Go-SDL/sdl"
)

func main() {

	grid := genGrid()

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	var screen = sdl.SetVideoMode(W*16, H*16, 32, sdl.RESIZABLE)

	sdl.EnableUNICODE(1)

	sdl.WM_SetCaption("Feu de fôret", "")

	running := true
	launch := false

	if sdl.GetKeyName(270) != "[+]" {
		panic("GetKeyName broken")
	}

	tree := sdl.Load("arbre.png")
	fire := sdl.Load("feu.png")

	screen.FillRect(&sdl.Rect{0, 0, W * 16, H * 16}, sdl.MapRGB(screen.Format, 255, 255, 255))
	screen.Flip()

	for running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch e := ev.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				println("")
				println(e.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keysym.Sym)))

				if e.Keysym.Sym == 27 {
					running = false
				}

				if e.Type == sdl.KEYDOWN {
					if e.Keysym.Sym == 32 {
						launch = true
					} else if e.Keysym.Sym == 114 {
						screen.FillRect(&sdl.Rect{0, 0, W * 16, H * 16}, sdl.MapRGB(screen.Format, 255, 255, 255))
						grid = genGrid()
						launch = false
					}
				}

			}
		}
		// Draw ici
		for row := 1; row < W-1; row++ {
			for col := 1; col < H-1; col++ {
				cell := grid[row][col]
				rect := &sdl.Rect{int16(row * 16), int16(col * 16), 16, 16}
				if cell == TREE {
					screen.Blit(rect, tree, nil)
				} else if cell == FIRE {
					screen.Blit(rect, fire, nil)
				} else if cell == EMPTY {
					screen.FillRect(rect, sdl.MapRGB(screen.Format, 255, 255, 255))
				}
			}
		}
		if launch {
			grid = evolve(grid)
		}
		screen.Flip()
		sdl.Delay(1 / 60)
	}
}
