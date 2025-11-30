package consumer

import (
	"encoding/json"
	"fmt"

	"queue-worker/httpclient"
	"queue-worker/models"
)

func ProcessMessage(msg []byte, client *httpclient.APIClient) error {
	var event models.WeatherEvent

	if err := json.Unmarshal(msg, &event); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	if event.Timestamp == "" {
		return fmt.Errorf("missing timestamp")
	}

	if err := client.PostWeather(event); err != nil {
		return err
	}

	return nil
}
