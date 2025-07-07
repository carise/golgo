package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jroimartin/gocui"
)

/**
 * Conway's Game of Life:
 *
 * 1. LIVE curr_cell: < 2 live neighbors: die
 * 2. LIVE curr_cell: 1 < x <= 3 live neighbors: live
 * 3. LIVE curr_cell: > 3 live neighbors: die
 * 4. DEAD curr_cell: == 3 live neighbors: live
 */

type CellState int

const (
	CellDead CellState = iota
	CellAlive
)

type Cell struct {
	State CellState
}

type Grid struct {
	View      *gocui.View
	Size      int
	Cells     [][]Cell
	Iteration int
}

var grid *Grid

var speed int
var cols int
var rows int
var initialPctAlive float64

func main() {
	flag.IntVar(&speed, "speed", 300, "tick speed in millis")
	flag.IntVar(&cols, "cols", 50, "width of grid")
	flag.IntVar(&rows, "rows", 30, "height of grid")
	flag.Float64Var(&initialPctAlive, "initialPctAlive", 0.25, "initial percentage of alive cells")
	flag.Parse()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	xMax := (cols * 2) + 5
	yMax := rows + 5

	vInit, err1 := g.SetView("initParams", 5, 2, xMax, 4)
	if err1 != nil && err1 != gocui.ErrUnknownView {
		return err1
	}

	vGrid, err2 := g.SetView("grid", 5, 5, xMax, yMax)
	if err2 != nil && err2 != gocui.ErrUnknownView {
		return err2
	}
	
	vIters, err3 := g.SetView("iterations", 5, yMax+1, xMax, yMax+3)
	if err3 != nil && err3 != gocui.ErrUnknownView {
		return err3
	}

	if grid == nil {
		fmt.Fprintf(vInit, "cols=%d, rows=%d, initialPctAlive=%f", cols, rows, initialPctAlive)

		grid = &Grid{View: vGrid}
		grid.Init(cols, rows)

		// Start simulation in separate goroutine
		go runSimulation(g, vIters)
	}

	return nil
}

func runSimulation(g *gocui.Gui, iterView *gocui.View) {
	ticker := time.NewTicker(time.Duration(speed) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		changed := grid.Tick()
		g.Update(func(g *gocui.Gui) error {
			iterView.Clear()
			fmt.Fprint(iterView, "Iterations: "+fmt.Sprint(grid.Iteration))

			if !changed {
				fmt.Fprint(iterView, " (no change)")
			}
			return nil
		})
		if !changed {
			break
		}
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (grid *Grid) Init(cols int, rows int) {
	cells := make([][]Cell, cols)
	for i := range cells {
		cells[i] = make([]Cell, rows)
	}
	grid.Size = cols * rows
	grid.Cells = cells
	pct_alive := int(float64(grid.Size) * initialPctAlive)
	for i := range grid.Cells {
		row := grid.Cells[i]
		for j := range row {
			r := rand.Intn(grid.Size)
			if r < pct_alive {
				row[j].State = CellAlive
			} else {
				row[j].State = CellDead
			}
		}
	}
	grid.printGrid()
}

func (grid *Grid) printGrid() {
	grid.View.Clear()
	for i := range rows {
		for j := range cols {
			if grid.Cells[j][i].State == CellAlive {
				fmt.Fprint(grid.View, " @")
			} else {
				fmt.Fprint(grid.View, " .")
			}
		}
		fmt.Fprint(grid.View, " \n")
	}
}

func (grid *Grid) Tick() bool {
	changed := false
	grid.printGrid()
	for i := range grid.Cells {
		col := grid.Cells[i]
		for j := range col {
			count := grid.countLiveNeighbors(i, j)
			if col[j].State == CellAlive {
				if count == 2 || count == 3 {
					col[j].State = CellAlive
				} else {
					col[j].State = CellDead
					changed = true
				}
			} else if count == 3 {
				col[j].State = CellAlive
				changed = true
			}
		}
	}
	grid.Iteration++
	return changed
}

func (grid Grid) countLiveNeighbors(col int, row int) int {
	x_l := col - 1
	x_r := col + 1
	y_t := row - 1
	y_b := row + 1

	if x_l < 0 {
		x_l = 0
	}

	if x_r >= cols {
		x_r = cols - 1
	}

	if y_t < 0 {
		y_t = 0
	}

	if y_b >= rows {
		y_b = rows - 1
	}

	count := 0

	for i := x_l; i <= x_r; i++ {
		for j := y_t; j <= y_b; j++ {
			if i == col && j == row {
				continue
			}
			if grid.Cells[i][j].State == CellAlive {
				count++
			}
		}
	}
	return count
}
