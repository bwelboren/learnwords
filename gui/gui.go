package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	wrts "bwelboren.github.io/cmd"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))

const (
	email    = ""
	password = ""
)

type model struct {
	choices  []string
	cursor   int
	Chosen   bool
	Quitting bool
	selected map[int]struct{}
	wordlist []wrts.WordList
	viewport viewport.Model
	ready    bool
}

var SwapLang bool
var Update bool

func initialModel() model {

	// wrts.SetAuthToken(email, password)
	// wrts.GetOfficialLists("Russisch")
	// wl := wrts.GetAllWordsFromWordLists()

	lists := []string{}

	wl := []wrts.WordList{}

	wordList := []string{}

	// Fill up []wrts.WordLists.[]Words

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, strings.Split(scanner.Text(), "=")...)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Test Code
	for i := 0; i < 5; i++ {
		wl = append(wl, wrts.WordList{
			ID:    fmt.Sprintf("%d", i),
			Name:  fmt.Sprintf("Russian Words %d", i),
			Lang:  []string{"Nederlands", "Russisch"},
			Words: wordList,
		})

	}

	for _, wordlists := range wl {
		lists = append(lists, wordlists.Name)
	}

	return model{
		choices:  lists,
		selected: make(map[int]struct{}),
		wordlist: wl,
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
		if m.Chosen {

			switch k := msg.String(); k {
			case "s":
				if !SwapLang {
					SwapLang = true

				} else if SwapLang {
					SwapLang = false

				}
			case "b":
				SwapLang = false
				m.Chosen = false
			}
			m.viewport.SetContent(PrintWords(m))

		} else {
			return updateLists(msg, m)
		}

	case tea.WindowSizeMsg:
		m.viewport = viewport.New(msg.Width, msg.Height-4)
		m.viewport.YPosition = 0
		m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
		m.viewport.SetContent(PrintWords(m))
	}

	return m, nil
}

func updateLists(msg tea.Msg, m model) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			// Tell program we selected a wordlist
			m.Chosen = true

		}

	}

	return m, nil
}

func viewLists(m model) string {
	s := "Selecteer een woordenlijst.\n\n"

	for i, choice := range m.choices {

		if m.cursor == i {
			choice = selectedItemStyle.Render(choice)

		}
		s += choice + "\n"

	}

	return s
}

func (m model) View() string {

	s := ""
	hotkeys := ""

	if !m.Chosen {
		s = viewLists(m)
	} else {
		s = LanguageHeader(m) + "\n"
		s += m.viewport.View()

	}

	s += fmt.Sprintf("\n\n%sctrl+c to quit\n", hotkeys)

	return s
}

func PrintWords(m model) string {

	buf := new(bytes.Buffer)
	w := new(tabwriter.Writer)

	w.Init(buf, 0, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)

	for i := 0; i < len(m.wordlist[m.cursor].Words); i = i + 2 {
		if !SwapLang {
			fmt.Fprintln(w, m.wordlist[m.cursor].Words[i]+"\t"+m.wordlist[m.cursor].Words[i+1])

		} else {
			fmt.Fprintln(w, m.wordlist[m.cursor].Words[i+1]+"\t"+m.wordlist[m.cursor].Words[i])
		}
	}

	fmt.Fprintln(w)
	data := buf.Bytes()
	s := string(data)
	w.Flush()

	return s

}

func LanguageHeader(m model) string {
	var s string
	if !SwapLang {
		s = m.wordlist[m.cursor].Lang[1] + " " + m.wordlist[m.cursor].Lang[0]
	} else {
		s = m.wordlist[m.cursor].Lang[0] + " " + m.wordlist[m.cursor].Lang[1]
	}
	return s
}

func main() {

	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

}
