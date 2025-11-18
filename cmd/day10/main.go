package main

import (
	"aichallenge/internal/day10"
	//"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: day10 <mcp-server-command> [args...]")
		fmt.Println()
		fmt.Println("Example:")
		//day10 npx -y @modelcontextprotocol/server-filesystem /Users/sergeyusachev/Projects/GoProjects/AIGladkovChallenge/scripts
		fmt.Println("day10 npx -y @modelcontextprotocol/server-filesystem /Users/sergeyusachev/Projects/GoProjects/AIGladkovChallenge/scripts")
		fmt.Println("day10 uvx mcp-server-sqlite --db-path /path/to/db.sqlite")
		os.Exit(1)
	}

	serverCommand := os.Args[1]
	serverArgs := []string{}
	if len(os.Args) > 2 {
		serverArgs = os.Args[2:]
	}

	fmt.Printf("Connecting to MCP server: %s %v\n", serverCommand, serverArgs)
	fmt.Println()

	client, err := day10.NewMCPClient(serverCommand, serverArgs...)
	if err != nil {
		fmt.Printf("Failed to create MCP client: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Println("Initializing MCP connection...")
	if err := client.Initialize(); err != nil {
		fmt.Printf("Failed to initialize MCP connection: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully initialized MCP connection")
	fmt.Println()

	fmt.Println("Fetching available tools...")
	tools, err := client.ListTools()
	if err != nil {
		fmt.Printf("Failed to list tools: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d tools:\n", len(tools))
	fmt.Println()

	for i, tool := range tools {
		fmt.Printf("%d. %s\n", i+1, tool.Name)
		/*if tool.Description != "" {
			fmt.Printf("   Description: %s\n", tool.Description)
		}*/
		/*if tool.InputSchema != nil {
			schemaJSON, _ := json.MarshalIndent(tool.InputSchema, "   ", "  ")
			fmt.Printf("   Input Schema: %s\n", string(schemaJSON))
		}*/
		//fmt.Println()
	}
}
