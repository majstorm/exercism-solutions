// Package matrix
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	rows_len := 0
	for row := range rows {
		elements := strings.Split(strings.TrimSpace(rows[row]), " ")
		if rows_len == 0 {
			rows_len = len(elements)
		}
		if len(elements) != rows_len {
			return nil, errors.New("Incorrect row length")
		}
		matrix[row] = make([]int, len(elements))
		for element := range elements {
			value, err := strconv.Atoi(elements[element])
			if err != nil {
				return nil, errors.New("Incorrect input")
			}
			matrix[row][element] = value
		}
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	matrix := make(Matrix, len((*m)[0]))
	for row := range matrix {
		matrix[row] = make([]int, len(*m))
		for elem := range matrix[row] {
			matrix[row][elem] = (*m)[elem][row]
		}
	}
	return matrix
}

func (m *Matrix) Rows() [][]int {
	matrix := make(Matrix, len(*m))
	for row := range *m {
		matrix[row] = make([]int, len((*m)[row]))
		for elem := range (*m)[row] {
			matrix[row][elem] = (*m)[row][elem]
		}
	}
	return matrix
}

func (m *Matrix) Set(row, col, val int) bool {
	if len(*m) <= col || len((*m)[0]) <= row || col < 0 || row < 0 {
		return false
	}
	(*m)[row][col] = val
	return true
}
