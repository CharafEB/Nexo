package Test

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	famwork  []string
	cursor   int
	selacted map[int]struct{}
}

func initModel() model {
	return model{
		famwork:  []string{"Node", "Golang", "Djolang"},
		selacted: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
		switch msg.String() {
			case "q" , "ctrl+c":
				return m, tea.Quit
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.famwork)-1 {
					m.cursor++
				}

			case "enter":
				if _, ok := m.selacted[m.cursor]; ok {
					delete(m.selacted, m.cursor) 
				} else {
					m.selacted[m.cursor] = struct{}{} 
				}

	}
}
return m, nil

}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"
	for i , selact := range m.famwork {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "

		if _, ok := m.selacted[i] ; ok {
			checked = " X "
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, selact)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main(){
	p := tea.NewProgram(initModel())
	if _,err := p.Run() ; err != nil {
		log.Fatal(err)
	}
}
