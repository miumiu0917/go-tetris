package tetris

import (
	"fmt"
)

type Map struct {
	Height int
	Width int
	Field [][]*Block
}

type Block struct {
	Falling bool
}

func (m *Map) Display() {
	s := fmt.Sprintf("\033[%dA", m.Height-4)
	fmt.Printf(s)
	for index, row := range m.Field {
		if index < 4 {
			continue
		}
		for _, p := range row {
			var color string
			if p != nil {
				color = "\x1b[37m\x1b[42m%s\x1b[0m"
			} else {
				color = "\x1b[30m\x1b[47m%s\x1b[0m"
			}
			fmt.Printf(color, "  ")
		}
		fmt.Printf("\n")
	}
}

func (m *Map) Next() {
	for i := m.Height-2; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			if m.Field[i+1][j] == nil {
				m.Field[i+1][j] = m.Field[i][j]
				m.Field[i][j] = nil
			}
		}
	}
	m.Field[0] = make([]*Block, m.Width)
}

func NewMap(width int, height int) *Map {
	height += 4
	field := make([][]*Block, height)
	for i := 0; i < height; i++ {
			field[i] = make([]*Block, width)
	}
	m := Map { height, width,  field }
	return &m
}
