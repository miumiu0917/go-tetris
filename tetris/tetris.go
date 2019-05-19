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
	Center bool
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

		// 横列が埋まっているか検査
		for i := m.Height-1; i >= 0; i-- {
			tmp := true
			for j := 0; j < m.Width; j++ {
				if m.Field[i][j] == nil {
					tmp = false
				}
			}
			// 埋まっていたら削除
			if tmp {
				for j := 0; j < m.Width; j++ {
					m.Field[i][j] = nil
				}
				for k := i-1; k >= 0; k-- {
					for j := 0; j < m.Width; j++ {
						m.Field[k+1][j] = m.Field[k][j]
					}
				}
				i++
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
	m.Field[0][5] = &Block{ Falling: true, Center: false }
	m.Field[1][5] = &Block{ Falling: true, Center: true }
	m.Field[2][5] = &Block{ Falling: true, Center: false }
	m.Field[3][5] = &Block{ Falling: true, Center: false }
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

	// TODO 移動可能かどうかチェック
	
	// 移動前のブロックを削除
	var tmp []*Block
	var center [2]*int
	for _, p := range target {
		x := p[1]
		y := p[0]
		tmp = append(tmp, m.Field[y][x])
		if m.Field[y][x].Center {
			center[0] = &y
			center[1] = &x
		}
		m.Field[y][x] = nil
	}

	// 移動後の座標にブロックを配置
	for i, p := range target {
		x := p[1]
		y := p[0]
		if direction == 0 {
			m.Field[y][x-1] = tmp[i]
		} else if direction == 1 {
			m.Field[y][x+1] = tmp[i]
		} else if direction == 2 {
			if center[0] != nil {
				centerX := *center[1]
				centerY := *center[0]
				dx := centerX - x
				dy := centerY - y
				m.Field[centerY - dx][centerX + dy] = tmp[i]
			} else {
				m.Field[y][x] = tmp[i]
			}
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
