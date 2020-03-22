package meili

func (c *Client) Write(indexUID string, document interface{}) (int, error) {
	return 0, nil
}

func (c *Client) SynchronousWrite(indexUID string, document interface{}) error {
	return nil
}
