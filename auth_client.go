package authClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
	Method    int8   `json:"method"` // 1、GET 2、POST 3、PUT 4、DELETE
	Path      string `json:"path"`
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"` // 根据AUTH系统查看
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
	// 构建查询参数
	params := url.Values{}
	params.Add("method", strconv.Itoa(int(req.Method)))
	params.Add("path", req.Path)
	params.Add("user_id", strconv.Itoa(req.UserID))
	params.Add("project_id", strconv.Itoa(req.ProjectID))

	// 构建完整的 URL，包括查询参数
	fullURL := fmt.Sprintf("%s/checkInterfaceAuth?%s", c.BaseURL, params.Encode())

	// 发送 GET 请求
	return c.makeRequest("GET", fullURL, nil)
}

func (c *AuthClient) HasAnyPermission(req CheckHasPermissionReq) (bool, error) {
	return c.makeRequest("POST", "/hasAnyPermission", req)
}

func (c *AuthClient) HasAnyRole(req CheckHasRoleReq) (bool, error) {
	return c.makeRequest("POST", "/hasAnyRole", req)
}

func (c *AuthClient) makeRequest(method, path string, payload interface{}) (bool, error) {
	var req *http.Request
	var err error

	if method == "GET" {
		// 对于 GET 请求，我们已经在 URL 中包含了参数，所以这里不需要 body
		req, err = http.NewRequest(method, path, nil)
	} else {
		// 对于其他方法，我们需要将 payload 编码为 JSON
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return false, fmt.Errorf("error marshaling payload: %w", err)
		}
		req, err = http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(jsonPayload))
	}

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
