package life

import (
	"math/rand"
)

type Board struct {
	Width int
	Height int
	Cells [][]bool
}
func NewCells(width int, height int) [][]bool {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return cells
}
func NewBoard(width int, height int) Board {
	return Board{
		width,
		height,
		NewCells(width, height),
	}
}

func getNextState(currentState bool, score int) bool {
	return score == 3 || score == 2 && currentState
}

func (b Board) setState(x int, y int, alive bool) {
	b.Cells[y][x] = alive
}

func (b Board) Seed() {
	for i := 0; i < (b.Width * b.Height / 4); i++ {
		b.setState(rand.Intn(b.Width), rand.Intn(b.Height), true)
	}
}

func (b Board) alive(x int, y int) bool {
	if(x < 0 || x >= b.Width) {
		return false
	}
	if(y < 0 || y >= b.Height) {
		return false
	}
	if(b.Cells[y] == nil) {
		return false
	}
	return b.Cells[y][x]
}

func (b Board) neighborsAlive(x int, y int) int {
	score := 0
	for _,dx := range []int{-1,0,1} {
		for _,dy := range []int{-1,0,1} {
			if (dx == 0 && dy == 0) {
				continue
			}
			if (b.alive(x + dx, y + dy)) {
				score++
			}
		}
	}
	return score
}

func (b Board) nextState(x int, y int) bool {
	return getNextState(
		b.alive(x, y),
		b.neighborsAlive(x, y),
	)
}

func (b Board) getNextCells() [][]bool {
	nextCells := NewCells(b.Width, b.Height)
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			nextCells[y][x] = b.nextState(x,y)
		}
	}
	return nextCells
}

func (b Board) GoToNext() bool {
	nextCells := b.getNextCells()
	changed := false
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			changed = changed || nextCells[y][x] != b.alive(x, y)
			b.setState(x,y, nextCells[y][x])
		}
	}
	return changed
}

func (b Board) String() string {
	var bt byte
	buf := make([]byte, 0, (b.Width+1)*b.Height)

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			bt = ' '
			if(b.alive(x,y)) {
				bt = '*'
			}
			buf = append(buf, bt)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}
