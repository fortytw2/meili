package meili

import "testing"

func TestClientKeyHelpers(t *testing.T) {
	var (
		masterKey  = "link"
		privateKey = "private"
		pubKey     = "public"
	)

	var cases = []struct {
		Name      string
		Options   []ClientOption
		Get       func(c *Client) func() (string, error)
		ExpectKey string
		ShouldErr bool
	}{
		{
			"basic",
			[]ClientOption{WithNoKeys()},
			func(c *Client) func() (string, error) { return c.getMasterKey },
			"",
			false,
		},
		{
			"basic-2",
			[]ClientOption{WithNoKeys()},
			func(c *Client) func() (string, error) { return c.getAnyKey },
			"",
			false,
		},
		{
			"basic-master-key",
			[]ClientOption{WithMasterKey(masterKey)},
			func(c *Client) func() (string, error) { return c.getMasterKey },
			masterKey,
			false,
		},
		{
			"basic-private-key",
			[]ClientOption{WithPrivateKey(privateKey)},
			func(c *Client) func() (string, error) { return c.getMasterOrPrivateKey },
			privateKey,
			false,
		},
		{
			"basic-public-key",
			[]ClientOption{WithPublicKey(pubKey)},
			func(c *Client) func() (string, error) { return c.getAnyKey },
			pubKey,
			false,
		},
	}

	for _, c := range cases {
		client, err := NewClient("", c.Options...)
		if err != nil {
			t.Errorf("failed to initialize client, case %q: %s", c.Name, err)
		}

		key, err := c.Get(client)()
		if err != nil {
			t.Errorf("failed to get key, case %q: %s", c.Name, err)
		}

		if key != c.ExpectKey {
			t.Errorf("did not get the right key for case %q: got %q want %q", c.Name, key, c.ExpectKey)
		}
	}
}
