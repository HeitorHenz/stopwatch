package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

const TickRate = time.Second / 60

func tickCmd() tea.Cmd {
	return tea.Tick(TickRate, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type Model struct {
	running     bool
	startedAt   time.Time
	elapsed     time.Duration
	accumulated time.Duration
}

func New() Model { return Model{} }

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if !m.running {
			return m, nil
		}
		m.elapsed = m.accumulated + time.Since(m.startedAt)
		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ":
			return m.toggle()
		case "r":
			return m.reset()
		}
	}
	return m, nil
}

func (m Model) View() string {
	status := "ready"
	switch {
	case m.running:
		status = "running"
	case m.elapsed > 0:
		status = "stopped"
	}
	return "\n  " + m.elapsed.String() + "\n\n  " + status + "\n\n  space start/stop   r reset   q quit\n"
}

func (m Model) toggle() (tea.Model, tea.Cmd) {
	if m.running {
		m.running = false
		m.accumulated += time.Since(m.startedAt)
		m.elapsed = m.accumulated
		return m, nil
	}
	m.startedAt = time.Now()
	m.running = true
	return m, tickCmd()
}

func (m Model) reset() (tea.Model, tea.Cmd) {
	m.accumulated = 0
	m.elapsed = 0
	m.startedAt = time.Time{}
	m.running = false
	return m, nil
}
