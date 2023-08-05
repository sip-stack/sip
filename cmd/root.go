package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/paulmanoni/sip/pkg/registry"
	"github.com/paulmanoni/sip/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	useCustomTemplate                  bool // define custom templates
	createAnswers, customCreateAnswers registry.CreateAnswers
	surveyIconsConfig                  = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "[?]"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "yellow"
		icons.Error.Text = "Note ->"
	}
)

var rootCmd = &cobra.Command{
	Use:   "sip",
	Short: "sip is a CLI for managing SIP accounts",
	Long:  `sip is a CLI for managing SIP accounts. It is designed to be used with`,
}

func init() {
	rootCmd.SetOut(utils.Stdout)
	rootCmd.SetErr(utils.Stderr)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
