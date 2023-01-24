package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockHandler struct{}

func (h *mockHandler) about(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"about": "this auth service api"})
}

func (h *mockHandler) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func TestAbout(t *testing.T) {
	h := &mockHandler{}
	r := gin.Default()
	r.GET("/about", h.about)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHealthcheck(t *testing.T) {
	h := &mockHandler{}
	r := gin.Default()
	r.GET("/healthcheck", h.healthcheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
