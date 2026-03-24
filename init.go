package main

import tea "charm.land/bubbletea/v2"

func initialModel() model {
	return model{
		list:     []string{"aboba", "bimbim", "GoGoGo"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
