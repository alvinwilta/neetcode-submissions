func isValidSudoku(board [][]byte) bool {
	sc := NewSudokuChecker()
	
	for row, subBoard := range board {
		for col := range subBoard {
			if !(sc.checkValidRow(row,board[row][col]) &&
			sc.checkValidCol(col,board[row][col]) &&
			sc.checkValidQuad(row,col,board[row][col])) {
				return false
			}
		} 
	}
	return true
}

type SudokuChecker struct {
	rowMap []map[byte]bool
	colMap []map[byte]bool
	quadMap [][]map[byte]bool
}

func NewSudokuChecker() *SudokuChecker {
	rowMap := make([]map[byte]bool,9)
	colMap := make([]map[byte]bool,9)
	for i:=0;i<9;i++ {
		rowMap[i] = make(map[byte]bool)
		colMap[i] = make(map[byte]bool)
	}
	quadMap := make([][]map[byte]bool,3)
	for i := range quadMap {
		quadMap[i] = make([]map[byte]bool,3)
		for j := range quadMap {
			quadMap[i][j] = make(map[byte]bool)
		}
	}

	return &SudokuChecker{
		rowMap: rowMap,
		colMap: colMap,
		quadMap: quadMap,
	}
}

func (s *SudokuChecker) checkValidRow(row int, char byte) bool {
	if char == '.' {
		return true
	}
	if !s.rowMap[row][char] {
		s.rowMap[row][char] = true
		return true
	}

	// not valid
	return false
}

func (s *SudokuChecker) checkValidCol(col int, char byte) bool {
	if char == '.' {
		return true
	}

	if !s.colMap[col][char] {
		s.colMap[col][char] = true
		return true
	}

	// not valid
	return false
}

func (s *SudokuChecker) checkValidQuad(row, col int, char byte) bool {
	if char == '.' {
		return true
	}
	if !s.quadMap[row/3][col/3][char] {
		s.quadMap[row/3][col/3][char] = true
		return true
	}

	// not valid
	return false
}