package main

import (
    "fmt"
    "math/rand"
    "time"
)

/**
 * Conway's Game of Life:
 *
 * 1. LIVE curr_cell: < 2 live neighbors: die
 * 2. LIVE curr_cell: 1 < x <= 3 live neighbors: live
 * 3. LIVE curr_cell: > 3 live neighbors: die
 * 4. DEAD curr_cell: == 3 live neighbors: live
 */

type Cell struct {
    CurrAlive bool
    NextAlive bool
}

type Grid struct {
    Size int
    Cells [][]Cell
}

func main() {
    row := 20
    col := 30
    seed := -1 
    
    if seed > -1 {
        rand.Seed(int64(seed))
    } else {
        rand.Seed(time.Now().UnixNano())
    }

    grid := &Grid{}
    grid.Init(row, col)

    for grid.Tick() {
        time.Sleep(time.Second)
    }
}

func (grid *Grid) Init(row int, col int) {
    cells := make([][]Cell, row)
    for i := range cells {
        cells[i] = make([]Cell, col)
    }
    grid.Size = row * col
    grid.Cells = cells
    grid.Seed()
}

func (grid *Grid) Seed() {
    pct_alive := int(float32(grid.Size) * 0.30)
    for i := range grid.Cells {
        row := grid.Cells[i]
        for j := range row {
            r := rand.Intn(grid.Size)
            row[j].CurrAlive, row[j].NextAlive = r < pct_alive, r < pct_alive
        }
    }
    grid.PrintGrid()
}

func (grid *Grid) PrintGrid() {
    fmt.Println("----------------------------")
    for i := range grid.Cells {
        row := grid.Cells[i]
        for j := range row {
            if row[j].NextAlive {
                fmt.Printf(" x ")
            } else {
                fmt.Printf(" . ")
            }
            row[j].CurrAlive = row[j].NextAlive
        }
        fmt.Printf("\n")
    }
}

func (grid *Grid) Tick() bool {
    changed := false
    for i := range grid.Cells {
        row := grid.Cells[i]
        for j := range row {
            count := grid.CountLiveNeighbors(i, j)
            if row[j].CurrAlive {
                row[j].NextAlive = count == 2 || count == 3
            } else {
                row[j].NextAlive = count == 3
            }
            changed = changed || row[j].CurrAlive != row[j].NextAlive
        }
    }
    grid.PrintGrid()
    return changed
}

func (grid Grid) CountLiveNeighbors(row int, col int) int {
    x_l := col - 1
    x_r := col + 1
    y_t := row - 1
    y_b := row + 1

    if x_l < 0 {
        x_l = 0
    }

    if x_r >= len(grid.Cells[row]) {
        x_r = len(grid.Cells[row]) - 1
    }

    if y_t < 0 {
        y_t = 0
    }

    if y_b >= len(grid.Cells) {
        y_b = len(grid.Cells) - 1
    }

    count := 0

    for i := y_t; i <= y_b; i++ {
        for j := x_l; j <= x_r; j++ {
            if i == row && j == col {
                continue
            }
            if grid.Cells[i][j].CurrAlive {
                count++
            }
        }
    }
    return count
}
