package authClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthClient struct {
	BaseURL string
	Token   string
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type CheckInterfaceAuthReq struct {
	Method    int8   `json:"method"`
	Path      string `json:"path"`
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"`
}

type CheckHasPermissionReq struct {
	UserID         int      `json:"user_id"`
	ProjectID      int      `json:"project_id"`
	PermissionList []string `json:"permission_list"`
}

type CheckHasRoleReq struct {
	UserID     int   `json:"user_id"`
	ProjectID  int   `json:"project_id"`
	RoleIDList []int `json:"role_id_list"`
}

func NewAuthClient(baseURL, token string) *AuthClient {
	return &AuthClient{BaseURL: baseURL, Token: token}
}

func (c *AuthClient) CheckInterfaceAuth(req CheckInterfaceAuthReq) (bool, error) {
	return c.makeRequest("GET", "/checkInterfaceAuth", req)
}

func (c *AuthClient) HasAnyPermission(req CheckHasPermissionReq) (bool, error) {
	return c.makeRequest("POST", "/hasAnyPermission", req)
}

func (c *AuthClient) HasAnyRole(req CheckHasRoleReq) (bool, error) {
	return c.makeRequest("POST", "/hasAnyRole", req)
}

func (c *AuthClient) makeRequest(method, path string, payload interface{}) (bool, error) {
	url := c.BaseURL + path
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("error marshaling payload: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, fmt.Errorf("error decoding response: %w", err)
	}

	if response.Code != 2000 {
		return false, fmt.Errorf("unexpected response code: %d, message: %s", response.Code, response.Msg)
	}

	hasPermission, ok := response.Data.(bool)
	if !ok {
		return false, fmt.Errorf("unexpected data type in response")
	}

	return hasPermission, nil
}
