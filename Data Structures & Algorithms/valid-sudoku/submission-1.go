func isValidSudoku(board [][]byte) bool {
	sc := NewSudokuChecker(len(board))
	
	for row, subBoard := range board {
		for col := range subBoard {
			if board[row][col] == '.' {
				continue
			}
			if !(sc.checkValidRow(row,board[row][col]-'1') &&
			sc.checkValidCol(col,board[row][col]-'1') &&
			sc.checkValidQuad(row,col,board[row][col]-'1')) {
				return false
			}
		} 
	}
	return true
}

type SudokuChecker struct {
	rowMap [][]bool
	colMap [][]bool
	quadMap [][]bool
}

func NewSudokuChecker(n int) *SudokuChecker {
	rowMap := make([][]bool,n)
    colMap := make([][]bool,n)
    quadMap := make([][]bool,n)
    for i:=0;i<n;i++ {
        rowMap[i] = make([]bool,n)
        colMap[i] = make([]bool,n)
        quadMap[i] = make([]bool,n)
    }

	return &SudokuChecker{
		rowMap: rowMap,
		colMap: colMap,
		quadMap: quadMap,
	}
}

func (s *SudokuChecker) checkValidRow(row int, char byte) bool {
	if !s.rowMap[row][char] {
		s.rowMap[row][char] = true
		return true
	}

	// not valid
	return false
}

func (s *SudokuChecker) checkValidCol(col int, char byte) bool {
	if !s.colMap[col][char] {
		s.colMap[col][char] = true
		return true
	}

	// not valid
	return false
}

func (s *SudokuChecker) checkValidQuad(row, col int, char byte) bool {
	if !s.quadMap[row/3*3+col/3][char] {
		s.quadMap[row/3*3+col/3][char] = true
		return true
	}

	// not valid
	return false
}