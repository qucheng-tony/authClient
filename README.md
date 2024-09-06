# Auth Client

This package provides a client for interacting with authentication and authorization endpoints.

## Installation

To install this package, use:

```
go get e.coding.net/your-team/auth-client
```

## Usage

```go
import "e.coding.net/your-team/auth-client"

client := authclient.NewAuthClient("http://your-api-base-url", "your token")

// Check interface auth
authReq := authclient.CheckInterfaceAuthReq{
    Method:    1,
    Path:      "/some/path",
    UserID:    123,
    ProjectID: 456,
}
hasAuth, err := client.CheckInterfaceAuth(authReq)
if err != nil {
    // Handle error
}
if hasAuth {
    // User has auth for this interface
} else {
    // User doesn't have auth for this interface
}
```

## License

[MIT License](https://opensource.org/licenses/MIT)
