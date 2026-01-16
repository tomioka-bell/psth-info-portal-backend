package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type ldapAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LdapAuthenticate(username, password string) (bool, string) {
	baseURL := os.Getenv("AUTH_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}

	url := fmt.Sprintf("%s/api/auth/domain-login", baseURL)

	reqBody := ldapAuthRequest{
		Username: username,
		Password: password,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return false, "failed to marshal request to auth-service: " + err.Error()
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return false, "failed to create request to auth-service: " + err.Error()
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, "failed to call auth-service: " + err.Error()
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&res)

	msg := ""
	if m, ok := res["error"].(string); ok {
		msg = m
	} else if m, ok := res["message"].(string); ok {
		msg = m
	}

	if resp.StatusCode != http.StatusOK {
		if msg != "" {
			return false, msg
		}
		return false, fmt.Sprintf("auth-service returned status %d", resp.StatusCode)
	}

	success, _ := res["success"].(bool)
	if !success {
		if msg == "" {
			msg = "authentication failed"
		}
		return false, msg
	}

	return true, msg
}
