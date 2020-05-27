package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/shachat/router"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("unexpected error:%s\n", err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}
