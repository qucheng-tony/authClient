package test

import (
	"fmt"
	"testing"

	"github.com/qucheng-tony/authClient"
)

func TestClient(t *testing.T) {
	server := authClient.NewAuthClient("a", "b")
	fmt.Println(server)
}
