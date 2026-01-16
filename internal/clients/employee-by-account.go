package clients

import (
	"backend/internal/core/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func FindEmployeeFromMicroservice(account string) (*models.EmployeeViewResp, error) {
	fmt.Println("account : ", account)
	baseURL := os.Getenv("AUTH_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}
	account = strings.ToLower(strings.TrimSpace(account))

	endpoint := fmt.Sprintf(
		"%s/api/employee/find-employee-by-account?account=%s",
		baseURL,
		url.QueryEscape(account),
	)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("employee service returned status %d", resp.StatusCode)
	}

	var wrapper struct {
		Data *models.EmployeeViewResp `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}

	if wrapper.Data == nil {
		return nil, nil
	}

	return wrapper.Data, nil
}
