package testutils

import (
	"reflect"
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
)

func MsgKey(runeKey rune) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{runeKey}, Alt: false}
}

func MsgKeyByType(keyType tea.KeyType) tea.KeyMsg {
	return tea.KeyMsg{Type: keyType, Runes: []rune{}, Alt: false}
}

var ignoredMessages = []string{
	"textinput.initialBlinkMsg",
	"textinput.blinkMsg",
}

func ModelUpdate[M tea.Model](model *M, msg tea.Msg) {
	var cmd tea.Cmd
	var teaModel tea.Model
	teaModel = *model

	for teaModel, cmd = (*model).Update(msg); cmd != nil; teaModel, cmd = teaModel.Update(msg) {
		msg = cmd()

		if shouldBeIgnored(msg) {
			break
		}

		switch cmds := msg.(type) {
		case []tea.Cmd:
			for _, cmd = range cmds {
				ModelUpdate(model, cmd())
			}
		}

	}
	*model = teaModel.(M)
}

func shouldBeIgnored(msg tea.Msg) bool {
	msgTypeStr := reflect.TypeOf(msg).String()

	for _, e := range ignoredMessages {
		if e == msgTypeStr {
			return true
		}
	}
	return false
}

var ignoreAnsiEscapes = regexp.MustCompile(`\x1b\[.*?m`)

func ToPlainText(data string) string {
	x := ignoreAnsiEscapes.ReplaceAllString(data, "")
	return x
}
