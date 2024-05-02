package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"variant-inc_dx-demo-go-api/app/api"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Index)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "Successfully performed HTTP request\n"
	assert.Equal(t, expected, rr.Body.String())
}
