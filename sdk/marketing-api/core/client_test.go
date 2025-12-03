package core

import (
	"net/http"
	"testing"
	"time"
)

func TestNewSDKClient(t *testing.T) {
	appID := uint64(12345678)
	secret := "test_secret"

	client := NewSDKClient(appID, secret)

	if client == nil {
		t.Fatal("NewSDKClient returned nil")
	}

	if client.AppID != appID {
		t.Errorf("AppID = %d, want %d", client.AppID, appID)
	}

	if client.Secret != secret {
		t.Errorf("Secret = %s, want %s", client.Secret, secret)
	}

	if client.client == nil {
		t.Error("http client is nil")
	}
}

func TestSDKClient_SetDebug(t *testing.T) {
	client := NewSDKClient(12345678, "secret")

	client.SetDebug(true)
	if !client.debug {
		t.Error("SetDebug(true) did not set debug to true")
	}

	client.SetDebug(false)
	if client.debug {
		t.Error("SetDebug(false) did not set debug to false")
	}
}

func TestSDKClient_SetHttpClient(t *testing.T) {
	client := NewSDKClient(12345678, "secret")

	customClient := &http.Client{
		Timeout: time.Second * 30,
	}

	client.SetHttpClient(customClient)

	if client.client != customClient {
		t.Error("SetHttpClient did not set custom client")
	}
}

func TestSDKClient_Sandbox(t *testing.T) {
	client := NewSDKClient(12345678, "secret")

	if client.sandbox {
		t.Error("sandbox should be false by default")
	}

	client.UseSandbox()
	if !client.sandbox {
		t.Error("UseSandbox() did not enable sandbox")
	}

	client.DisableSandbox()
	if client.sandbox {
		t.Error("DisableSandbox() did not disable sandbox")
	}
}

func TestSDKClient_SetOperatorIP(t *testing.T) {
	client := NewSDKClient(12345678, "secret")

	testIP := "192.168.1.1"
	client.SetOperatorIP(testIP)

	if client.operatorIP != testIP {
		t.Errorf("operatorIP = %s, want %s", client.operatorIP, testIP)
	}
}

func TestSDKClient_Copy(t *testing.T) {
	client := NewSDKClient(12345678, "secret")
	client.SetDebug(true)
	client.UseSandbox()
	client.SetOperatorIP("127.0.0.1")

	copied := client.Copy()

	if copied == client {
		t.Error("Copy() returned same pointer")
	}

	if copied.AppID != client.AppID {
		t.Errorf("copied.AppID = %d, want %d", copied.AppID, client.AppID)
	}

	if copied.Secret != client.Secret {
		t.Errorf("copied.Secret = %s, want %s", copied.Secret, client.Secret)
	}

	if copied.debug != client.debug {
		t.Errorf("copied.debug = %v, want %v", copied.debug, client.debug)
	}

	if copied.sandbox != client.sandbox {
		t.Errorf("copied.sandbox = %v, want %v", copied.sandbox, client.sandbox)
	}

	if copied.operatorIP != client.operatorIP {
		t.Errorf("copied.operatorIP = %s, want %s", copied.operatorIP, client.operatorIP)
	}
}

func TestDefaultHttpClient(t *testing.T) {
	client1 := defaultHttpClient()
	client2 := defaultHttpClient()

	if client1 != client2 {
		t.Error("defaultHttpClient() should return singleton")
	}

	if client1 == nil {
		t.Error("defaultHttpClient() returned nil")
	}

	if client1.Timeout != time.Second*60 {
		t.Errorf("client timeout = %v, want %v", client1.Timeout, time.Second*60)
	}
}
