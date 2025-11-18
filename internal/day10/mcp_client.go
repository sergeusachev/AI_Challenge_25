package day10

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
)

func NewMCPClient(serverCommand string, args ...string) (*MCPClient, error) {
	cmd := exec.Command(serverCommand, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start MCP server: %w", err)
	}

	client := &MCPClient{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		reader: bufio.NewReader(stdout),
		nextID: 1,
	}

	return client, nil
}

func (c *MCPClient) Initialize() error {
	params := InitializeParams{
		ProtocolVersion: "2024-11-05",
		Capabilities:    map[string]interface{}{},
		ClientInfo: ClientInfo{
			Name:    "go-mcp-client",
			Version: "1.0.0",
		},
	}

	_, err := c.sendRequest("initialize", params)
	if err != nil {
		return fmt.Errorf("initialize failed: %w", err)
	}
	
	err = c.sendNotification("notifications/initialized", map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("initialized notification failed: %w", err)
	}

	return nil
}

func (c *MCPClient) ListTools() ([]Tool, error) {
	response, err := c.sendRequest("tools/list", map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("tools/list failed: %w", err)
	}

	var result ToolsListResult
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, fmt.Errorf("failed to parse tools list: %w", err)
	}

	return result.Tools, nil
}

func (c *MCPClient) Close() error {
	if c.stdin != nil {
		c.stdin.Close()
	}
	if c.stdout != nil {
		c.stdout.Close()
	}
	if c.cmd != nil && c.cmd.Process != nil {
		c.cmd.Process.Kill()
		c.cmd.Wait()
	}
	return nil
}

func (c *MCPClient) sendRequest(method string, params interface{}) (*JSONRPCResponse, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      c.nextID,
		Method:  method,
		Params:  params,
	}
	c.nextID++

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	requestBytes = append(requestBytes, '\n')

	if _, err := c.stdin.Write(requestBytes); err != nil {
		return nil, fmt.Errorf("failed to write request: %w", err)
	}

	responseBytes, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response JSONRPCResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("MCP error: %s (code: %d)", response.Error.Message, response.Error.Code)
	}

	return &response, nil
}

func (c *MCPClient) sendNotification(method string, params interface{}) error {
	notification := JSONRPCNotification{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	notificationBytes = append(notificationBytes, '\n')

	if _, err := c.stdin.Write(notificationBytes); err != nil {
		return fmt.Errorf("failed to write notification: %w", err)
	}

	return nil
}
