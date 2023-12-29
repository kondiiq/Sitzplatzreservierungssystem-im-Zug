package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPingRoute(t *testing.T) {
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/route", nil)
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
	req2, _ := http.NewRequest("GET", "/train", nil)
	router.ServeHTTP(w, req2)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
	req3, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req3)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
	req4, _ := http.NewRequest("GET", "/ticket", nil)
	router.ServeHTTP(w, req4)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
