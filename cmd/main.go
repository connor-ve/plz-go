package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var display bool = true
var default_path string = "C:\\plz_scripts\\"

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
	}
	path := parseInputs(os.Args)
	fmt.Println(path)
	if path == "" {
		os.Exit(0)
	}
	cmd := exec.Command(path)
	if display {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Execution Failed : %s", err)
	}
}

func parseInputs(args []string) string {
	var path string
	for _, v := range args {
		if v[0] != '-' {
			path = v
			continue
		}
		if v == "-h" || v == "--help" {
			manual := `Welcome to PLZ
A simple command to keep basic bat files, and run them! 
-s, --silent : turns off output
-v, --view   : shows the plz_scripts files
-l, --list   : shows the bat files in the current folder`
			fmt.Print(manual)
			return ""
		}
		if v == "-s" || v == "--silent" {
			display = false
		}
		if v == "-l" || v == "--list" {
			list, _ := os.ReadDir(".")
			path = selectFile(list, ".\\")
			break
		}
		if v == "-v" || v == "--view" {
			list, _ := os.ReadDir(default_path)
			path = selectFile(list, default_path)
			break
		}
	}
	return path
}

func selectFile(list []fs.DirEntry, path string) string {
	var batFiles []string
	for _, v := range list {
		if v.IsDir() {
			continue
		}
		if len(v.Name()) > 4 && v.Name()[len(v.Name())-4:] == ".bat" {
			batFiles = append(batFiles, v.Name())
		}
	}

	if len(batFiles) == 0 {
		fmt.Println("No .bat files found.")
		return ""
	}

	var selectedFile string
	prompt := &survey.Select{
		Message: "Choose a .bat file:",
		Options: batFiles,
	}
	err := survey.AskOne(prompt, &selectedFile)
	if err != nil {
		log.Fatalf("Failed to get user input: %s", err)
	}

	return path + selectedFile
}
