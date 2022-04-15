package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_helloHandler(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"base case", `{"name": "liwenzhou"}`, "hello liwenzhou"},
		{"bad case", "", "we need a name"},
	}
	r := SetupRouter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			req := httptest.NewRequest(
				"POST",
				"/hello",
				strings.NewReader(tt.param))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)

			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})

	}
}
