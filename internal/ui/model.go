package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	gloss "github.com/charmbracelet/lipgloss"
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

	width  int
	height int
}

func New() Model { return Model{} }

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

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

var (
	clockStyle = gloss.NewStyle().
			Bold(true).
			Foreground(gloss.AdaptiveColor{Light: "#1971c2", Dark: "#4dabf7"})

	cardStyle = gloss.NewStyle().
			Border(gloss.RoundedBorder()).
			BorderForeground(gloss.AdaptiveColor{Light: "#dee2e6", Dark: "#343a40"}).
			Padding(1, 4).
			Align(gloss.Center)

	dimStyle = gloss.NewStyle().
			Foreground(gloss.AdaptiveColor{Light: "#adb5bd", Dark: "#6c757d"})

	dotReady = gloss.NewStyle().
			Foreground(gloss.AdaptiveColor{Light: "#adb5bd", Dark: "#6c757d"})

	dotRunning = gloss.NewStyle().
			Foreground(gloss.AdaptiveColor{Light: "#2f9e44", Dark: "#51cf66"})

	dotStopped = gloss.NewStyle().
			Foreground(gloss.AdaptiveColor{Light: "#e8590c", Dark: "#ffd43b"})
)

func (m Model) View() string {
	status, dot := "ready", dotReady
	switch {
	case m.running:
		status, dot = "running", dotRunning
	case m.elapsed > 0:
		status, dot = "stopped", dotStopped
	}

	card := cardStyle.Render(gloss.JoinVertical(
		gloss.Center,
		clockStyle.Render(format(m.elapsed)),
		"",
		dot.Render("●")+" "+dimStyle.Render(status),
	))

	body := lipgloss.JoinVertical(gloss.Center,
		card,
		"",
		dimStyle.Render("space start/stop   r reset   q quit"),
	)

	if m.width == 0 || m.height == 0 {
		return body
	}

	return gloss.Place(m.width, m.height, gloss.Center, gloss.Center, body)
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
