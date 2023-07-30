package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
)

type Project struct {
	Name      string
	ParentDir string
	Framework string
	Tooling   string
}

func (p *Project) Init(wg *sync.WaitGroup) {
	defer wg.Done()

	// isNew := GetYesOrNo("Is this a new project?", true)
	isNew := true
	if isNew {
		p.Name = GetProjectName()
		dir, err := getLocalDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p.ParentDir = dir

		p.Tooling = GetTooling()
		p.Framework = GetFramework()

		p.Setup()
	}

	if !isNew {
		var err error
		p.Name, p.ParentDir, err = getExistingProjectNameAndParentDirectory()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p.Tooling = GetTooling()
		p.Framework = GetFramework()

		fmt.Println(p.Name)
		fmt.Println(p.ParentDir)
		fmt.Println(p.Tooling)
		fmt.Println(p.Framework)
	}

	return

	for {
		quit := !GetYesOrNo("Would you like to add a feature?", true)
		if quit {
			break
		}

		feature := GetFeatureRequest()
		switch feature {
		}
		fmt.Println(feature)
	}

}

func (p *Project) Setup() {
	sstInstalled := p.SetupFramework()
	if !sstInstalled {
		err := install_sst(p.Tooling, "")
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	}
	color.Green("Project Setup Complete! 🚀")
}

// SetupFramework will setup the framework and return if sst was installed in the process.
func (p *Project) SetupFramework() bool {
	switch p.Framework {
	case "Standalone":
		err := install_sst(p.Tooling, p.Name)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
		return true
	case "Next":
		err := next_create_new_app(p.Name, p.Tooling)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
	case "Remix":
		err := remix_create_new_app(p.Name, p.Tooling)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
	case "Svelte":
		err := svelte_create_new_app(p.Name, p.Tooling)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
	case "Astro":
		err := astro_create_new_app(p.Name, p.Tooling)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
	case "Solid":
		err := solid_create_new_app(p.Name, p.Tooling)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		os.Chdir(p.Name)
	}

	return false
}
