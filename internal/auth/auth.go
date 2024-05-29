package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no authentication header provided")
	}
	parts := strings.Split(apiKey, " ")
	if len(parts) != 2 {
		return "", errors.New("malformed auth header")
	}
	if parts[0] != "ApiKey" {
		return "", errors.New("invalid authentication header")
	}
	return parts[1], nil
}