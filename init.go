package main

import tea "charm.land/bubbletea/v2"

func initialModel() model {
	return model{
		lists: [][]string{
			{"abo44ba", "bimbim", "GoGoG141o"},
			{"bib", "bi414n", "bim", "eg41e"},
			{"bom", "ag4141414141a"},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
