package main

import (
	tea "charm.land/bubbletea/v2"
	lipg "charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	var viewList []string

	for _, list := range m.lists {
		s := lipg.JoinVertical(lipg.Center, list...)
		viewList = append(viewList, titleStyle.Render(s))
	}

	s := lipg.JoinHorizontal(lipg.Center, viewList...)

	return tea.NewView(s)
}
