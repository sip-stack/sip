package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/paulmanoni/sip/pkg/registry"
	"github.com/paulmanoni/sip/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend and frontend templates",
	)
}

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI",
	Long: `Create a new project via interactive UI
		- Project name
		- Project description`,
	RunE: runCreateCmd,
}

func runCreateCmd(cmd *cobra.Command, args []string) error {
	utils.ShowMessage(
		"",
		fmt.Sprintf(
			"Create a new project via Create Go App CLI v%v...",
			"1.0.0",
		),
		true, true,
	)
	backend := ""
	frontend := "none"
	folder := "backend"

	if useCustomTemplate {
		if err := survey.Ask(registry.CreateQuestions, &customCreateAnswers, survey.WithIcons(surveyIconsConfig)); err != nil {
			return utils.ShowError(err.Error())
		}
		backend = createAnswers.Backend
		frontend = createAnswers.Frontend
		folder = createAnswers.Project
	} else {
		if err := survey.Ask(registry.CreateQuestions, &createAnswers, survey.WithIcons(surveyIconsConfig)); err != nil {
			return utils.ShowError(err.Error())
		}
		backend = createAnswers.Backend
		frontend = createAnswers.Frontend
		folder = createAnswers.Project
	}
	if err := utils.GitClone(folder, fmt.Sprintf("https://github.com/paulmanoni/%v-template", backend)); err != nil {
		return utils.ShowError(err.Error())
	}
	if frontend != "none" {
		switch frontend {
		case "nuxt":
			if err := utils.ExecCommand(
				"npx", []string{"nuxi", "init", "frontend"}, true,
			); err != nil {
				return err
			}
		default:
			// Create a default frontend template from Vite (Pure JS/TS, React, Preact, Vue, Svelte, Lit).
			if err := utils.ExecCommand(
				"npm", []string{"init", "vite@latest", "frontend", "--", "--template", frontend}, true,
			); err != nil {
				return err
			}
		}
	}
	utils.ShowMessage(
		"",
		"Have a happy new project! :)",
		false, true,
	)
	return nil
}
