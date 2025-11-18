package common

import (
	"encoding/json"
	"fmt"
	"os"
)

const messagesDBPath = "./data/messages.json"

type DbService struct {
	dbPath string
}

func NewDbService() (*DbService, error) {
	if err := os.MkdirAll("./data", 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	return &DbService{
		dbPath: messagesDBPath,
	}, nil
}

func (db *DbService) SaveMessages(messages []Message) error {
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal messages: %w", err)
	}
	if err := os.WriteFile(db.dbPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write messages to file: %w", err)
	}

	return nil
}

func (db *DbService) LoadMessages() ([]Message, error) {
	if _, err := os.Stat(db.dbPath); os.IsNotExist(err) {
		return []Message{}, nil
	}
	data, err := os.ReadFile(db.dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read messages from file: %w", err)
	}
	var messages []Message
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, fmt.Errorf("failed to unmarshal messages: %w", err)
	}
	return messages, nil
}
