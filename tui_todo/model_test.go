package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func newTestModel(items []TodoItem) Model {
	if items == nil {
		items = []TodoItem{
			{Text: "item a"},
			{Text: "item b"},
			{Text: "item c"},
		}
	}
	return Model{items: items, windowWidth: 60}
}

func TestNavigation(t *testing.T) {
	m := newTestModel(nil)

	tests := []struct {
		name   string
		keys   []tea.KeyType
		runes  []rune
		want   int
	}{
		{name: "start at 0", keys: nil, want: 0},
		{name: "j moves down", keys: []tea.KeyType{tea.KeyRunes}, runes: []rune{'j'}, want: 1},
		{name: "k moves up", keys: []tea.KeyType{tea.KeyRunes, tea.KeyRunes, tea.KeyRunes}, runes: []rune{'j', 'j', 'k'}, want: 1},
		{name: "cannot go below 0", keys: []tea.KeyType{tea.KeyRunes}, runes: []rune{'k'}, want: 0},
		{name: "cannot exceed len-1", keys: []tea.KeyType{tea.KeyRunes, tea.KeyRunes, tea.KeyRunes, tea.KeyRunes}, runes: []rune{'j', 'j', 'j', 'j'}, want: 2},
		{name: "arrow down", keys: []tea.KeyType{tea.KeyDown}, want: 1},
		{name: "arrow up", keys: []tea.KeyType{tea.KeyDown, tea.KeyUp}, want: 0},
		{name: "mixed arrows and jk", keys: []tea.KeyType{tea.KeyRunes, tea.KeyDown, tea.KeyRunes, tea.KeyUp}, runes: []rune{'j', 'k'}, want: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := m
			for i, kt := range tc.keys {
				msg := tea.KeyMsg{Type: kt}
				if kt == tea.KeyRunes && i < len(tc.runes) {
					msg.Runes = []rune{tc.runes[i]}
				}
	
				result, _ := got.Update(msg)
				got = result.(Model)
			}
			if got.listIdx != tc.want {
				t.Errorf("listIdx = %d, want %d", got.listIdx, tc.want)
			}
		})
	}
}

func TestToggleComplete(t *testing.T) {
	tests := []struct {
		name   string
		idx    int
		toggle int
		want   bool
	}{
		{name: "toggle first to true", idx: 0, toggle: 1, want: true},
		{name: "toggle first back to false", idx: 0, toggle: 2, want: false},
		{name: "toggle last", idx: 2, toggle: 1, want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newTestModel(nil)
			m.listIdx = tc.idx
			msg := tea.KeyMsg{Type: tea.KeySpace}
			for range tc.toggle {
				result, _ := m.Update(msg)
				m = result.(Model)
			}
			if m.items[tc.idx].Completed != tc.want {
				t.Errorf("items[%d].Completed = %v, want %v", tc.idx, m.items[tc.idx].Completed, tc.want)
			}
		})
	}
}

func TestDeleteItem(t *testing.T) {
	tests := []struct {
		name    string
		initial []TodoItem
		idx     int
		wantLen int
		wantIdx int
	}{
		{
			name:    "delete first",
			initial: []TodoItem{{Text: "a"}, {Text: "b"}, {Text: "c"}},
			idx:     0,
			wantLen: 2,
			wantIdx: 0,
		},
		{
			name:    "delete middle",
			initial: []TodoItem{{Text: "a"}, {Text: "b"}, {Text: "c"}},
			idx:     1,
			wantLen: 2,
			wantIdx: 1,
		},
		{
			name:    "delete last adjusts index",
			initial: []TodoItem{{Text: "a"}, {Text: "b"}},
			idx:     1,
			wantLen: 1,
			wantIdx: 0,
		},
		{
			name:    "delete only item",
			initial: []TodoItem{{Text: "a"}},
			idx:     0,
			wantLen: 0,
			wantIdx: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newTestModel(tc.initial)
			m.listIdx = tc.idx
			msg := tea.KeyMsg{Type: tea.KeyBackspace}
			result, _ := m.Update(msg)
			m = result.(Model)

			if len(m.items) != tc.wantLen {
				t.Errorf("len(items) = %d, want %d", len(m.items), tc.wantLen)
			}
			if m.listIdx != tc.wantIdx {
				t.Errorf("listIdx = %d, want %d", m.listIdx, tc.wantIdx)
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	m := newTestModel(nil)

	// Enter input mode
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
	result, _ := m.Update(msg)
	m = result.(Model)

	if !m.inputMode {
		t.Fatal("expected inputMode to be true after pressing 'a'")
	}

	// Type some text (space via tea.KeySpace)
	for _, r := range "buy" {
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		result, _ = m.Update(msg)
		m = result.(Model)
	}
	msg = tea.KeyMsg{Type: tea.KeySpace}
	result, _ = m.Update(msg)
	m = result.(Model)
	for _, r := range "milk" {
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		result, _ = m.Update(msg)
		m = result.(Model)
	}
	if m.inputBuf != "buy milk" {
		t.Errorf("inputBuf = %q, want %q", m.inputBuf, "buy milk")
	}

	// Submit
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	result, _ = m.Update(msg)
	m = result.(Model)

	if m.inputMode {
		t.Fatal("expected inputMode to be false after Enter")
	}
	if len(m.items) != 4 {
		t.Fatalf("len(items) = %d, want 4", len(m.items))
	}
	if m.items[3].Text != "buy milk" {
		t.Errorf("items[3].Text = %q, want %q", m.items[3].Text, "buy milk")
	}
}

func TestCancelInput(t *testing.T) {
	m := newTestModel(nil)

	// Enter input mode and type something
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
	result, _ := m.Update(msg)
	m = result.(Model)
	msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h', 'i'}}
	result, _ = m.Update(msg)
	m = result.(Model)

	// Cancel
	msg = tea.KeyMsg{Type: tea.KeyEsc}
	result, _ = m.Update(msg)
	m = result.(Model)

	if m.inputMode {
		t.Fatal("expected inputMode to be false after Esc")
	}
	if m.inputBuf != "" {
		t.Errorf("inputBuf = %q, want empty", m.inputBuf)
	}
}

func TestWindowSize(t *testing.T) {
	m := newTestModel(nil)
	msg := tea.WindowSizeMsg{Width: 100, Height: 50}
	result, _ := m.Update(msg)
	m = result.(Model)
	if m.windowWidth != 100 {
		t.Errorf("windowWidth = %d, want 100", m.windowWidth)
	}
}

func TestQuit(t *testing.T) {
	m := newTestModel(nil)
	msg := tea.KeyMsg{Type: tea.KeyEsc}
	_, cmd := m.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil quit command, got nil")
	}
}

func TestViewRenders(t *testing.T) {
	m := newTestModel([]TodoItem{
		{Text: "alpha", Completed: false},
		{Text: "beta", Completed: true},
	})
	view := m.View()
	if view == "" {
		t.Fatal("View() returned empty string")
	}
	if !contains(view, "alpha") {
		t.Error("expected view to contain 'alpha'")
	}
	if !contains(view, "beta") {
		t.Error("expected view to contain 'beta'")
	}
	// Check completed marker
	if !contains(view, "☑") {
		t.Error("expected view to contain ☑ for completed item")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
