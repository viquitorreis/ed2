package main

import "fmt"

func main() {
	graph := NewGraph().GenBoard(8, 8)
	graph.Move(6, 6)
	graph.PrintBoard()
	graph.Move(4, 5)
	fmt.Printf("%v\n", graph.GetValidMoves())
	graph.PrintBoard()
}

type Graph struct {
	Board     [][]int // 0 - vazio, 1 - cavalo, 2 - visitado
	knightPos *Position
}

type Position struct {
	row, col int
}

type IGraph interface {
	GenBoard(rows, cols int) *Graph
	Move(row, col int) bool
	IsMoveValid(row, col int) bool
	GetValidMoves() []Position
}

func NewGraph() IGraph {
	return &Graph{
		Board: nil,
	}
}

func (g *Graph) GenBoard(rows, cols int) *Graph {
	g.Board = make([][]int, rows)
	for i := range g.Board {
		g.Board[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g.Board[i][j] = 0
		}
	}

	return g
}

func (g *Graph) GetValidMoves() []Position {
	var moves []Position
	validKnightMoves := []Position{
		{-2, -1},
		{-2, 1},
		{-1, -2},
		{-1, 2},
		{1, -2},
		{1, 2},
		{2, -1},
		{2, 1},
	}

	for _, move := range validKnightMoves {
		pos := &Position{
			row: g.knightPos.row + move.row,
			col: g.knightPos.col + move.col,
		}

		if !g.IsMoveValid(pos.row, pos.col) {
			continue
		}

		moves = append(moves, *pos)
	}

	return moves
}

func (g *Graph) IsMoveValid(row, col int) bool {
	if row >= 0 && col >= 0 && row <= 7 && col <= 7 {
		return true
	}

	return false
}

func (g *Graph) IsMoveOnRange(row, col int) bool {
	validMoves := g.GetValidMoves()
	for _, move := range validMoves {
		if move.col == col && move.row == row {
			return true
		}
	}

	return false
}

func (g *Graph) Move(row, col int) bool {
	if !g.IsMoveValid(row, col) {
		return false
	}

	if g.knightPos == nil {
		g.Board[row][col] = 1
		g.knightPos = &Position{row, col}
	}

	if g.IsMoveOnRange(row, col) {
		// marcar previous como 2
		g.Board[g.knightPos.row][g.knightPos.col] = 2
		g.Board[row][col] = 1
		g.knightPos = &Position{row, col}
	}

	return true
}

func (g *Graph) PrintBoard() {
	for i := range g.Board {
		for j := range g.Board[i] {
			fmt.Printf("%d ", g.Board[i][j])
		}
		fmt.Printf("\n")
	}
}
