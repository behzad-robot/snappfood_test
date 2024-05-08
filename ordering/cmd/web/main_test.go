package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_IntergationAgents(t *testing.T) {
	go main()
	time.Sleep(2 * time.Second)
	const URL = "http://localhost:3000"
	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get(URL)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})
	t.Run("Get Agent 1", func(t *testing.T) {
		t.Log("Get Agent 1")
		resp, err := http.Get(URL + "/api/agents/1")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})
	t.Run("Get Agent 1700", func(t *testing.T) {
		t.Log("Get Agent 1700")
		resp, err := http.Get(URL + "/api/agents/1700")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, 404, resp.StatusCode)
	})
}
