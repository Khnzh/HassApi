package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	hassBaseURL  = "http://192.168.1.18:8123"
	automationID = "1773179468507"
)

type hassAutomation struct {
	ID          string `json:"id"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
	Triggers    []any  `json:"triggers"`
	Conditions  []any  `json:"conditions"`
	Actions     []any  `json:"actions"`
	Mode        string `json:"mode"`
}

func updateHassAutomationTime(token, hh, mm string) error {
	client := &http.Client{}

	body, err := hassGet(client, token, "/api/config/automation/config/"+automationID)
	if err != nil {
		return fmt.Errorf("fetch automation: %w", err)
	}

	var auto hassAutomation
	if err := json.Unmarshal(body, &auto); err != nil {
		return fmt.Errorf("parse automation: %w", err)
	}

	auto.Triggers = []any{
		map[string]any{
			"trigger": "time",
			"at":      hh + ":" + mm + ":00",
		},
	}

	payload, _ := json.Marshal(auto)
	if _, err := hassPost(client, token, "/api/config/automation/config/"+automationID, string(payload)); err != nil {
		return fmt.Errorf("update automation: %w", err)
	}

	if _, err := hassPost(client, token, "/api/services/automation/reload", "{}"); err != nil {
		return fmt.Errorf("reload automations: %w", err)
	}

	return nil
}

func hassGet(client *http.Client, token, path string) ([]byte, error) {
	req, err := http.NewRequest("GET", hassBaseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func hassPost(client *http.Client, token, path, payload string) ([]byte, error) {
	req, err := http.NewRequest("POST", hassBaseURL+path, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
