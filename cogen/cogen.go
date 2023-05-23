package cogen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	cogenDir     = ".cogen"
	configDir    = "config"
	projectFile  = "project"
	planFile     = "plan"
	definePrompt = `System: COGEN
Project Name: cogen

# The description describes your project system giving additional details
Description: An autocoding command line system

# The goal is overarching guidance to focus and hone the end results
Goal: To create a clear, concise and professional solution

# The tech stack includes the languages, tooling, and technologies that your project system is implemented with
Tech Stack: Golang

# The dependencies are the tech stack specific list of dependencies that the system should ensure is included in the end result
Dependencies: github.com/sashabaranov/go-openai

# Primary project implementation language
Language: Go

# The following lays down specific 
All code should be written in the tech stack above. 
Requirements should be implemented in the language specified above.
Be sure to use SOLID principals.

# Personas are guiding examples of coding styles that you want the system to follow
Personas: Combine the personas Rob Pike, Brad Fitzpatrick and Eric Evans. 
`
)

// Init initializes a new cogen project
func Init(projectName string) error {
	if err := os.MkdirAll(cogenDir+"/"+configDir, os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(cogenDir+"/"+configDir+"/"+projectFile, []byte("System: COGEN\nProject Name: "+projectName+"\n"), os.ModePerm)
}

// Define defines the project configuration
func Define() error {
	if _, err := os.Stat(cogenDir + "/" + configDir); os.IsNotExist(err) {
		return errors.New("cogen config directory not found, run 'cogen init' first")
	}

	tempFile, err := ioutil.TempFile("", "cogen-define-")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(definePrompt); err != nil {
		return err
	}
	if err := tempFile.Close(); err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	editedContent, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cogenDir+"/"+configDir+"/"+projectFile, editedContent, os.ModePerm)
}

func checkConfigDir() error {
	if _, err := os.Stat(cogenDir + "/" + configDir); os.IsNotExist(err) {
		return errors.New("config directory not found, run 'cogen init' first")
	}
	return nil
}

func openEditorWithPrompt(prompt string) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	tempFile, err := ioutil.TempFile("", "cogen-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(prompt); err != nil {
		return nil, err
	}
	if err := tempFile.Close(); err != nil {
		return nil, err
	}

	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(tempFile.Name())
}

// Plan creates a project plan
func Plan() error {
	if err := checkConfigDir(); err != nil {
		return err
	}

	prompt := `Provide an overall solution following the guidance mentioned above. 
Don’t generate code. 
Describe the solution, break it down as a detailed numbered task list based on the guidance mentioned above. 
And we will refer this task list as the project plan.
`

	editedText, err := openEditorWithPrompt(prompt)
	if err != nil {
		return err
	}

	if len(editedText) == 0 {
		return errors.New("empty plan, please provide a valid plan")
	}

	planFilePath := cogenDir + "/" + configDir + "/" + planFile
	return ioutil.WriteFile(planFilePath, editedText, 0644)
}

// Refine refines the project
func Refine() error {
	configPath := filepath.Join(".cogen", "config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config directory not found, run 'cogen init' first")
	}

	timestamp := time.Now().Format("20060102150405")
	timestampDir := filepath.Join(".cogen", timestamp)
	err := os.MkdirAll(timestampDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create timestamp directory: %v", err)
	}

	promptPath := filepath.Join(timestampDir, "prompt")
	responsePath := filepath.Join(timestampDir, "response")

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, promptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open editor: %v", err)
	}

	promptContent, err := ioutil.ReadFile(promptPath)
	if err != nil {
		return fmt.Errorf("failed to read prompt file: %v", err)
	}

	if len(promptContent) == 0 {
		return fmt.Errorf("empty prompt file, refine operation aborted")
	}

	// Save the response to the responsePath
	// Replace this with the actual response generation logic
	responseContent := "Generated response content"
	err = ioutil.WriteFile(responsePath, []byte(responseContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write response file: %v", err)
	}

	return nil
}
