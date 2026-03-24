package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	list     []string
	cursor   int
	selected map[int]struct{}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERRRRRRRRRRR: %v", err)
		os.Exit(1)
	}
}
