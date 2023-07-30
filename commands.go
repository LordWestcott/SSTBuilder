package main

import (
	"fmt"

	"github.com/fatih/color"
)

func next_create_new_app(name string, tooling string) error {
	color.Yellow("Creating new Next App...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "create-next-app@latest", name, "--ts", "--use-npm")
	case "yarn":
		return RunCommand("yarn", "create", "next-app", name, "--ts", "--use-yarn")
	case "pnpm":
		return RunCommand("pnpm", "create", "next-app", name, "--ts", "--use-pnpm")
	}
	return nil
}

func remix_create_new_app(name string, tooling string) error {
	color.Yellow("Creating new Remix App...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "create-remix@latest", name)
	case "yarn":
		return RunCommand("yarn", "create", "remix", name)
	case "pnpm":
		return RunCommand("pnpm", "create", "remix", name)
	}
	return nil
}

func svelte_create_new_app(name string, tooling string) error {
	color.Yellow("Creating new Svelte App...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "create-svelte@latest", name)
	case "yarn":
		return RunCommand("yarn", "create", "svelte", name)
	case "pnpm":
		return RunCommand("pnpm", "create", "svelte", name)
	}
	return nil
}

func astro_create_new_app(name string, tooling string) error {
	color.Yellow("Creating new Astro App...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "create-astro@latest", name)
	case "yarn":
		return RunCommand("yarn", "create", "astro", name)
	case "pnpm":
		return RunCommand("pnpm", "create", "astro", name)
	}
	return nil
}

func solid_create_new_app(name string, tooling string) error {
	color.Yellow("Creating new Solid App...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "create-solid@latest", name)
	case "yarn":
		return RunCommand("yarn", "create", "solid", name)
	case "pnpm":
		return RunCommand("pnpm", "create", "solid", name)
	}
	return nil
}

func install_sst(tooling string, name string) error {
	color.Yellow("Installing SST...")

	useTemplate := GetYesOrNo("Would you like to use a SST 'other' template?", false)
	template := ""
	if useTemplate {
		template = GetSSTTemplate()
	}

	if name != "" {
		switch tooling {
		case "npm":
			if template != "" {
				return RunCommand("npx", "create-sst@latest", name, fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("npx", "create-sst@latest", name)
		case "yarn":
			if template != "" {
				return RunCommand("yarn", "create", "sst", name, fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("yarn", "create", "sst", name)
		case "pnpm":
			if template != "" {
				return RunCommand("pnpm", "create", "sst", name, fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("pnpm", "create", "sst", name)
		}
	} else {

		switch tooling {
		case "npm":
			if template != "" {
				return RunCommand("npx", "create-sst@latest", fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("npx", "create-sst@latest")
		case "yarn":
			if template != "" {
				return RunCommand("yarn", "create", "sst", fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("yarn", "create", "sst")
		case "pnpm":
			if template != "" {
				return RunCommand("pnpm", "create", "sst", fmt.Sprintf("--template=%s", template))
			}
			return RunCommand("pnpm", "create", "sst")
		}
	}
	return nil
}

func start_sst_locally(tooling string) error {
	color.Yellow("Starting SST locally...")
	switch tooling {
	case "npm":
		return RunCommand("npx", "sst", "dev")
	case "yarn":
		return RunCommand("yarn", "sst", "dev")
	case "pnpm":
		return RunCommand("pnpm", "sst", "dev")
	}
	return nil
}
