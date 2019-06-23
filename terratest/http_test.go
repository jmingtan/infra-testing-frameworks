package test

import (
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"net"
	"net/http"
	"testing"
)

func TestPort80(t *testing.T) {
	conn, err := net.Dial("tcp", "web:80")
	if err != nil {
		t.Errorf("web:80 is unreachable: %v", err)
	} else {
		defer conn.Close()
	}
}

func TestWeb(t *testing.T) {
	status, _ := http_helper.HttpGet(t, "http://web")
	if status != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, status)
	}
}
