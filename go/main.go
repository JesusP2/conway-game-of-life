package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	WIDTH      = 25
	HEIGHT     = 25
	aliveColor = "#7D56F4"
	dedColor   = "#000"
)

type Conway struct {
	Width        int
	Height       int
	AliveBlock   string
	DedBlock     string
	previousTick time.Time
	Fps          int64
	State        [][]string
}

func (c *Conway) Init() {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			random := rand.Intn(100)
			if random > 80 {
				c.State[y][x] = c.AliveBlock
			} else {
				c.State[y][x] = c.DedBlock
			}
		}
	}
	c.Draw()
	for {
		c.Tick()
	}
}

func (c *Conway) Tick() {
	currentTime := time.Now()
	requiredTime := 1000 / c.Fps
	if currentTime.Sub(c.previousTick).Milliseconds() < requiredTime {
		return
	}
  c.previousTick = currentTime
	newState := make([][]string, c.Height)
	for y := 0; y < c.Height; y++ {
		newState[y] = make([]string, c.Width)
		for x := 0; x < c.Width; x++ {
			neighbours := c.GetNeighbours(x, y)
			aliveNeighbours := 0
			for _, neighbour := range neighbours {
				if neighbour {
					aliveNeighbours++
				}
			}
			if aliveNeighbours < 2 || aliveNeighbours > 3 {
				newState[y][x] = c.DedBlock
			} else if aliveNeighbours == 3 {
				newState[y][x] = c.AliveBlock
			} else {
				newState[y][x] = c.State[y][x]
			}
		}
	}
	c.State = newState
	c.Draw()
}

func (c *Conway) Draw() {
	fmt.Print("\033[H\033[2J")
	for y := 0; y < c.Height; y++ {
		row := lipgloss.JoinHorizontal(lipgloss.Left, c.State[y]...)
		fmt.Println(row)
	}
}

func (c *Conway) GetNeighbours(x, y int) [8]bool {
	neighbours := [8]bool{}
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			_y := y + i
			_x := x + j
			if (j == 0 && i == 0) || _y < 0 || _x < 0 || _y >= c.Height || _x >= c.Width {
				continue
			}
			neighbours[count] = c.State[_y][_x] == c.AliveBlock
			count++
		}
	}
	return neighbours
}

func NewConway(width, height int, fps int64, aliveBlock string, dedBlock string) *Conway {
	state := make([][]string, height)
	for y := 0; y < height; y++ {
		state[y] = make([]string, width)
	}
	return &Conway{Width: width, Height: height, State: state, Fps: fps, previousTick: time.Now(), AliveBlock: aliveBlock, DedBlock: dedBlock}
}

func main() {
	aliveBlock := lipgloss.NewStyle().
		Background(lipgloss.Color(aliveColor)).
		Width(4).Height(1).Padding(0).Margin(0).Render()
	dedBlock := lipgloss.NewStyle().
		Background(lipgloss.Color(dedColor)).
		Width(4).Height(1).Padding(0).Margin(0).Render()
	conway := NewConway(WIDTH, HEIGHT, 2, aliveBlock, dedBlock)
	conway.Init()
}
