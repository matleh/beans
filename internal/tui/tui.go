package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"hmans.dev/beans/internal/bean"
	"hmans.dev/beans/internal/config"
)

// viewState represents which view is currently active
type viewState int

const (
	viewList viewState = iota
	viewDetail
)

// App is the main TUI application model
type App struct {
	state   viewState
	list    listModel
	detail  detailModel
	history []detailModel // stack of previous detail views for back navigation
	store   *bean.Store
	config  *config.Config
	width   int
	height  int
}

// New creates a new TUI application
func New(store *bean.Store, cfg *config.Config) *App {
	return &App{
		state:  viewList,
		store:  store,
		config: cfg,
		list:   newListModel(store, cfg),
	}
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	return a.list.Init()
}

// Update handles messages
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		case "q":
			if a.state == viewDetail {
				return a, tea.Quit
			}
			// For list, only quit if not filtering
			if a.state == viewList && a.list.list.FilterState() != 1 {
				return a, tea.Quit
			}
		}

	case selectBeanMsg:
		// Push current detail view to history if we're already viewing a bean
		if a.state == viewDetail {
			a.history = append(a.history, a.detail)
		}
		a.state = viewDetail
		a.detail = newDetailModel(msg.bean, a.store, a.config, a.width, a.height)
		return a, a.detail.Init()

	case backToListMsg:
		// Pop from history if available, otherwise go to list
		if len(a.history) > 0 {
			a.detail = a.history[len(a.history)-1]
			a.history = a.history[:len(a.history)-1]
			// Stay in viewDetail state
		} else {
			a.state = viewList
		}
		return a, nil
	}

	// Forward all messages to the current view
	switch a.state {
	case viewList:
		a.list, cmd = a.list.Update(msg)
	case viewDetail:
		a.detail, cmd = a.detail.Update(msg)
	}

	return a, cmd
}

// View renders the current view
func (a *App) View() string {
	switch a.state {
	case viewList:
		return a.list.View()
	case viewDetail:
		return a.detail.View()
	}
	return ""
}

// Run starts the TUI application
func Run(store *bean.Store, cfg *config.Config) error {
	p := tea.NewProgram(New(store, cfg), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
