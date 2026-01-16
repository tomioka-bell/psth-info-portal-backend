package clients

import (
	"backend/internal/core/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetEmployeeByEmpCodeFromMicroservice(empCode string) (*models.EmployeeViewResp, error) {
	if empCode == "" {
		return nil, errors.New("empCode is required")
	}
	baseURL := os.Getenv("AUTH_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}

	empCode = strings.ToLower(strings.TrimSpace(empCode))
	endpoint := fmt.Sprintf(
		"%s/api/employee/get-employee-by-emp-code?empCode=%s",
		baseURL,
		url.QueryEscape(empCode),
	)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
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
