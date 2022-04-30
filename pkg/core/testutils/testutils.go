package testutils

import (
	tea "github.com/charmbracelet/bubbletea"
)

func MsgKey(runeKey rune) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{runeKey}, Alt: false}
}

func ModelUpdate[M tea.Model](model *M, msg tea.Msg) {
	var cmd tea.Cmd
	var teaModel tea.Model
	teaModel = *model

	for teaModel, cmd = teaModel.Update(msg); cmd != nil; teaModel, cmd = teaModel.Update(msg) {
		msg = cmd()

		switch cmds := msg.(type) {
		case []tea.Cmd:
			for _, cmd = range cmds {
				ModelUpdate(model, cmd())
			}
		}

	}
	*model = teaModel.(M)
}
