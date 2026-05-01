package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testAPIKey, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Errorf(`NewRequest error: %v`, err)
	}
	testAPIKey.Header.Set("Authorization", "ApiKey jieksojc=semovhsef3239==")
	strAPIKEY, err := GetAPIKey(testAPIKey.Header)
	if err != nil {
		t.Errorf(`GetBearerToken error: %v`, err)
	} else {
		t.Logf(`GetBearerToken works well, the bearer token is: %v`, strAPIKEY)
	}
}
