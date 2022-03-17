package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	sum := 0
	for _, val := range cb[rank] {
		if val {
			sum++
		}
	}
	return sum
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	sum := 0
	if file < 0 || file > 8 {
		return 0
	}
	for _, val := range cb {
		if val[file-1] {
			sum++
		}
	}
	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	sum := 0
	for _, rank := range cb {
		for range rank {
			sum++
		}
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	sum := 0
	for _, rank := range cb {
		for _, val := range rank {
			if val {
				sum++
			}
		}
	}
	return sum
}
