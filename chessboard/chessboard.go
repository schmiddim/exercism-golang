package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	ctr := 0
	for _, field := range cb[file] {
		if field {
			ctr += 1
		}
	}
	return ctr
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	if rank < 1 || rank > 8 {
		return 0
	}
	ctr := 0
	for _, file := range cb {
		if file[rank-1] {
			ctr += 1
		}
	}
	return ctr
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	ctr := 0
	for _, fields := range cb {
		for _, file := range fields {
			if file {
				ctr += 1
			}
		}
	}
	return ctr
}
