package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TodoItem struct {
	Text      string
	Completed bool
}

type Model struct {
	items      []TodoItem
	listIdx    int
	inputBuf   string
	inputMode  bool
	windowWidth int
}

func NewModel() Model {
	return Model{
		items: []TodoItem{
			{Text: "Add new item with a"},
			{Text: "Delete an item with backspace"},
			{Text: "Check an item with space"},
			{Text: "Quit with Esc"},
		},
		windowWidth: 60,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {
			return m.updateInput(msg), nil
		}
		return m.updateNormal(msg)
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		return m, nil
	}
	return m, nil
}

func (m Model) updateInput(msg tea.KeyMsg) Model {
	switch msg.Type {
	case tea.KeyRunes:
		m.inputBuf += string(msg.Runes)
	case tea.KeySpace:
		m.inputBuf += " "
	case tea.KeyBackspace:
		if len(m.inputBuf) > 0 {
			m.inputBuf = m.inputBuf[:len(m.inputBuf)-1]
		}
	case tea.KeyEnter:
		if m.inputBuf != "" {
			m.items = append(m.items, TodoItem{Text: m.inputBuf})
			m.listIdx = len(m.items) - 1
		}
		m.inputMode = false
		m.inputBuf = ""
	case tea.KeyEsc:
		m.inputMode = false
		m.inputBuf = ""
	}
	return m
}

func (m Model) updateNormal(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyRunes:
		for _, r := range msg.Runes {
			switch r {
			case 'j':
				if m.listIdx < len(m.items)-1 {
					m.listIdx++
				}
			case 'k':
				if m.listIdx > 0 {
					m.listIdx--
				}
			case 'a', 'n':
				m.inputMode = true
				m.inputBuf = ""
			}
		}
	case tea.KeyDown:
		if m.listIdx < len(m.items)-1 {
			m.listIdx++
		}
	case tea.KeyUp:
		if m.listIdx > 0 {
			m.listIdx--
		}
	case tea.KeySpace:
		if m.listIdx >= 0 && m.listIdx < len(m.items) {
			m.items[m.listIdx].Completed = !m.items[m.listIdx].Completed
		}
	case tea.KeyBackspace:
		if m.listIdx >= 0 && m.listIdx < len(m.items) {
			m.items = append(m.items[:m.listIdx], m.items[m.listIdx+1:]...)
			if m.listIdx >= len(m.items) {
				m.listIdx = max(0, len(m.items)-1)
			}
		}
	case tea.KeyEsc:
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() string {
	win := m.windowWidth
	if win < 20 {
		win = 60
	}

	// Border + padding add 4 to total width, so content width is win-4.
	inner := max(0, win-4)

	headerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("220")).
		Padding(0, 1).
		Width(inner)

	header := headerStyle.Render(" Todo list ")

	var items string
	for i, item := range m.items {
		check := "☐"
		if item.Completed {
			check = "☑"
		}

		itemStr := fmt.Sprintf("%s %s", check, item.Text)

		var style lipgloss.Style
		if i == m.listIdx {
			style = lipgloss.NewStyle().Background(lipgloss.Color("70")).Bold(true)
		} else {
			style = lipgloss.NewStyle()
		}
		if item.Completed {
			style = style.Strikethrough(true)
		}

		items += style.Render(itemStr) + "\n"
	}

	content := header + "\n\n" + items

	if m.inputMode {
		inputStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("51")).
			Padding(0, 1).
			Width(inner)
		content += "\n\n" + inputStyle.Render(" " + m.inputBuf + " ")
	}

	return lipgloss.NewStyle().Width(win).Margin(1, 0).Render(content)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
