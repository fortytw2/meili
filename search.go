package meili

import "context"

type SearchParams struct {
	Filters map[string]interface{}

	AttributesToHighlight []string

	ReturnMatches bool

	AttributesToCrop []string
	CropLength       int
}

type SearchOption func(sp *SearchParams)

func WithAttributesToHighlight(attrs []string) SearchOption {
	return func(sp *SearchParams) {
		sp.AttributesToHighlight = attrs
	}
}

// Search is a simple API into meili search
func (c *Client) Search(ctx context.Context, query string, destination interface{}, searchOptions ...SearchOption) error {

	return nil
}
