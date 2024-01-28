package main
 
import "fmt"

// define chess pieces 
const (
	E = iota // empty
	WP  // white pawn
	WR // white rook
	WN // white knight
	WB // white bishop
	WQ // white queen
	WK // white king
	BP // black pawn
	BR  // black rook
	BN  // black knight
	BB  // black bishop
	BQ // black queen
	BK  // black king
)

// define a 2D array representing the initial chessboard
var initialBoard = [][]int{
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
    for _, row := range initialBoard {
        for _, piece := range row {
            fmt.Printf("%2d ", piece)
        }
        fmt.Println()
    }
}

var startPos string 
var endPos string

func MakeMove() {

    // get move from user 
    fmt.Print("Which piece to move: ")
    fmt.Scanln(&startPos)
    fmt.Print("Where to move: ")
    fmt.Scanln(&endPos)

    // get X and Y coordinates (startPos)
    startPos_split := []rune(startPos)
    X_coord_start := startPos_split[0]
    Y_coord_start := startPos_split[1]

    // get piece from startPos
    var piece int

    if X_coord_start == 'E' {
        X_coord_start = 4 
    }

    piece = initialBoard[X_coord_start][Y_coord_start]
    Print(piece)

    // get X and Y coordinates (endPos)
    endPos_split := []rune(endPos)
    X_coord_end:= endPos_split[0]
    Y_coord_end := endPos_split[1]
    fmt.Print(X_coord_end)
    fmt.Print(Y_coord_end)
}

func main() {
    PrintBoard()
    MakeMove()
}





