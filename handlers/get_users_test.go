package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/shachat/router"
	"github.com/stretchr/testify/assert"
)

func TestAllGetUsers(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("unexpected error:%s\n", err)
	}
	router.ServeHTTP(w, req)

	assert := assert.New(t)

	assert.Equal(http.StatusOK, w.Code)
	assert.NotEqual("", w.Body.String())
}

func TestSomeGetUsers(t *testing.T) {

}

func TestFailGetUsers(t *testing.T) {

}
