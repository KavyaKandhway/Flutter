package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInit(t *testing.T) {
	r := gin.New()
	middleware, err := New("../credentials.json", nil)
	if err != nil {
		t.Error(err)
	}
	r.Use(middleware.TokenValidation())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer example_token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("w.Code should be 401:%d", w.Code)
	}

}
