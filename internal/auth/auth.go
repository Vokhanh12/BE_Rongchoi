package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Missing authorization header")
	}

	parts := strings.SplitN(val, " ", 2)
	if len(parts) != 2 {
		return "", errors.New("Malformed authorization header")
	}

	const TYPE = "ApiKey"
	authType := strings.TrimSpace(parts[0])
	if authType != TYPE {
		return "", errors.New("Unsupported authorization type")
	}

	apiKey := strings.TrimSpace(parts[1])
	if apiKey == "" {
		return "", errors.New("Empty API key")
	}

	return apiKey, nil
}

// GetBearer extracts a Bearer token from
// the headers of an HTTP request
// Example:
// Authorization: Bearer {insert token here}
func GetBearer(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Missing authorization header")
	}

	parts := strings.SplitN(val, " ", 2)
	if len(parts) != 2 {
		return "", errors.New("Malformed authorization header")
	}

	const TYPE = "Bearer"
	authType := strings.TrimSpace(parts[0])
	if authType != TYPE {
		return "", errors.New("Unsupported authorization type")
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("Empty token")
	}

	return token, nil
}
