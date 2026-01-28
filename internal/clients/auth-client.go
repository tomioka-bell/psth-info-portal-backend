package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"backend/internal/core/models"
)

type ldapAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LdapAuthenticate(username, password string) (models.LdapAuthResponse, error) {
	baseURL := os.Getenv("AUTH_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}

	url := fmt.Sprintf("%s/auth/ldap-authen", baseURL)

	reqBody := ldapAuthRequest{
		Username: username,
		Password: password,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return models.LdapAuthResponse{
			Err:     true,
			Message: "failed to marshal request to auth-service: " + err.Error(),
		}, err
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return models.LdapAuthResponse{
			Err:     true,
			Message: "failed to create request to auth-service: " + err.Error(),
		}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return models.LdapAuthResponse{
			Err:     true,
			Message: "failed to call auth-service: " + err.Error(),
		}, err
	}
	defer resp.Body.Close()

	var res models.LdapAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return models.LdapAuthResponse{
			Err:     true,
			Message: "failed to decode response from auth-service: " + err.Error(),
		}, err
	}

	if resp.StatusCode != http.StatusOK {
		res.Err = true
		if res.Message == "" {
			res.Message = fmt.Sprintf("auth-service returned status %d", resp.StatusCode)
		}
		return res, fmt.Errorf("auth-service error: %s", res.Message)
	}

	return res, nil
}
