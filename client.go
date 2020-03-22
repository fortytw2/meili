package meili

import (
	"errors"
	"net/http"
)

var (
	// ErrNoKeysProvided occurs when instantiating a meili client with
	// no access keys set and keys not disabled
	ErrNoKeysProvided = errors.New("meili: no key provided and WithNoKeys not set")
)

// a ClientOption is used to configure the client with less-standard options
type ClientOption func(c *Client) error

// DefaultClientOptions are applied automatically
// but can be easily overridden by user supplied ClientOptions
var DefaultClientOptions = []ClientOption{
	WithHTTPClient(http.DefaultClient),
}

func WithNoKeys() ClientOption {
	return func(c *Client) error {
		if !c.hasNoKeys() {
			return errors.New("meili: cannot set WithNoKeys if any keys are provided")
		}

		c.allowNoKeys = true
		return nil
	}
}

// WithMasterKey allows the Client to access all MeiliSearch routes
func WithMasterKey(masterKey string) ClientOption {
	return func(c *Client) error {
		if c.allowNoKeys {
			return errors.New("meili: cannot set a master key if WithNoKeys was explicitly set")
		}

		c.masterKey = masterKey
		return nil
	}
}

// WithPrivateKey allows for access to all MeiliSearch except for the ability to manage keys
func WithPrivateKey(privateKey string) ClientOption {
	return func(c *Client) error {
		if c.allowNoKeys {
			return errors.New("meili: cannot set a private key if WithNoKeys was explicitly set")
		}

		c.privateKey = privateKey
		return nil
	}
}

// WithPublicKey only allows for the ability to search
func WithPublicKey(publicKey string) ClientOption {
	return func(c *Client) error {
		if c.allowNoKeys {
			return errors.New("meili: cannot set a public key if WithNoKeys was explicitly set")
		}

		c.publicKey = publicKey
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// Client is the core meili search client
type Client struct {
	httpClient *http.Client

	address string

	allowNoKeys bool
	masterKey   string
	privateKey  string
	publicKey   string
}

func NewClient(address string, opts ...ClientOption) (*Client, error) {
	c := &Client{
		address: address,
	}

	for _, opt := range DefaultClientOptions {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	if !c.allowNoKeys {
		if c.masterKey == "" && c.publicKey == "" && c.privateKey == "" {
			return nil, ErrNoKeysProvided
		}
	}

	return c, nil
}
