package main

import (
	"context"
	"fmt"
	"os"
	"google.golang.org/genai"
	"hdi/harness"
	"os/exec"
	"strings"
)

func main() {
	ctx := context.Background()

	userQuestion := os.Args[2]
	helpPageContent := getHelpPageContent(os.Args[1]) 

	prompt := buildPrompt(helpPageContent, userQuestion)

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GENAI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Printf("Error creating GenAI client: %v\n", err)
		return
	}

	agent := harness.New(client, "gemini-3.5-flash-lite")
	
	response, err := agent.AskLLM(ctx, prompt)
	if err != nil {
		fmt.Printf("Error asking LLM: %v\n", err)
		return
	}

	response = postProcessResponse(response)

	fmt.Println("hdi:\r\n", response)
}

func buildPrompt(helpPage string, userQuestion string) string {
	return fmt.Sprintf("Help page Documentation:\n%s\n\nUser Question: %s", helpPage, userQuestion)
} 


func postProcessResponse(response string) string {
	response = strings.ReplaceAll(response, "```", "`")
	return response
}

func getHelpPageContent(clitool string) string {
	helpPage := exec.Command(clitool, "--help")

	output, err := helpPage.Output()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	return string(output)
}
