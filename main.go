package main

import "fmt"

// define users in two-player game (temporary)
const userA = "Yin"
const userB = "Yang"

// define checkmate, stalemate, and quit status
var checkmateStatus = false
var stalemateStatus = false
var userResignedStatus = false

// define chess pieces
const (
	EM = "EM" // empty
	WP = "WP" // white pawn
	WR = "WR" // white rook
	WN = "WN" // white knight
	WB = "WB" // white bishop
	WQ = "WQ" // white queen
	WK = "WK" // white king
	BP = "BP" // black pawn
	BR = "BR" // black rook
	BN = "BN" // black knight
	BB = "BB" // black bishop
	BQ = "BQ" // black queen
	BK = "BK" // black king
)

// define mapping between letter and column
const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
)

// define a 2D array representation of the (initial) chessboard
var board = [8][8]string{
	{WR, WN, WB, WQ, WK, WB, WN, WR},
	{WP, WP, WP, WP, WP, WP, WP, WP},
	{EM, EM, EM, EM, EM, EM, EM, EM},
	{EM, EM, EM, EM, EM, EM, EM, EM},
	{EM, EM, EM, EM, EM, EM, EM, EM},
	{EM, EM, EM, EM, EM, EM, EM, EM},
	{BP, BP, BP, BP, BP, BP, BP, BP},
	{BR, BN, BB, BQ, BK, BB, BN, BR},
}

// Print the chessboard
func PrintBoard() {
	for _, row := range board {
		for _, piece := range row {
			fmt.Printf("%v ", piece)
		}
		fmt.Println()
	}
}

// randomly assign each user either white or black
func AssignColours() {
	// do stuff here
}

// var startPos string
// var endPos string

// func MakeMove() {

// 	// get move from user
// 	fmt.Print("Which piece to move: ")
// 	fmt.Scanln(&startPos)
// 	fmt.Print("Where to move: ")
// 	fmt.Scanln(&endPos)

// 	// get X and Y coordinates (startPos)
// 	startPos_split := []rune(startPos)
// 	X_coord_start := startPos_split[0]
// 	Y_coord_start := startPos_split[1]
// 	fmt.Print(X_coord_start)
// 	fmt.Print(Y_coord_start)

// 	// get piece from startPos
// 	var piece string

// 	if X_coord_start == 'E' {
// 		X_coord_start = 4
// 	}

// 	piece = initialBoard[X_coord_start][Y_coord_start]
// 	fmt.Print(piece)

// 	// get X and Y coordinates (endPos)
// 	endPos_split := []rune(endPos)
// 	X_coord_end := endPos_split[0]
// 	Y_coord_end := endPos_split[1]
// 	fmt.Print(X_coord_end)
// 	fmt.Print(Y_coord_end)
// }

func main() {
	AssignColours()
	// loop repeatedly until checkmate, stalemate or user resignation
	for {
		// if end condition met, exit loop & game
		// else: alternate userA and userB MakeMove
		PrintBoard()
	}
}
