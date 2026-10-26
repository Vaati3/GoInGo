package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Vector struct {
	x, y int
}

const Empty rune = '🟨'
const Black rune = '⚫'
const White rune = '⚪'

type Board struct {
	size  Vector
	tiles [19][]rune
}

func (board *Board) Init() {
	board.size = Vector{19, 19}
	for i := range board.size.y {
		var line strings.Builder

		for range board.size.x {
			line.WriteRune(Empty)
		}
		board.tiles[i] = []rune(line.String())
	}
}

func (board *Board) Print() {
	var letters strings.Builder
	letters.WriteString("   ")
	for j := range board.size.x {
		letters.WriteString(string(j+'A') + "  ")
	}
	fmt.Println(letters.String())
	for i := range board.size.y {
		var line strings.Builder
		if i < 9 {
			line.WriteString(" ")
		}
		line.WriteString(strconv.Itoa(i + 1))
		line.WriteString(string(board.tiles[i]))
		fmt.Println(line.String())
	}
}

func (board *Board) Move(x, y int, isWhite bool) bool {
	if board.tiles[y][x] != Empty {
		return false
	}

	if isWhite {
		board.tiles[y][x] = White
	} else {
		board.tiles[y][x] = Black
	}
	board.CheckCaptures(x, y, board.tiles[y][x])
	board.Print()
	return true
}

func (board *Board) GetNeighbours(x, y int) []Vector {
	var neighbours []Vector
	pos := [4]Vector{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := range pos {
		x, y := x+pos[i].x, y+pos[i].y
		if x >= 0 && x < board.size.x && y >= 0 && y < board.size.y && board.tiles[y][x] != Empty {
			neighbours = append(neighbours, Vector{x, y})
		}
	}

	return neighbours
}

func VectorInArray(array []Vector, pos Vector) bool {
	for _, p := range array {
		if p.x == pos.x && p.y == pos.y {
			return true
		}
	}
	return false
}

func (board *Board) GetGroup(x, y int) ([]Vector, []Vector) {
	if board.tiles[y][x] == Empty {
		return nil, nil
	}
	colour := board.tiles[y][x]
	groups := make([]Vector, 0)
	liberties := make([]Vector, 0)

	for _, pos := range board.GetNeighbours(x, y) {
		if board.tiles[pos.y][pos.x] == Empty {
			liberties = append(liberties, pos)
		} else if board.tiles[pos.y][pos.x] == colour && !VectorInArray(groups, pos) {
			groups = append(groups, pos)
			g, l := board.GetGroup(pos.x, pos.y)
			groups = append(groups, g...)
			liberties = append(liberties, l...)
		}
	}

	return groups, liberties
}

func (board *Board) CheckCaptures(moveX, moveY int, colour rune) int {
	var otherColour rune
	if colour == White {
		otherColour = Black
	} else {
		otherColour = White
	}
	count := 0

	for _, pos := range board.GetNeighbours(moveX, moveY) {
		if board.tiles[pos.y][pos.x] == otherColour {
			groups, liberties := board.GetGroup(pos.x, pos.y)

			if len(liberties) == 0 {
				for _, group := range groups {
					if board.tiles[group.y][group.x] != Empty {
						board.tiles[group.y][group.x] = Empty
						count++
					}
				}
			}
		}
	}

	//check suicide
	if count == 0 {
		groups, liberties := board.GetGroup(moveX, moveY)
		if len(liberties) == 0 {
			for _, group := range groups {
				if board.tiles[group.y][group.x] != Empty {
					board.tiles[group.y][group.x] = Empty
					count++
				}
			}
		}
	}
	return count
}

func (board *Board) IsGameOver() bool {
	//TO DO

	return false
}
