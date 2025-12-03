package oceanengine

import (
	"strings"
	"testing"
)

func TestOAuthServiceGetAuthURL(t *testing.T) {
	client := NewClient("test_app_id", "test_secret")
	oauthSvc := NewOAuthService(client)

	tests := []struct {
		name         string
		redirectURI  string
		state        string
		wantContains []string
	}{
		{
			name:        "basic auth URL",
			redirectURI: "https://example.com/callback",
			state:       "test_state",
			wantContains: []string{
				"https://ad.oceanengine.com/openapi/audit/oauth.html",
				"app_id=test_app_id",
				"redirect_uri=https://example.com/callback",
				"state=test_state",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oauthSvc.GetAuthURL(tt.state, tt.redirectURI)
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("GetAuthURL() = %v, want to contain %v", got, want)
				}
			}
		})
	}
}

func TestOAuthServiceGetAuthURLWithScope(t *testing.T) {
	client := NewClient("test_app_id", "test_secret")
	oauthSvc := NewOAuthService(client)

	tests := []struct {
		name         string
		redirectURI  string
		state        string
		scope        []string
		materialAuth bool
		wantContains []string
	}{
		{
			name:         "auth URL with scope and material_auth",
			redirectURI:  "https://example.com/callback",
			state:        "test_state",
			scope:        []string{"ad_manage", "ad_read"},
			materialAuth: true,
			wantContains: []string{
				"app_id=test_app_id",
				"scope=ad_manage,ad_read",
				"material_auth=1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oauthSvc.GetAuthURLWithScope(tt.state, tt.redirectURI, tt.scope, tt.materialAuth)
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("GetAuthURLWithScope() = %v, want to contain %v", got, want)
				}
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient("my_app_id", "my_secret")

	if client.appID != "my_app_id" {
		t.Errorf("NewClient() appID = %v, want my_app_id", client.appID)
	}
	if client.secret != "my_secret" {
		t.Errorf("NewClient() secret = %v, want my_secret", client.secret)
	}
	if client.httpClient == nil {
		t.Error("NewClient() httpClient should not be nil")
	}
}
