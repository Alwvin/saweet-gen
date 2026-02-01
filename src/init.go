package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var name_dir string

func CheckOs(detect_os string) {
	switch detect_os {
	case "windows":
		name_dir = "\\.saweet-template\\"
	case "linux", "android":
		name_dir = "/.saweet-template/"
	default:
		fmt.Println("Unknown Operating System")
	}
}

func Init() {
	home_dir, _ := os.UserHomeDir()

	CheckOs(runtime.GOOS)

	// create directory for contains template
	err := os.Mkdir(home_dir+name_dir, 0750)

	if os.IsExist(err) {
		color.Blue("Welcome to Saweet-gen")
	} else {
		color.Red("Initialize")
	}

	get_err_env := godotenv.Load()
	if get_err_env != nil {
		panic("Cannot found .env file")
	}

	// for custom url
	uri := os.Getenv("URL_TEMPLATE")

	GetTemplate(uri, home_dir+name_dir)
	Input()
}
