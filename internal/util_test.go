package internal

import (
	"testing"
)


func Test_parseAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		want    string
		wantErr bool
	}{
		// Invalid API keys
		{"too short", "abc123", "", true},
		{"too long", "111222333444555666777888999000aaabbbcccdddeeefffggg111222333444555666777888999000aaabbbcccdddeeefffggg", "", true},
		{"invalid characters", "invalid-key!", "", true},
		// Valid API keys
		{"ok", "bgmsElBKJgo5GJMDL1121R0UiOEj86RpyM", "bgmsElBKJgo5GJMDL1121R0UiOEj86RpyM", false},
		{"ok with whitespace", " bgmsElBKJgo5GJMDL1121R0UiOEj86RpyM     ", "bgmsElBKJgo5GJMDL1121R0UiOEj86RpyM", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAPIKey(tt.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		want    string
		wantErr bool
	}{
		// Invalid URLs
		{"empty", "", "", true},
		{"missing scheme", "api.example.com/path", "", true},
		{"invalid scheme", "ftp://api.example.com/path", "", true},
		{"invalid scheme", "http://api.example.com/path", "", true},
		// Valid URLs
		{"valid", "https://api.example.com/path", "https://api.example.com/path", false},
		{"valid with whitespace", "   https://api.example.com/path \n", "https://api.example.com/path", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseURL(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}