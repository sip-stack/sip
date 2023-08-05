package registry

import "github.com/AlecAivazis/survey/v2"

var (
	CreateQuestions = []*survey.Question{
		{
			Name: "project",
			Prompt: &survey.Input{
				Message: "Project Name:",
				Default: "my-project",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"net/http",
					"gin-gonic",
					"fiber",
				},
				Default:  "gin-gonic",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend framework/library:",
				Help:    "Option with a `*-ts` tail will create a TypeScript template.",
				Options: []string{
					"none",
					"nuxt",
					"vue",
					"vue-ts",
				},
				Default:  "none",
				PageSize: 4,
			},
		},
	}
)
