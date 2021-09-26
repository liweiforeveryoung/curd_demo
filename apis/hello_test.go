package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	g := gin.Default()
	SetUpRoutes(g)
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
