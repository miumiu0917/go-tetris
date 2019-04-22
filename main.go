package main

import (
	"github.com/go-tetris/tetris"
	"time"
	"fmt"
	"strings"
)

func main() {
	width, height := 20, 20
	t := tetris.NewMap(width, height)
	fmt.Printf(strings.Repeat("\n", height))
	t.Field[5][5] = &tetris.Block{ true }
	t.Field[6][5] = &tetris.Block{ true }
	t.Field[7][5] = &tetris.Block{ true }
	t.Field[6][6] = &tetris.Block{ true }
	for {
		t.Display()
		time.Sleep(500 * time.Millisecond)
		t.Next()
		if t.IsAllFreeze() {
			t.NextBlock()
		}
	}
}