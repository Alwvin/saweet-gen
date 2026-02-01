package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func Input() {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid Project Name")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Project Name",
		Validate: validate,
	}

	result, _ := prompt.Run()
	create_dir_err := os.Mkdir(result, 0750)
	if create_dir_err != nil {
		panic("Cannot create directory")
	}

	fmt.Printf("You choose %q\n", result)
}
