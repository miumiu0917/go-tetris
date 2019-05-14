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
	freeze := false
	// 落下中のものをチェック
	for i := m.Height-1; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			if m.Field[i][j] != nil && m.Field[i][j].Falling  {
				if i == m.Height-1 || (m.Field[i+1][j] != nil && !m.Field[i+1][j].Falling) {
					freeze = true
				}
			}
		}
	}

	// ブロックが下に着いたらフリーズ
	if freeze {
		for i := m.Height-1; i >= 0; i-- {
			for j := 0; j < m.Width; j++ {
				if m.Field[i][j] != nil && m.Field[i][j].Falling  {
					m.Field[i][j].Falling = false
				}
			}
		}
	}

	// フリーズされていないブロックを1マス下げる
	for i := m.Height-2; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			if m.Field[i+1][j] == nil && m.Field[i][j] != nil && m.Field[i][j].Falling {
				m.Field[i+1][j] = m.Field[i][j]
				m.Field[i][j] = nil
			}
		}
	}
	m.Field[0] = make([]*Block, m.Width)
}

func (m *Map) IsAllFreeze() bool {
	for i := m.Height-1; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			if m.Field[i][j] != nil && m.Field[i][j].Falling  {
					return false
			}
		}
	}
	return true
}

func (m *Map) NextBlock() {
	m.Field[0][5] = &Block{ true }
	m.Field[1][5] = &Block{ true }
	m.Field[2][5] = &Block{ true }
	m.Field[1][6] = &Block{ true }
}

func (m *Map) Move(direction int) {
	var target [][2]int
	for i := m.Height-1; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			if m.Field[i][j] != nil && m.Field[i][j].Falling  {
				target = append(target, [2]int{i, j})
			}
		}
	}
	for _, p := range target {
		x := p[1]
		y := p[0]
		m.Field[y][x] = nil
	}
	for _, p := range target {
		x := p[1]
		y := p[0]
		if direction == 0 {
			m.Field[y][x-1] = &Block{ true }
		} else if direction == 1 {
			m.Field[y][x+1] = &Block{ true }
		}
	}
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
