package main

import "fmt"

// define chess pieces 
const (
	E = iota // empty
	WP // white pawn
	WR // white rook
	WN // white knight
	WB // white bishop
	WQ // white queen
	WK // white king
	BP // black pawn
	BR // black rook
	BN // black knight
	BB // black bishop
	BQ // black queen
	BK // black king
)

// define a 2D array representing the initial chessboard
initialBoard = [][]int{
	{WR, WN, WB, WQ, WK, WB, WN, WR},
    {WP, WP, WP, WP, WP, WP, WP, WP},
    {E, E, E, E, E, E, E, E},
    {E, E, E, E, E, E, E, E},
    {E, E, E, E, E, E, E, E},
    {E, E, E, E, E, E, E, E},
    {BP, BP, BP, BP, BP, BP, BP, BP},
    {BR, BN, BB, BQ, BK, BB, BN, BR},
}

// Print the initial chessboard
func PrintBoard() {
	for _, row := range initial_board {
		for _, piece := range row {
			fmt.Printf("%2d ", piece)
		}
		fmt.Println()
	}
}