package main

import (
	"fmt"
	"strconv"
	"strings"
)

// define users in two-player game (temporary)
const userA = "Yin"
const userB = "Yang"
const userAColour = "white"
const userBColour = "black"

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

// define collection of white pieces
var whitePieces = [6]string{
	WP, WR, WN, WB, WQ, WK,
}

// define collection of black pieces
var blackPieces = [6]string{
	BP, BR, BN, BB, BQ, BK,
}

// define mapping between letter and column
var letterToColumnIndex = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

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

// // randomly assign each user either white or black
// func assignColours(userA string, userB string) {
// 	// do stuff here
// }

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

func makeMove(user string, color string) {

	// ask the user what piece to move
	fmt.Println("What piece would you like to move?")
	var pieceCoordinates string
	fmt.Scan(&pieceCoordinates)
	fmt.Println(pieceCoordinates)

	// get piece from coordinates
	pieceCoords_split := strings.Split(pieceCoordinates, "")
	fmt.Println(pieceCoords_split)
	columnCoord := letterToColumnIndex[pieceCoords_split[0]]
	rowCoord, err := strconv.Atoi(pieceCoords_split[1])
	if err == nil {
		// do stuff
	}
	piece := board[rowCoord][columnCoord]
	fmt.Println(piece)

	// check if piece is OK colour

	// if color == white {
	// 	if !checkColor(whitePieces[:], piece) {
	// 		fmt.Println("Wrong side! You're supposed to move %v...", color)
	// 	}
	// }

	// if color == black {
	// 	if !checkColor(blackPieces[:], piece) {
	// 		fmt.Println("Wrong side! You're supposed to move %v...", color)
	// 	}
	// }

	// ask the user where they would like to move it
	fmt.Println("Where would you like to move it?")
	var moveCoordinates string
	fmt.Scan(&moveCoordinates)
	fmt.Println(moveCoordinates)

	// get exact move coordinates
	moveCoords_split := strings.Split(moveCoordinates, "")
	columnCoord_move := letterToColumnIndex[moveCoords_split[0]]
	rowCoord_move, err := strconv.Atoi(moveCoords_split[1])
	if err == nil {
		// do stuff
	}

	// assuming valid move, move piece
	board[rowCoord][columnCoord] = EM
	board[rowCoord_move][columnCoord_move] = piece

	PrintBoard()

}

// func checkColor(arr []string, piece string) {
// 	fmt.Println("hello")
// }

// TEMP: delete after
var turns = 10

func main() {
	PrintBoard()
	var currTurn = userA
	// assignColours(userA, userB)
	// loop infinitely until checkmate, stalemate or user resignation
	for turns > 5 {

		// alternate making moves between players
		if currTurn == userA {
			makeMove(userA, userAColour)
			currTurn = userB
		}

		if currTurn == userB {
			makeMove(userB, userBColour)
			currTurn = userA
		}

		turns = turns - 1
		// if end condition met, exit loop & game
		// else: alternate userA and userB MakeMove
	}
}
