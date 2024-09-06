package test

import (
	"fmt"
	"testing"

	"github.com/qucheng-tony/authClient"
)

func TestClient(t *testing.T) {
	server := authClient.NewAuthClient("https://invengo-auth.rfidtour.com/api/auth/permissionManage",
		"Authorization:token:d6a6d699-b1dd-4c1e-bf92-1e9b184aade5")
	ok, err := server.CheckInterfaceAuth(authClient.CheckInterfaceAuthReq{
		Method:    1,
		Path:      "b",
		UserID:    45,
		ProjectID: 4,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)
}
