package tetris

import (
	"fmt"
)

type Map struct {
	Height int
	Width int
	Field [][]bool
}

func (m *Map) Display() {
	s := fmt.Sprintf("\033[%dA", m.Height)
	fmt.Printf(s)
	for _, row := range m.Field {
		for _, p := range row {
			var color string
			if p {
				color = "\x1b[37m\x1b[42m%s\x1b[0m"
			} else {
				color = "\x1b[30m\x1b[47m%s\x1b[0m"
			}
			fmt.Printf(color, "  ")
		}
		fmt.Printf("\n")
	}
}

// func (m *Map) Next() {
// 	for i := range m.Height-1 {
// 		for j := range m.Width-1 {
// 			var color string
// 			if p {
// 				color = "\x1b[37m\x1b[42m%s\x1b[0m"
// 			} else {
// 				color = "\x1b[30m\x1b[47m%s\x1b[0m"
// 			}
// 			fmt.Printf(color, "  ")
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func NewMap(width int, height int) *Map {
	field := make([][]bool, height)
	for i := 0; i < height; i++ {
			field[i] = make([]bool, width)
	}
	m := Map { height, width,  field }
	return &m
}
