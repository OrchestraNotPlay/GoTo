package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	lists  [][]string
	cursor int
	width  int
	height int
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERRRRRRRRRRR: %v", err)
		os.Exit(1)
	}
}
