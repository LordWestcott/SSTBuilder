package main

import (
	"errors"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/manifoldco/promptui"
)

func GetProjectName() string {
	prompt := promptui.Prompt{
		Label:    "Project Name",
		Validate: validateProjectName,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed - %v\n \nPlease Retry", err)
		return GetProjectName()
	}

	result = strcase.ToKebab(result)

	return result
}

func validateProjectName(input string) error {
	if len(input) < 1 {
		return errors.New("Project name must be at least 1 character")
	}
	return nil
}

func GetTooling() string {
	prompt := promptui.Select{
		Label: "What tooling would you like to use?",
		Items: []string{"npm", "yarn", "pnpm"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return GetTooling()
	}

	return result
}

func GetFramework() string {
	prompt := promptui.Select{
		Label: "Which Framework would you like to use?",
		Items: []string{"Standalone", "Next", "Remix", "Svelte", "Astro", "Solid"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return GetFramework()
	}

	return result
}

func GetSSTTemplate() string {
	prompt := promptui.Select{
		Label: "Which Template would you like to use?",
		Items: []string{
			"other/go",
			"other/csharp",
			"other/fsharp",
			"other/java",
			"other/python",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return GetFramework()
	}

	return result
}

func GetYesOrNo(label string, reversed bool) bool {

	var prompt promptui.Select
	if reversed {
		prompt = promptui.Select{
			Label: label,
			Items: []string{"Yes", "No"},
		}
	} else {
		prompt = promptui.Select{
			Label: label,
			Items: []string{"No", "Yes"},
		}
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return GetYesOrNo(label, false)
	}

	if result == "Yes" {
		return true
	}

	return false
}

func GetFeatureRequest() string {
	prompt := promptui.Select{
		Label: "What feature would you like to implement?",
		Items: []string{
			"api",
			"auth",
			"config",
			"event",
			"queue",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return GetFramework()
	}

	return result
}
