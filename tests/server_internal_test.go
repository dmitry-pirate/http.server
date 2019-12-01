package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"vpn_api/app/server"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestApiServer_HandlePac(t *testing.T) {
	s := server.New(server.NewConfig())
	body := gin.H{
		"mode": "vpn",
	}
	router := s.GetRouter()
	w := performRequest(router, http.MethodPost, "/pac")
	var response map[string]string
	r := w.Body.String()
	err := json.Unmarshal([]byte(r), &response)
	value, exists := response["mode"]
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["mode"], value)
}

func TestApiServer_HandlePing(t *testing.T) {
	s := server.New(server.NewConfig())
	body := gin.H{
		"mode": "vpn",
	}
	router := s.GetRouter()
	w := performRequest(router, http.MethodPost, "/ping")
	var response map[string]string
	r := w.Body.String()
	err := json.Unmarshal([]byte(r), &response)
	value, exists := response["mode"]
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["mode"], value)
}
