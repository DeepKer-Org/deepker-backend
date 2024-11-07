package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the structure for the notification payload
type ExpoPushMessage struct {
	To       []string `json:"to"`
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Priority string   `json:"priority"`
	Sound    string   `json:"sound"`
}

func SendExponentPushNotifications(pushTokens []string, title string, body string) error {
	// Define the API endpoint for Expo push notifications
	url := "https://exp.host/--/api/v2/push/send"

	// Create the payload with the list of tokens, title, body, and other settings
	message := ExpoPushMessage{
		To:       pushTokens,
		Title:    title,
		Body:     body,
		Priority: "high",
		Sound:    "default",
	}

	// Encode the payload as JSON
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal notification payload: %w", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notifications, status code: %d", resp.StatusCode)
	}

	fmt.Println("Notifications sent successfully")
	return nil
}
