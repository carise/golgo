package main

import "fmt"

/**
 * Game of Life:
 *
 * 1. LIVE curr_cell: < 2 live neighbors: die
 * 2. LIVE curr_cell: 1 < x <= 3 live neighbors: live
 * 3. LIVE curr_cell: > 3 live neighbors: die
 * 4. DEAD curr_cell: == 3 live neighbors: live
 */

func main() {
    Tick()
}

type Cell struct {
    Alive bool
}

func Tick() {
    PrintCells()
}

func PrintCells() {
    fmt.Printf("Print cells:\n")
}

func (currCell Cell) CountLiveNeighbors() int {
    return 0
}
