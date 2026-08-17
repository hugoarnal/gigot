package commands

import (
	"fmt"
	"gigot/internal/config"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/spf13/cobra"
)

type model struct {
	choices []config.GitConfigConfig
	cursor  int

	// The selected config name
	Config *config.GitConfigConfig
}

func (m model) View() tea.View {
	s := "Select which config to switch to\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		enabled := ""
		if choice.Enabled {
			enabled = " (enabled)"
		}

		s += fmt.Sprintf("%s %s%s\n", cursor, choice.Name, enabled)
	}

	s += "\nPress q to quit\n"
	return tea.NewView(s)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			m.Config = &m.choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}

func SwitchList(parsedConfig []config.GitConfigConfig, filename string) error {
	p := tea.NewProgram(model{
		choices: parsedConfig,
		Config:  nil,
	})
	teaModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	m := teaModel.(model)

	if m.Config == nil {
		return nil
	}

	if err := config.SwitchGitConfigConfig(&parsedConfig, m.Config.Name); err != nil {
		return err
	}

	if err := config.WriteGitConfigFile(filename, parsedConfig); err != nil {
		return err
	}

	printSwitchConfig(m.Config)
	return nil
}

func printSwitchConfig(config *config.GitConfigConfig) {
	particle := "on"
	if !config.Enabled {
		particle = "off"
	}

	fmt.Printf("Toggled %s \"%s\" configuration\n", particle, config.Name)
}

func SwitchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch",
		Short: "Switch to another gitconfig",
		RunE:  Switch,
	}
	cmd.Flags().Bool("disable", false, "Disable the currently enabled gitconfig")
	return cmd
}

func Switch(cmd *cobra.Command, args []string) error {
	disable, _ := cmd.Flags().GetBool("disable")

	config.GigotShellWarning()

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		return err
	}

	if disable {
		for i := range parsedConfig {
			parsedConfig[i].Enabled = false
		}

		if err := config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig); err != nil {
			return err
		}

		fmt.Println("Disabled all enabled configurations")
		return nil
	}

	switch len(args) {
	case 0:
		err := SwitchList(parsedConfig, filename)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case 1:
		selectedConfig := args[0]

		if err := config.SwitchGitConfigConfig(&parsedConfig, selectedConfig); err != nil {
			return err
		}

		if err := config.WriteGitConfigFile(filename, parsedConfig); err != nil {
			return err
		}

		configIndex, err := config.FindGitConfigByName(parsedConfig, selectedConfig)

		if err != nil {
			return err
		}

		printSwitchConfig(&parsedConfig[configIndex])
	default:
		return fmt.Errorf("incorrect amount of arguments")
	}

	return nil
}
