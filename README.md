# Auth Client

This package provides a robust client for interacting with authentication and authorization endpoints. It simplifies the process of checking user permissions, roles, and interface access in your Go applications.

## Features

- Check interface authentication
- Verify user permissions
- Check user roles
- Easy-to-use API
- Configurable base URL and authentication token

## Installation

To install this package, use the `go get` command:

```bash
go get github.com/qucheng-tony/authClient
```

## Usage

Here's a quick example of how to use the Auth Client:

```go
package main

import (
    "fmt"
    "github.com/qucheng-tony/authClient"
)

func main() {
    // Initialize the client with your API base URL and authentication token
    client := authclient.NewAuthClient("http://your-api-base-url", "your-auth-token")

    // Check interface authentication
    authReq := authclient.CheckInterfaceAuthReq{
        Method:    1, // 1: GET, 2: POST, 3: PUT, 4: DELETE
        Path:      "/some/path",
        UserID:    123,
        ProjectID: 456,
    }

    hasAuth, err := client.CheckInterfaceAuth(authReq)
    if err != nil {
        fmt.Printf("Error checking auth: %v\n", err)
        return
    }

    if hasAuth {
        fmt.Println("User has auth for this interface")
    } else {
        fmt.Println("User doesn't have auth for this interface")
    }

    // Check user permissions
    permReq := authclient.CheckHasPermissionReq{
        UserID:         123,
        ProjectID:      456,
        PermissionList: []string{"read", "write"},
    }

    hasPermission, err := client.HasAnyPermission(permReq)
    if err != nil {
        fmt.Printf("Error checking permissions: %v\n", err)
        return
    }

    fmt.Printf("User has required permissions: %v\n", hasPermission)

    // Check user roles
    roleReq := authclient.CheckHasRoleReq{
        UserID:     123,
        ProjectID:  456,
        RoleIDList: []int{1, 2, 3},
    }

    hasRole, err := client.HasAnyRole(roleReq)
    if err != nil {
        fmt.Printf("Error checking roles: %v\n", err)
        return
    }

    fmt.Printf("User has required roles: %v\n", hasRole)
}
```

## API Reference

### `NewAuthClient(baseURL, token string) *AuthClient`

Creates a new instance of the Auth Client.

### `CheckInterfaceAuth(req CheckInterfaceAuthReq) (bool, error)`

Checks if a user has authentication for a specific interface.

### `HasAnyPermission(req CheckHasPermissionReq) (bool, error)`

Checks if a user has any of the specified permissions.

### `HasAnyRole(req CheckHasRoleReq) (bool, error)`

Checks if a user has any of the specified roles.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Support

If you encounter any problems or have any questions, please open an issue on the [GitHub repository](https://github.com/qucheng-tony/authClient/issues).