package meili

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Write asynchronously adds a document (or updates an existing one) into meili
// this function returns an updateId that is able to be polled to check on the
// status of the update.
// use SynchronousWrite if you plan on doing that inline and just waiting for
// the write operation to complete.
func (c *Client) WriteOne(ctx context.Context, indexUID string, document json.RawMessage) (int, error) {
	return c.Write(ctx, indexUID, []json.RawMessage{document})
}

func (c *Client) Write(ctx context.Context, indexUID string, documents []json.RawMessage) (int, error) {
	key, err := c.getMasterOrPrivateKey()
	if err != nil {
		return 0, err
	}

	docs, err := json.Marshal(documents)
	if err != nil {
		return 0, err
	}

	r, err := http.NewRequest(http.MethodPost, c.getRoute("/indexes/"+indexUID+"/documents"), bytes.NewReader(docs))
	if err != nil {
		return 0, err
	}

	r = r.WithContext(ctx)
	r.Header.Set("X-Meili-API-Key", key)

	resp, err := c.makeRequest(r)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return 0, nil
}

func (c *Client) SynchronousWrite(indexUID string, document json.RawMessage) error {
	return nil
}
