package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Empty rune = '🟨'
const Black rune = '⚫'
const White rune = '⚪'

type Board struct {
	tiles [20][]rune
}

func (board *Board) Init() {
	for i := range 20 {
		var line strings.Builder
		if i > 0 {
			if i < 10 {
				line.WriteString(" ")
			}
			line.WriteString(strconv.Itoa(i))
		} else {
			line.WriteString("   ")
		}

		for j := range 19 {
			if i < 1 {
				line.WriteString(string(j+'A') + "  ")
			} else {
				line.WriteString(string(Empty))
			}
		}
		board.tiles[i] = []rune(line.String())
	}
}

func (board *Board) Print() {
	for i := range board.tiles {
		fmt.Println(string(board.tiles[i]))
	}
}

func (board *Board) Move(x, y int, isWhite bool) bool {
	if board.tiles[y+1][x+2] != Empty {
		return false
	}

	if isWhite {
		board.tiles[y+1][x+2] = White
	} else {
		board.tiles[y+1][x+2] = Black
	}
	board.Print()
	return true
}

func (board *Board) IsGameOver() bool {
	//TO DO

	return false
}