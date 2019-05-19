package main

import (
	"github.com/go-tetris/tetris"
	"time"
	"fmt"
	"strings"
	term "github.com/nsf/termbox-go"
	"os"
)

func main() {
	width, height := 20, 20
	t := tetris.NewMap(width, height)
	fmt.Printf(strings.Repeat("\n", height))

	q := make(chan int, 2)
	go interact(q)

	for {
		t.Display()
		time.Sleep(500 * time.Millisecond)
		if len(q) > 0 {
			event := <- q
			if event == 0 || event == 1 || event == 2 {
				t.Move(event)
			}
		}
		t.Next()
		if t.IsAllFreeze() {
			t.NextBlock()
		}
	}
}

func reset() {
	term.Sync()
}

func interact(q chan int) {
	err := term.Init()
	if err != nil {
			panic(err)
	}
	defer term.Close()
	
	for {
		switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
					switch ev.Key {
						case term.KeyEsc:
							term.Close()
							os.Exit(0)
						case term.KeyArrowLeft:
							q <- 0
						case term.KeyArrowRight:
							q <- 1
						case term.KeyArrowUp:
							q <- 2
						default:
								// we only want to read a single character or one key pressed event
					}
			case term.EventError:
					panic(ev.Err)
		}
	}
}