package main

import (
	"math/rand"
)

const (
	W     = 80
	H     = 50
	PROBA = 0.45
)
const (
	EMPTY = 1
	TREE  = 2
	FIRE  = 3
)

func genGrid() (grid [W][H]int) {
	for row := 1; row < W-1; row++ {
		for col := 1; col < H-1; col++ {
			random := rand.Float64()
			if random < PROBA {
				grid[row][col] = TREE
			} else {
				grid[row][col] = EMPTY
			}
		}
	}
	grid[randInt(1, W-1)][randInt(1, H-1)] = FIRE

	return grid
}

func evolve(grid [W][H]int) (newGrid [W][H]int) {
	newGrid = grid
	for row := 1; row < W-1; row++ {
		for col := 1; col < H-1; col++ {
			cell := grid[row][col]
			if cell == TREE {
				if row-1 >= 0 && row+1 < W && col-1 >= 0 && col+1 < H {
					if grid[row-1][col-1] == FIRE || grid[row-1][col] == FIRE || grid[row-1][col+1] == FIRE ||
						grid[row][col-1] == FIRE || grid[row][col+1] == FIRE ||
						grid[row+1][col-1] == FIRE || grid[row+1][col] == FIRE || grid[row+1][col+1] == FIRE {
						newGrid[row][col] = FIRE
					}
				}
			} else if cell == FIRE {
				newGrid[row][col] = EMPTY
			}
		}
	}
	return newGrid
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
