package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Command struct {
	Cmd         string `yaml:"cmd"`
	Description string `yaml:"description,omitempty"`
}

type Commands struct {
	Commands map[string]interface{} `yaml:"commands"`
}

func showAvailableApps() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	entries, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	var availableApps []string
	for _, entry := range entries {
		if entry.IsDir() {
			// Check if commands.yml exists in this directory
			commandsFile := filepath.Join(currentDir, entry.Name(), "commands.yml")
			if _, err := os.Stat(commandsFile); err == nil {
				availableApps = append(availableApps, entry.Name())
			}
		}
	}

	if len(availableApps) == 0 {
		fmt.Println("No applications found with commands.yml files")
		fmt.Println("\nUsage: commander <app> [command]")
		fmt.Println("Example: commander caddy reload")
		return
	}

	fmt.Println("Available applications:")
	for _, app := range availableApps {
		fmt.Printf("  %s\n", app)
	}
	fmt.Println("\nUse 'commander <app>' to see available commands for an application")
}

func main() {
	if len(os.Args) < 2 {
		// No arguments - show available applications
		showAvailableApps()
		return
	}

	appName := os.Args[1]
	var commandName string
	if len(os.Args) >= 3 {
		commandName = os.Args[2]
	}

	// Get current directory (should be ~/apps)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Construct app directory path
	appDir := filepath.Join(currentDir, appName)

	// Check if app directory exists
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		fmt.Printf("App directory '%s' does not exist\n", appDir)
		os.Exit(1)
	}

	// Load commands.yml from app directory
	commandsFile := filepath.Join(appDir, "commands.yml")
	if _, err := os.Stat(commandsFile); os.IsNotExist(err) {
		fmt.Printf("commands.yml not found in %s\n", appDir)
		os.Exit(1)
	}

	// Read and parse YAML file
	data, err := os.ReadFile(commandsFile)
	if err != nil {
		fmt.Printf("Error reading commands.yml: %v\n", err)
		os.Exit(1)
	}

	var commands Commands
	err = yaml.Unmarshal(data, &commands)
	if err != nil {
		fmt.Printf("Error parsing commands.yml: %v\n", err)
		os.Exit(1)
	}

	// Parse commands to handle both string and object formats
	parsedCommands := make(map[string]Command)
	for name, rawCmd := range commands.Commands {
		switch v := rawCmd.(type) {
		case string:
			// Simple string format: "command string"
			parsedCommands[name] = Command{Cmd: v}
		case map[string]interface{}:
			// Object format with cmd and description
			cmd := Command{}
			if cmdStr, ok := v["cmd"].(string); ok {
				cmd.Cmd = cmdStr
			}
			if desc, ok := v["description"].(string); ok {
				cmd.Description = desc
			}
			parsedCommands[name] = cmd
		default:
			fmt.Printf("Invalid command format for '%s' in commands.yml\n", name)
			os.Exit(1)
		}
	}

	// Find the requested command or list all if no command provided
	if commandName == "" {
		// List all available commands
		fmt.Printf("Available commands for '%s':\n", appName)
		for name, cmd := range parsedCommands {
			if cmd.Description != "" {
				fmt.Printf("  %-15s %s - %s\n", name, cmd.Cmd, cmd.Description)
			} else {
				fmt.Printf("  %-15s %s\n", name, cmd.Cmd)
			}
		}
		return
	}

	command, exists := parsedCommands[commandName]
	if !exists {
		fmt.Printf("Command '%s' not found for app '%s'\n", commandName, appName)
		fmt.Printf("\nAvailable commands for '%s':\n", appName)
		for name, cmd := range parsedCommands {
			if cmd.Description != "" {
				fmt.Printf("  %-15s %s - %s\n", name, cmd.Cmd, cmd.Description)
			} else {
				fmt.Printf("  %-15s %s\n", name, cmd.Cmd)
			}
		}
		os.Exit(1)
	}

	// Execute the command in the app directory
	fmt.Printf("Running: %s (in %s)\n", command.Cmd, appDir)
	
	// Split command into parts for exec
	parts := strings.Fields(command.Cmd)
	if len(parts) == 0 {
		fmt.Println("Empty command")
		os.Exit(1)
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = appDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
		os.Exit(1)
	}
}
