package meili

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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
func (c *Client) Search(ctx context.Context, indexUID, query string, searchOptions ...SearchOption) ([]json.RawMessage, error) {
	key, err := c.getAnyKey()
	if err != nil {
		return nil, err
	}

	newURL := c.getRoute("/indexes/" + indexUID + "/search")
	parsed, err := url.Parse(newURL)
	if err != nil {
		return nil, err
	}

	values := url.Values{"q": []string{query}}
	parsed.RawQuery = values.Encode()

	r, err := http.NewRequest(http.MethodGet, parsed.String(), nil)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)
	r.Header.Set("X-Meili-API-Key", key)

	resp, err := c.makeRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("meili: could not search index, got an unexpected status code '%d' (wanted a 200)", resp.StatusCode)
	}

	var searchResp struct {
		Offset           int               `json:"offset"`
		Limit            int               `json:"limit"`
		ProcessingTimeMs int               `json:"processingTimeMs"`
		Query            string            `json:"query"`
		Hits             []json.RawMessage `json:"hits"`
	}

	err = json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		return nil, err
	}

	return searchResp.Hits, nil
}
