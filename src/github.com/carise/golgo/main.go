package main

import (
    "fmt"
)

/**
 * Game of Life:
 *
 * 1. LIVE curr_cell: < 2 live neighbors: die
 * 2. LIVE curr_cell: 1 < x <= 3 live neighbors: live
 * 3. LIVE curr_cell: > 3 live neighbors: die
 * 4. DEAD curr_cell: == 3 live neighbors: live
 */

type Cell struct {
    Alive bool
}

type Grid struct {
    Cells [][]Cell
}

func main() {
    row := 10
    col := 10
    grid := &Grid{}
    grid.Init(row, col)
    grid.Tick()
}

func (grid *Grid) Init(row int, col int) {
    cells := make([][]Cell, row)
    for i := range cells {
        cells[i] = make([]Cell, col)
    }
    grid.Cells = cells
}

func (grid Grid) Tick() {
    grid.PrintGrid()
}

func (grid Grid) PrintGrid() {
    fmt.Printf("Print cells:\n")
    for i := 0; i < len(grid.Cells); i++ {
        row := grid.Cells[i]
        for j := 0; j < len(row); j++ {
            if row[j].Alive {
                fmt.Printf(" x ")
            } else {
                fmt.Printf(" . ")
            }
        }
        fmt.Printf("\n")
    }
}

func (grid Grid) SeedGrid() {
}

func (currCell Cell) CountLiveNeighbors() int {
    return 0
}
