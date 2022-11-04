package main

import (
	"fmt"
	"strings"
)

type State [][]byte

func xor(a, b byte) byte {
	return a ^ b
}

func xorTable(state, other State) State {
	newState := NewZeroedState(len(state), len(state[0]))
	for i, row := range state {
		for j, val := range row {
			newState[i][j] = xor(val, other[i][j])
		}
	}

	return newState
}

func NewZeroedState(rows, cols int) State {
	newState := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		newState[i] = make([]byte, cols)
	}
	return newState
}

func (s *State) appendColumn(column State) State {
	if len((*s)[0]) == 0 {
		return column
	}
	numRows := len(*s)
	numCols := len((*s)[0]) + 1

	newState := NewZeroedState(numRows, numCols)

	for i, row := range *s {
		copy(newState[i], row)
		newState[i][numCols-1] = column[i][0]
	}

	return newState
}

func (s *State) getTableColumn(colNum int) State {
	column := NewZeroedState(len(*s), 1)

	for i, row := range *s {
		column[i][0] = row[colNum]
	}

	return column
}

func (s *State) tableToString() string {
	var sb strings.Builder

	for _, row := range *s {
		sb.WriteString(" ")
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%02x ", val))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
