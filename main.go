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
	t.Display()
	time.Sleep(3 * time.Second)
	t.Field[5][5] = true
	time.Sleep(3 * time.Second)
	t.Display()
	t.Field[5][6] = true
	time.Sleep(3 * time.Second)
	t.Display()
	t.Field[5][7] = true
	time.Sleep(3 * time.Second)
	t.Display()
	t.Field[4][6] = true
}