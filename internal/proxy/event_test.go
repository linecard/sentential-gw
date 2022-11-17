package proxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestBuildEvent(t *testing.T) {
	requests := []http.Request{
		{
			Method: "GET",
			URL: &url.URL{
				Path: "/",
			},
			Header: http.Header{},
			Body:   io.NopCloser(strings.NewReader("")),
		},
		{
			Method: "POST",
			URL: &url.URL{
				Path:     "/device/add",
				RawQuery: "name=humidity&location=rooftop",
			},
			Header: http.Header{
				"Content-Type": {"application/json"},
				"Cookie":       {"SESSIONID=abcde12345"},
			},
			RemoteAddr: "127.0.0.1:12345",
			Body:       io.NopCloser(strings.NewReader("{\"sensor_id\": \"abc123\"}")),
		},
	}

	for _, request := range requests {
		_, err := BuildEvent(&request)
		if err != nil {
			t.Errorf("error building event: %v", err)
		}
	}
}
