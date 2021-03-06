package profitbricks

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type tokenHeader struct {
	Typ string
	Kid string
	Alg string
}

// ExtractIDFromToken returns the given token's key ID
func ExtractIDFromToken(token string) (string, error) {
	headerB64 := strings.Split(token, ".")[0]
	headerMarshaled, err := base64.StdEncoding.DecodeString(headerB64)
	if err != nil {
		return "", err
	}
	var header tokenHeader
	if err = json.Unmarshal(headerMarshaled, &header); err != nil {
		return "", err
	}
	return header.Kid, nil
}

// TokenID returns the client's token's key ID if a token is set.
// Returns an empty string when using basic auth.
func (c *Client) TokenID() (string, error) {
	if c.client.token != "" {
		return ExtractIDFromToken(c.client.token)
	}
	return "", nil
}

// DeleteTokenByID deletes the token with the given key ID
func (c *Client) DeleteTokenByID(tokenID string) error {
	url := tokenPath(tokenID)
	return c.client.do(c.client.authApiUrl+url, http.MethodDelete, nil, nil, http.StatusOK)
}

// DeleteToken deletes the given token
func (c *Client) DeleteToken(token string) error {
	tokenID, err := ExtractIDFromToken(token)
	if err != nil {
		return err
	}
	return c.DeleteTokenByID(tokenID)
}

// DeleteCurrentToken deletes the client's token if a token is set.
// Noop when using basic auth.
func (c *Client) DeleteCurrentToken() error {
	if c.client.token != "" {
		return c.DeleteToken(c.client.token)
	}
	return nil
}
