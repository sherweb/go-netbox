package main

import (
	"os"
	"testing"

	"github.com/sherweb/go-netbox/v4/netbox"
)

func HGetClient(t *testing.T) *netbox.APIClient {
	srv := os.Getenv("NETBOX_URL")
	token := os.Getenv("NETBOX_TOKEN")

	if srv == "" {
		t.Fatal("NETBOX_URL is not set")
	}
	if token == "" {
		t.Fatal("NETBOX_TOKEN is not set")
	}

	cfg := netbox.NewConfiguration()
	cfg.Servers[0].URL = srv
	cfg.AddDefaultHeader("Authorization", "Token "+token)

	client := netbox.NewAPIClient(cfg)
	return client

}
