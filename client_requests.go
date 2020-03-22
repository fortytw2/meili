package meili

import (
	"errors"
	"net/http"
)

func (c *Client) makeRequest(r *http.Request) (*http.Response, error) {
	return c.httpClient.Do(r)
}

func (c *Client) hasNoKeys() bool {
	if c.allowNoKeys {
		return false
	}

	if c.masterKey == "" && c.privateKey == "" && c.publicKey == "" {
		return true
	}

	return false
}

func (c *Client) getMasterKey() (string, error) {
	if c.hasNoKeys() {
		return "", ErrNoKeysProvided
	}

	if c.masterKey != "" {
		return c.masterKey, nil
	}

	if c.allowNoKeys {
		return "", nil
	}

	return "", errors.New("meili: operation requires a master key to execute")
}

func (c *Client) getMasterOrPrivateKey() (string, error) {
	if c.hasNoKeys() {
		return "", ErrNoKeysProvided
	}

	if c.masterKey != "" {
		return c.masterKey, nil
	}

	if c.privateKey != "" {
		return c.privateKey, nil
	}

	if c.allowNoKeys {
		return "", nil
	}

	return "", errors.New("meili: operation requires a master or private key to execute")
}

func (c *Client) getAnyKey() (string, error) {
	if c.hasNoKeys() {
		return "", ErrNoKeysProvided
	}

	if c.masterKey != "" {
		return c.masterKey, nil
	}

	if c.privateKey != "" {
		return c.privateKey, nil
	}

	if c.publicKey != "" {
		return c.publicKey, nil
	}

	return "", nil
}
