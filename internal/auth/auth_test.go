package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input  http.Header
		outStr string
		outErr string
	}

	tests := []test{
		{
			input:  http.Header{"Authorization": []string{"ApiKey MyAPIKey"}},
			outStr: "MyAPIKey",
			outErr: "",
		},
		{
			input:  http.Header{"Authorization": []string{"NotApiKey MyAPIKey"}},
			outStr: "",
			outErr: "malformed authorization header",
		},
		{
			input:  http.Header{"Other": []string{"ApiKey MyAPIKey"}},
			outStr: "",
			outErr: "no authorization header included",
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.input.Get("Authorization"))
		got, err := GetAPIKey(tc.input)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}

		if got != tc.outStr || errStr != tc.outErr {
			t.Fatalf("expected (error): %v (%v), got (error): %v (%v)", tc.outStr, tc.outErr, got, err)
		}
	}
}
