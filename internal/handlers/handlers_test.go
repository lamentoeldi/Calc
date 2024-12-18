package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	cases := []struct {
		name           string
		reqBody        []byte
		expectedStatus int
		method         string
	}{
		{
			name:           "Successful request",
			reqBody:        []byte(`{"expression": "2+2*2"}`),
			expectedStatus: http.StatusOK,
			method:         http.MethodPost,
		},
		{
			name:           "Invalid expression",
			reqBody:        []byte(`{"expression": "1/0"}`),
			expectedStatus: http.StatusUnprocessableEntity,
			method:         http.MethodPost,
		},
		{
			name:           "Internal server error",
			reqBody:        nil,
			expectedStatus: http.StatusInternalServerError,
			method:         http.MethodPost,
		},
		{
			name:           "Method not allowed",
			reqBody:        nil,
			expectedStatus: http.StatusMethodNotAllowed,
			method:         http.MethodGet,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(tc.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Method = tc.method

			res := httptest.NewRecorder()

			http.HandlerFunc(CalcHandler).ServeHTTP(res, req)

			if res.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, res.Code)
			}
		})
	}
}
