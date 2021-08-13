package demo_test

import (
	"demo"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	server := demo.New(demo.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.HandleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello world")
}