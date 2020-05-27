package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/shachat/router"
	"github.com/stretchr/testify/assert"
)

func TestPingPong(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Errorf("unexpected error:\n%s", err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
