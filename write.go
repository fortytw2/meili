package meili

import "encoding/json"

// Write asynchronously adds a document (or updates an existing one) into meili
// this function returns an updateId that is able to be polled to check on the
// status of the update.
// use SynchronousWrite if you plan on doing that inline and just waiting for
// the write operation to complete.
func (c *Client) WriteOne(indexUID string, document json.RawMessage) (int, error) {

	return 0, nil
}

func (c *Client) Write(indexUID string, documents []json.RawMessage) (int, error) {

	return 0, nil
}

func (c *Client) SynchronousWrite(indexUID string, document json.RawMessage) error {
	return nil
}
