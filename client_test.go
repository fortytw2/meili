package meili

import "testing"

func TestClientInitialization(t *testing.T) {

	var cases = []struct {
		Name      string
		Options   []ClientOption
		ShouldErr bool
	}{
		{
			"basic",
			[]ClientOption{WithNoKeys()},
			false,
		},
		{
			"basic-master-key",
			[]ClientOption{WithMasterKey("test-key")},
			false,
		},
		{
			"basic-private-key",
			[]ClientOption{WithPrivateKey("test-key")},
			false,
		},
		{
			"basic-public-key",
			[]ClientOption{WithPublicKey("test-key")},
			false,
		},
		{
			"basic-no-key",
			[]ClientOption{},
			true,
		},
	}

	for _, c := range cases {
		_, err := NewClient("", c.Options...)
		if !c.ShouldErr && err != nil {
			t.Errorf("failed to initialize client, case %q: %s", c.Name, err)
		}
	}
}
